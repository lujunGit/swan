// Code generated by protoc-gen-gogo.
// source: raft.proto
// DO NOT EDIT!

package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "gogoproto"

import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// StoreActionKind defines the operation to take on the store for the target of
// a storage action.
type StoreActionKind int32

const (
	StoreActionKindUnknown StoreActionKind = 0
	StoreActionKindCreate  StoreActionKind = 1
	StoreActionKindUpdate  StoreActionKind = 2
	StoreActionKindRemove  StoreActionKind = 3
)

var StoreActionKind_name = map[int32]string{
	0: "UNKNOWN",
	1: "STORE_ACTION_CREATE",
	2: "STORE_ACTION_UPDATE",
	3: "STORE_ACTION_REMOVE",
}
var StoreActionKind_value = map[string]int32{
	"UNKNOWN":             0,
	"STORE_ACTION_CREATE": 1,
	"STORE_ACTION_UPDATE": 2,
	"STORE_ACTION_REMOVE": 3,
}

func (x StoreActionKind) String() string {
	return proto.EnumName(StoreActionKind_name, int32(x))
}
func (StoreActionKind) EnumDescriptor() ([]byte, []int) { return fileDescriptorRaft, []int{0} }

// Contains one of many protobuf encoded objects to replicate
// over the raft backend with a request ID to track when the
// action is effectively applied
type InternalRaftRequest struct {
	ID     uint64         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Action []*StoreAction `protobuf:"bytes,2,rep,name=action" json:"action,omitempty"`
}

func (m *InternalRaftRequest) Reset()                    { *m = InternalRaftRequest{} }
func (m *InternalRaftRequest) String() string            { return proto.CompactTextString(m) }
func (*InternalRaftRequest) ProtoMessage()               {}
func (*InternalRaftRequest) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{0} }

// StoreAction defines a target and operation to apply on the storage system.
type StoreAction struct {
	Action StoreActionKind `protobuf:"varint,1,opt,name=action,proto3,enum=types.StoreActionKind" json:"action,omitempty"`
	// Types that are valid to be assigned to Target:
	//	*StoreAction_Application
	Target isStoreAction_Target `protobuf_oneof:"target"`
}

func (m *StoreAction) Reset()                    { *m = StoreAction{} }
func (m *StoreAction) String() string            { return proto.CompactTextString(m) }
func (*StoreAction) ProtoMessage()               {}
func (*StoreAction) Descriptor() ([]byte, []int) { return fileDescriptorRaft, []int{1} }

type isStoreAction_Target interface {
	isStoreAction_Target()
	Equal(interface{}) bool
	VerboseEqual(interface{}) error
	MarshalTo([]byte) (int, error)
	Size() int
}

type StoreAction_Application struct {
	Application *Application `protobuf:"bytes,2,opt,name=application,oneof"`
}

func (*StoreAction_Application) isStoreAction_Target() {}

func (m *StoreAction) GetTarget() isStoreAction_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *StoreAction) GetApplication() *Application {
	if x, ok := m.GetTarget().(*StoreAction_Application); ok {
		return x.Application
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StoreAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StoreAction_OneofMarshaler, _StoreAction_OneofUnmarshaler, _StoreAction_OneofSizer, []interface{}{
		(*StoreAction_Application)(nil),
	}
}

func _StoreAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StoreAction)
	// target
	switch x := m.Target.(type) {
	case *StoreAction_Application:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Application); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StoreAction.Target has unexpected type %T", x)
	}
	return nil
}

func _StoreAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StoreAction)
	switch tag {
	case 2: // target.application
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Application)
		err := b.DecodeMessage(msg)
		m.Target = &StoreAction_Application{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StoreAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StoreAction)
	// target
	switch x := m.Target.(type) {
	case *StoreAction_Application:
		s := proto.Size(x.Application)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*InternalRaftRequest)(nil), "types.InternalRaftRequest")
	proto.RegisterType((*StoreAction)(nil), "types.StoreAction")
	proto.RegisterEnum("types.StoreActionKind", StoreActionKind_name, StoreActionKind_value)
}
func (this *InternalRaftRequest) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*InternalRaftRequest)
	if !ok {
		that2, ok := that.(InternalRaftRequest)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *InternalRaftRequest")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *InternalRaftRequest but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *InternalRaftRequest but is not nil && this == nil")
	}
	if this.ID != that1.ID {
		return fmt.Errorf("ID this(%v) Not Equal that(%v)", this.ID, that1.ID)
	}
	if len(this.Action) != len(that1.Action) {
		return fmt.Errorf("Action this(%v) Not Equal that(%v)", len(this.Action), len(that1.Action))
	}
	for i := range this.Action {
		if !this.Action[i].Equal(that1.Action[i]) {
			return fmt.Errorf("Action this[%v](%v) Not Equal that[%v](%v)", i, this.Action[i], i, that1.Action[i])
		}
	}
	return nil
}
func (this *InternalRaftRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*InternalRaftRequest)
	if !ok {
		that2, ok := that.(InternalRaftRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if len(this.Action) != len(that1.Action) {
		return false
	}
	for i := range this.Action {
		if !this.Action[i].Equal(that1.Action[i]) {
			return false
		}
	}
	return true
}
func (this *StoreAction) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*StoreAction)
	if !ok {
		that2, ok := that.(StoreAction)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *StoreAction")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *StoreAction but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *StoreAction but is not nil && this == nil")
	}
	if this.Action != that1.Action {
		return fmt.Errorf("Action this(%v) Not Equal that(%v)", this.Action, that1.Action)
	}
	if that1.Target == nil {
		if this.Target != nil {
			return fmt.Errorf("this.Target != nil && that1.Target == nil")
		}
	} else if this.Target == nil {
		return fmt.Errorf("this.Target == nil && that1.Target != nil")
	} else if err := this.Target.VerboseEqual(that1.Target); err != nil {
		return err
	}
	return nil
}
func (this *StoreAction_Application) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*StoreAction_Application)
	if !ok {
		that2, ok := that.(StoreAction_Application)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *StoreAction_Application")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *StoreAction_Application but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *StoreAction_Application but is not nil && this == nil")
	}
	if !this.Application.Equal(that1.Application) {
		return fmt.Errorf("Application this(%v) Not Equal that(%v)", this.Application, that1.Application)
	}
	return nil
}
func (this *StoreAction) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*StoreAction)
	if !ok {
		that2, ok := that.(StoreAction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Action != that1.Action {
		return false
	}
	if that1.Target == nil {
		if this.Target != nil {
			return false
		}
	} else if this.Target == nil {
		return false
	} else if !this.Target.Equal(that1.Target) {
		return false
	}
	return true
}
func (this *StoreAction_Application) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*StoreAction_Application)
	if !ok {
		that2, ok := that.(StoreAction_Application)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Application.Equal(that1.Application) {
		return false
	}
	return true
}
func (this *InternalRaftRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&types.InternalRaftRequest{")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	if this.Action != nil {
		s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StoreAction) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&types.StoreAction{")
	s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	if this.Target != nil {
		s = append(s, "Target: "+fmt.Sprintf("%#v", this.Target)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StoreAction_Application) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&types.StoreAction_Application{` +
		`Application:` + fmt.Sprintf("%#v", this.Application) + `}`}, ", ")
	return s
}
func valueToGoStringRaft(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringRaft(m github_com_gogo_protobuf_proto.Message) string {
	e := github_com_gogo_protobuf_proto.GetUnsafeExtensionsMap(m)
	if e == nil {
		return "nil"
	}
	s := "proto.NewUnsafeXXX_InternalExtensions(map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "})"
	return s
}
func (m *InternalRaftRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InternalRaftRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaft(dAtA, i, uint64(m.ID))
	}
	if len(m.Action) > 0 {
		for _, msg := range m.Action {
			dAtA[i] = 0x12
			i++
			i = encodeVarintRaft(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StoreAction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreAction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Action != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRaft(dAtA, i, uint64(m.Action))
	}
	if m.Target != nil {
		nn1, err := m.Target.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *StoreAction_Application) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	if m.Application != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRaft(dAtA, i, uint64(m.Application.Size()))
		n2, err := m.Application.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}
func encodeFixed64Raft(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Raft(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintRaft(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedInternalRaftRequest(r randyRaft, easy bool) *InternalRaftRequest {
	this := &InternalRaftRequest{}
	this.ID = uint64(uint64(r.Uint32()))
	if r.Intn(10) != 0 {
		v1 := r.Intn(5)
		this.Action = make([]*StoreAction, v1)
		for i := 0; i < v1; i++ {
			this.Action[i] = NewPopulatedStoreAction(r, easy)
		}
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStoreAction(r randyRaft, easy bool) *StoreAction {
	this := &StoreAction{}
	this.Action = StoreActionKind([]int32{0, 1, 2, 3}[r.Intn(4)])
	oneofNumber_Target := []int32{2}[r.Intn(1)]
	switch oneofNumber_Target {
	case 2:
		this.Target = NewPopulatedStoreAction_Application(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedStoreAction_Application(r randyRaft, easy bool) *StoreAction_Application {
	this := &StoreAction_Application{}
	this.Application = NewPopulatedApplication(r, easy)
	return this
}

type randyRaft interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneRaft(r randyRaft) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringRaft(r randyRaft) string {
	v2 := r.Intn(100)
	tmps := make([]rune, v2)
	for i := 0; i < v2; i++ {
		tmps[i] = randUTF8RuneRaft(r)
	}
	return string(tmps)
}
func randUnrecognizedRaft(r randyRaft, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldRaft(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldRaft(dAtA []byte, r randyRaft, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateRaft(dAtA, uint64(key))
		v3 := r.Int63()
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		dAtA = encodeVarintPopulateRaft(dAtA, uint64(v3))
	case 1:
		dAtA = encodeVarintPopulateRaft(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateRaft(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateRaft(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateRaft(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateRaft(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *InternalRaftRequest) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovRaft(uint64(m.ID))
	}
	if len(m.Action) > 0 {
		for _, e := range m.Action {
			l = e.Size()
			n += 1 + l + sovRaft(uint64(l))
		}
	}
	return n
}

func (m *StoreAction) Size() (n int) {
	var l int
	_ = l
	if m.Action != 0 {
		n += 1 + sovRaft(uint64(m.Action))
	}
	if m.Target != nil {
		n += m.Target.Size()
	}
	return n
}

func (m *StoreAction_Application) Size() (n int) {
	var l int
	_ = l
	if m.Application != nil {
		l = m.Application.Size()
		n += 1 + l + sovRaft(uint64(l))
	}
	return n
}

func sovRaft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRaft(x uint64) (n int) {
	return sovRaft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InternalRaftRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InternalRaftRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InternalRaftRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Action = append(m.Action, &StoreAction{})
			if err := m.Action[len(m.Action)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StoreAction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StoreAction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreAction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			m.Action = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Action |= (StoreActionKind(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Application", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRaft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &Application{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Target = &StoreAction_Application{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRaft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRaft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRaft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRaft
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRaft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRaft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRaft
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRaft(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRaft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRaft   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("raft.proto", fileDescriptorRaft) }

var fileDescriptorRaft = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4a, 0x4c, 0x2b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x96, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xe8, 0x83, 0x58, 0x10, 0x49, 0x29, 0xc1, 0xc4, 0x82, 0x82, 0x9c,
	0xcc, 0xe4, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x88, 0x90, 0x52, 0x24, 0x97, 0xb0, 0x67, 0x5e, 0x49,
	0x6a, 0x51, 0x5e, 0x62, 0x4e, 0x10, 0xd0, 0x94, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x31, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x16, 0x27, 0xb6, 0x47, 0xf7, 0xe4,
	0x99, 0x3c, 0x5d, 0x82, 0x80, 0x22, 0x42, 0x5a, 0x5c, 0x6c, 0x89, 0xc9, 0x20, 0xed, 0x12, 0x4c,
	0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x42, 0x7a, 0x60, 0xfb, 0xf4, 0x82, 0x4b, 0xf2, 0x8b, 0x52, 0x1d,
	0xc1, 0x32, 0x41, 0x50, 0x15, 0x4a, 0xf5, 0x5c, 0xdc, 0x48, 0xc2, 0x42, 0x7a, 0x70, 0xad, 0x20,
	0x63, 0xf9, 0x8c, 0xc4, 0x30, 0xb5, 0x7a, 0x67, 0xe6, 0xa5, 0xc0, 0xb4, 0x0b, 0x99, 0x71, 0x71,
	0x23, 0x39, 0x17, 0x68, 0x1f, 0x23, 0x92, 0x7d, 0x8e, 0x08, 0x19, 0x0f, 0x86, 0x20, 0x64, 0x85,
	0x4e, 0x1c, 0x5c, 0x6c, 0x25, 0x89, 0x45, 0xe9, 0xa9, 0x25, 0x5a, 0xef, 0x19, 0xb9, 0xf8, 0xd1,
	0x4c, 0x17, 0x52, 0xe7, 0x62, 0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x90, 0x92,
	0xea, 0x9a, 0xab, 0x20, 0x86, 0xa6, 0x22, 0x34, 0x2f, 0x3b, 0x2f, 0xbf, 0x3c, 0x4f, 0xc8, 0x88,
	0x4b, 0x38, 0x38, 0xc4, 0x3f, 0xc8, 0x35, 0xde, 0xd1, 0x39, 0xc4, 0xd3, 0xdf, 0x2f, 0xde, 0x39,
	0xc8, 0xd5, 0x31, 0xc4, 0x55, 0x80, 0x51, 0x4a, 0x12, 0xa8, 0x49, 0x14, 0x4d, 0x93, 0x73, 0x51,
	0x6a, 0x62, 0x49, 0x2a, 0x86, 0x9e, 0xd0, 0x00, 0x17, 0x90, 0x1e, 0x26, 0xac, 0x7a, 0x42, 0x0b,
	0x52, 0xb0, 0xe9, 0x09, 0x72, 0xf5, 0xf5, 0x0f, 0x73, 0x15, 0x60, 0xc6, 0xaa, 0x27, 0x28, 0x35,
	0x37, 0xbf, 0x2c, 0x55, 0x4a, 0xbc, 0x63, 0xb1, 0x1c, 0xc3, 0xae, 0x25, 0x72, 0xe8, 0xbe, 0x73,
	0x52, 0x39, 0xf1, 0x50, 0x8e, 0xe1, 0xc1, 0x43, 0x39, 0xc6, 0x0f, 0x40, 0xfc, 0x03, 0x88, 0x57,
	0x3c, 0x92, 0x63, 0xdc, 0x01, 0xc4, 0x27, 0x80, 0xf8, 0x02, 0x10, 0x3f, 0x00, 0xe2, 0x08, 0x86,
	0x24, 0x36, 0x70, 0xe4, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x45, 0xa6, 0x11, 0xc9, 0x3a,
	0x02, 0x00, 0x00,
}
