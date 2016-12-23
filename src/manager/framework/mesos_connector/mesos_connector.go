package mesos_connector

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Dataman-Cloud/swan/src/config"
	"github.com/Dataman-Cloud/swan/src/manager/framework/event"
	"github.com/Dataman-Cloud/swan/src/mesosproto/mesos"
	"github.com/Dataman-Cloud/swan/src/mesosproto/sched"

	"github.com/Sirupsen/logrus"
	"github.com/andygrunwald/megos"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
)

var instance *MesosConnector

type MesosConnector struct {
	// mesos framework related
	ClusterId        string
	master           string
	client           *MesosHttpClient
	lastHearBeatTime time.Time
	config           config.Scheduler

	MesosCallChan chan *sched.Call

	// TODO make sure this chan doesn't explode
	MesosEventChan chan *event.MesosEvent
	Framework      *mesos.FrameworkInfo
}

func NewMesosConnector(config config.Scheduler) *MesosConnector {
	instance = &MesosConnector{
		config:         config,
		MesosEventChan: make(chan *event.MesosEvent, 1024), // make this unbound in future
		MesosCallChan:  make(chan *sched.Call, 1024),
	}

	return instance
}

func Instance() *MesosConnector {
	if instance == nil {
		logrus.Errorf("mesos connector is nil now, need reconnect")
		return nil
	} else {
		return instance
	}
}

func (s *MesosConnector) subscribe(ctx context.Context, mesosFailureChan chan error) {
	logrus.Infof("Subscribe with mesos master %s", s.master)
	call := &sched.Call{
		Type: sched.Call_SUBSCRIBE.Enum(),
		Subscribe: &sched.Call_Subscribe{
			FrameworkInfo: s.Framework,
		},
	}

	if s.Framework.Id != nil {
		call.FrameworkId = &mesos.FrameworkID{
			Value: proto.String(s.Framework.Id.GetValue()),
		}
	}

	resp, err := s.Send(call)
	if err != nil {
		mesosFailureChan <- err
	}

	// http might now be the default transport in future release
	if resp.StatusCode != http.StatusOK {
		mesosFailureChan <- fmt.Errorf("Subscribe with unexpected response status: %d", resp.StatusCode)
	}

	logrus.Info(s.client.StreamID)
	go s.handleEvents(ctx, resp, mesosFailureChan)
}

func (s *MesosConnector) handleEvents(ctx context.Context, resp *http.Response, mesosFailureChan chan error) {
	defer func() {
		resp.Body.Close()
	}()

	r := NewReader(resp.Body)
	dec := json.NewDecoder(r)

	for {
		select {
		case <-ctx.Done():
			logrus.Infof("handleEvents cancelled %s", ctx.Err())
			return
		default:
			event := new(sched.Event)
			if err := dec.Decode(event); err != nil {
				logrus.Errorf("Deocde event failed: %s", err)
				mesosFailureChan <- err
			}

			switch event.GetType() {
			case sched.Event_SUBSCRIBED:
				s.addEvent(sched.Event_SUBSCRIBED, event)
			case sched.Event_OFFERS:
				s.addEvent(sched.Event_OFFERS, event)
			case sched.Event_RESCIND:
				s.addEvent(sched.Event_RESCIND, event)
			case sched.Event_UPDATE:
				s.addEvent(sched.Event_UPDATE, event)
			case sched.Event_MESSAGE:
				s.addEvent(sched.Event_MESSAGE, event)
			case sched.Event_FAILURE:
				s.addEvent(sched.Event_FAILURE, event)
			case sched.Event_ERROR:
				s.addEvent(sched.Event_ERROR, event)
			case sched.Event_HEARTBEAT:
				s.addEvent(sched.Event_HEARTBEAT, event)
			}
		}
	}
}

func CreateFrameworkInfo(config config.Scheduler) *mesos.FrameworkInfo {
	fw := &mesos.FrameworkInfo{
		User:            proto.String(config.MesosFrameworkUser),
		Name:            proto.String("swan"),
		Hostname:        proto.String(config.Hostname),
		FailoverTimeout: proto.Float64(60 * 60 * 24 * 7),
	}

	return fw
}

func stateFromMasters(masters []string) (*megos.State, error) {
	masterUrls := make([]*url.URL, 0)
	for _, master := range masters {
		masterUrl, _ := url.Parse(fmt.Sprintf("http://%s", master))
		masterUrls = append(masterUrls, masterUrl)
	}

	mesos := megos.NewClient(masterUrls, nil)
	return mesos.GetStateFromCluster()
}

func (s *MesosConnector) Send(call *sched.Call) (*http.Response, error) {
	payload, err := proto.Marshal(call)
	if err != nil {
		return nil, err
	}
	return s.client.Send(payload)
}

func (s *MesosConnector) addEvent(eventType sched.Event_Type, e *sched.Event) {
	s.MesosEventChan <- &event.MesosEvent{EventType: eventType, Event: e}
}

func (s *MesosConnector) Start(ctx context.Context, mesosFailureChan chan error) {
	var err error
	state, err := stateFromMasters(strings.Split(s.config.MesosMasters, ","))
	if err != nil {
		logrus.Errorf("%s Check your mesos master configuration", err)
		mesosFailureChan <- err
	}

	s.master = state.Leader
	s.client = NewHTTPClient(state.Leader, "/api/v1/scheduler")

	s.ClusterId = state.Cluster
	if s.ClusterId == "" {
		s.ClusterId = "unamed"
	}

	s.subscribe(ctx, mesosFailureChan)

	for {
		select {
		case <-ctx.Done():
			logrus.Errorf("mesosConnector got signal %s", ctx.Err())
			return
		case call := <-s.MesosCallChan:
			logrus.WithFields(logrus.Fields{"sending-call": sched.Call_Type_name[int32(*call.Type)]}).Debugf("")
			resp, err := s.Send(call)
			if err != nil {
				logrus.Errorf("%s", err)
				mesosFailureChan <- err
			}
			if resp.StatusCode != 202 {
				logrus.Infof("send response not 202 but %d", resp.StatusCode)
				mesosFailureChan <- errors.New("http got respose not 202")
			}
		}
	}
}
