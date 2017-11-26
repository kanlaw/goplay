// Code generated by protoc-gen-gogo.
// source: lottery.proto
// DO NOT EDIT!

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 全民刮奖
type CLotteryInfo struct {
	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (m *CLotteryInfo) Reset()                    { *m = CLotteryInfo{} }
func (*CLotteryInfo) ProtoMessage()               {}
func (*CLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptorLottery, []int{0} }

func (m *CLotteryInfo) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type SLotteryInfo struct {
	Code      uint32     `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	List      []*Lottery `protobuf:"bytes,2,rep,name=list" json:"list,omitempty"`
	Single    uint32     `protobuf:"varint,3,opt,name=single,proto3" json:"single,omitempty"`
	Maxnumber uint32     `protobuf:"varint,4,opt,name=maxnumber,proto3" json:"maxnumber,omitempty"`
	Error     ErrCode    `protobuf:"varint,5,opt,name=error,proto3,enum=pb.ErrCode" json:"error,omitempty"`
}

func (m *SLotteryInfo) Reset()                    { *m = SLotteryInfo{} }
func (*SLotteryInfo) ProtoMessage()               {}
func (*SLotteryInfo) Descriptor() ([]byte, []int) { return fileDescriptorLottery, []int{1} }

func (m *SLotteryInfo) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SLotteryInfo) GetList() []*Lottery {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *SLotteryInfo) GetSingle() uint32 {
	if m != nil {
		return m.Single
	}
	return 0
}

func (m *SLotteryInfo) GetMaxnumber() uint32 {
	if m != nil {
		return m.Maxnumber
	}
	return 0
}

func (m *SLotteryInfo) GetError() ErrCode {
	if m != nil {
		return m.Error
	}
	return OK
}

// 刮奖机会
type Lottery struct {
	Niu    uint32 `protobuf:"varint,1,opt,name=niu,proto3" json:"niu,omitempty"`
	Number uint32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (m *Lottery) Reset()                    { *m = Lottery{} }
func (*Lottery) ProtoMessage()               {}
func (*Lottery) Descriptor() ([]byte, []int) { return fileDescriptorLottery, []int{2} }

func (m *Lottery) GetNiu() uint32 {
	if m != nil {
		return m.Niu
	}
	return 0
}

func (m *Lottery) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

// 刮奖
type CLottery struct {
	Code  uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Times uint32 `protobuf:"varint,2,opt,name=times,proto3" json:"times,omitempty"`
}

func (m *CLottery) Reset()                    { *m = CLottery{} }
func (*CLottery) ProtoMessage()               {}
func (*CLottery) Descriptor() ([]byte, []int) { return fileDescriptorLottery, []int{3} }

func (m *CLottery) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CLottery) GetTimes() uint32 {
	if m != nil {
		return m.Times
	}
	return 0
}

type SLottery struct {
	Code   uint32   `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Cards  []uint32 `protobuf:"varint,2,rep,packed,name=cards" json:"cards,omitempty"`
	Niu    uint32   `protobuf:"varint,3,opt,name=niu,proto3" json:"niu,omitempty"`
	Number uint32   `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	Error  ErrCode  `protobuf:"varint,5,opt,name=error,proto3,enum=pb.ErrCode" json:"error,omitempty"`
}

func (m *SLottery) Reset()                    { *m = SLottery{} }
func (*SLottery) ProtoMessage()               {}
func (*SLottery) Descriptor() ([]byte, []int) { return fileDescriptorLottery, []int{4} }

func (m *SLottery) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SLottery) GetCards() []uint32 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *SLottery) GetNiu() uint32 {
	if m != nil {
		return m.Niu
	}
	return 0
}

func (m *SLottery) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *SLottery) GetError() ErrCode {
	if m != nil {
		return m.Error
	}
	return OK
}

func init() {
	proto.RegisterType((*CLotteryInfo)(nil), "pb.CLotteryInfo")
	proto.RegisterType((*SLotteryInfo)(nil), "pb.SLotteryInfo")
	proto.RegisterType((*Lottery)(nil), "pb.Lottery")
	proto.RegisterType((*CLottery)(nil), "pb.CLottery")
	proto.RegisterType((*SLottery)(nil), "pb.SLottery")
}
func (this *CLotteryInfo) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*CLotteryInfo)
	if !ok {
		that2, ok := that.(CLotteryInfo)
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
	if this.Code != that1.Code {
		return false
	}
	return true
}
func (this *SLotteryInfo) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SLotteryInfo)
	if !ok {
		that2, ok := that.(SLotteryInfo)
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
	if this.Code != that1.Code {
		return false
	}
	if len(this.List) != len(that1.List) {
		return false
	}
	for i := range this.List {
		if !this.List[i].Equal(that1.List[i]) {
			return false
		}
	}
	if this.Single != that1.Single {
		return false
	}
	if this.Maxnumber != that1.Maxnumber {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *Lottery) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Lottery)
	if !ok {
		that2, ok := that.(Lottery)
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
	if this.Niu != that1.Niu {
		return false
	}
	if this.Number != that1.Number {
		return false
	}
	return true
}
func (this *CLottery) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*CLottery)
	if !ok {
		that2, ok := that.(CLottery)
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
	if this.Code != that1.Code {
		return false
	}
	if this.Times != that1.Times {
		return false
	}
	return true
}
func (this *SLottery) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SLottery)
	if !ok {
		that2, ok := that.(SLottery)
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
	if this.Code != that1.Code {
		return false
	}
	if len(this.Cards) != len(that1.Cards) {
		return false
	}
	for i := range this.Cards {
		if this.Cards[i] != that1.Cards[i] {
			return false
		}
	}
	if this.Niu != that1.Niu {
		return false
	}
	if this.Number != that1.Number {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *CLotteryInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.CLotteryInfo{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SLotteryInfo) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&pb.SLotteryInfo{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	if this.List != nil {
		s = append(s, "List: "+fmt.Sprintf("%#v", this.List)+",\n")
	}
	s = append(s, "Single: "+fmt.Sprintf("%#v", this.Single)+",\n")
	s = append(s, "Maxnumber: "+fmt.Sprintf("%#v", this.Maxnumber)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Lottery) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.Lottery{")
	s = append(s, "Niu: "+fmt.Sprintf("%#v", this.Niu)+",\n")
	s = append(s, "Number: "+fmt.Sprintf("%#v", this.Number)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CLottery) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.CLottery{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Times: "+fmt.Sprintf("%#v", this.Times)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *SLottery) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&pb.SLottery{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Cards: "+fmt.Sprintf("%#v", this.Cards)+",\n")
	s = append(s, "Niu: "+fmt.Sprintf("%#v", this.Niu)+",\n")
	s = append(s, "Number: "+fmt.Sprintf("%#v", this.Number)+",\n")
	s = append(s, "Error: "+fmt.Sprintf("%#v", this.Error)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringLottery(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CLotteryInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CLotteryInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Code))
	}
	return i, nil
}

func (m *SLotteryInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SLotteryInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Code))
	}
	if len(m.List) > 0 {
		for _, msg := range m.List {
			dAtA[i] = 0x12
			i++
			i = encodeVarintLottery(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Single != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Single))
	}
	if m.Maxnumber != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Maxnumber))
	}
	if m.Error != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Error))
	}
	return i, nil
}

func (m *Lottery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lottery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Niu != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Niu))
	}
	if m.Number != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Number))
	}
	return i, nil
}

func (m *CLottery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CLottery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Code))
	}
	if m.Times != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Times))
	}
	return i, nil
}

func (m *SLottery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SLottery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Code))
	}
	if len(m.Cards) > 0 {
		dAtA2 := make([]byte, len(m.Cards)*10)
		var j1 int
		for _, num := range m.Cards {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintLottery(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	if m.Niu != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Niu))
	}
	if m.Number != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Number))
	}
	if m.Error != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintLottery(dAtA, i, uint64(m.Error))
	}
	return i, nil
}

func encodeFixed64Lottery(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Lottery(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLottery(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CLotteryInfo) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovLottery(uint64(m.Code))
	}
	return n
}

func (m *SLotteryInfo) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovLottery(uint64(m.Code))
	}
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovLottery(uint64(l))
		}
	}
	if m.Single != 0 {
		n += 1 + sovLottery(uint64(m.Single))
	}
	if m.Maxnumber != 0 {
		n += 1 + sovLottery(uint64(m.Maxnumber))
	}
	if m.Error != 0 {
		n += 1 + sovLottery(uint64(m.Error))
	}
	return n
}

func (m *Lottery) Size() (n int) {
	var l int
	_ = l
	if m.Niu != 0 {
		n += 1 + sovLottery(uint64(m.Niu))
	}
	if m.Number != 0 {
		n += 1 + sovLottery(uint64(m.Number))
	}
	return n
}

func (m *CLottery) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovLottery(uint64(m.Code))
	}
	if m.Times != 0 {
		n += 1 + sovLottery(uint64(m.Times))
	}
	return n
}

func (m *SLottery) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovLottery(uint64(m.Code))
	}
	if len(m.Cards) > 0 {
		l = 0
		for _, e := range m.Cards {
			l += sovLottery(uint64(e))
		}
		n += 1 + sovLottery(uint64(l)) + l
	}
	if m.Niu != 0 {
		n += 1 + sovLottery(uint64(m.Niu))
	}
	if m.Number != 0 {
		n += 1 + sovLottery(uint64(m.Number))
	}
	if m.Error != 0 {
		n += 1 + sovLottery(uint64(m.Error))
	}
	return n
}

func sovLottery(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLottery(x uint64) (n int) {
	return sovLottery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CLotteryInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CLotteryInfo{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SLotteryInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SLotteryInfo{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`List:` + strings.Replace(fmt.Sprintf("%v", this.List), "Lottery", "Lottery", 1) + `,`,
		`Single:` + fmt.Sprintf("%v", this.Single) + `,`,
		`Maxnumber:` + fmt.Sprintf("%v", this.Maxnumber) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Lottery) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Lottery{`,
		`Niu:` + fmt.Sprintf("%v", this.Niu) + `,`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CLottery) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CLottery{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Times:` + fmt.Sprintf("%v", this.Times) + `,`,
		`}`,
	}, "")
	return s
}
func (this *SLottery) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SLottery{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Cards:` + fmt.Sprintf("%v", this.Cards) + `,`,
		`Niu:` + fmt.Sprintf("%v", this.Niu) + `,`,
		`Number:` + fmt.Sprintf("%v", this.Number) + `,`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringLottery(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CLotteryInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLottery
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
			return fmt.Errorf("proto: CLotteryInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CLotteryInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLottery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLottery
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
func (m *SLotteryInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLottery
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
			return fmt.Errorf("proto: SLotteryInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SLotteryInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
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
				return ErrInvalidLengthLottery
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, &Lottery{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Single", wireType)
			}
			m.Single = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Single |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Maxnumber", wireType)
			}
			m.Maxnumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Maxnumber |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Error |= (ErrCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLottery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLottery
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
func (m *Lottery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLottery
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
			return fmt.Errorf("proto: Lottery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lottery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Niu", wireType)
			}
			m.Niu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Niu |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLottery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLottery
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
func (m *CLottery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLottery
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
			return fmt.Errorf("proto: CLottery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CLottery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Times", wireType)
			}
			m.Times = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Times |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLottery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLottery
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
func (m *SLottery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLottery
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
			return fmt.Errorf("proto: SLottery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SLottery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLottery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint32(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Cards = append(m.Cards, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowLottery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthLottery
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowLottery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Cards = append(m.Cards, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Cards", wireType)
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Niu", wireType)
			}
			m.Niu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Niu |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			m.Error = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLottery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Error |= (ErrCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLottery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLottery
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
func skipLottery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLottery
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
					return 0, ErrIntOverflowLottery
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
					return 0, ErrIntOverflowLottery
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
				return 0, ErrInvalidLengthLottery
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLottery
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
				next, err := skipLottery(dAtA[start:])
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
	ErrInvalidLengthLottery = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLottery   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("lottery.proto", fileDescriptorLottery) }

var fileDescriptorLottery = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x4a,
	0xce, 0x4f, 0x49, 0x85, 0xf0, 0x95, 0x94, 0xb8, 0x78, 0x9c, 0x7d, 0x20, 0x2a, 0x3c, 0xf3, 0xd2,
	0xf2, 0x85, 0x84, 0xb8, 0x58, 0x40, 0xb2, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x60, 0xb6,
	0xd2, 0x1c, 0x46, 0x2e, 0x9e, 0x60, 0x02, 0x8a, 0x84, 0xe4, 0xb9, 0x58, 0x72, 0x32, 0x8b, 0x4b,
	0x24, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xb8, 0xf5, 0x0a, 0x92, 0xf4, 0xa0, 0x5a, 0x82, 0xc0,
	0x12, 0x42, 0x62, 0x5c, 0x6c, 0xc5, 0x99, 0x79, 0xe9, 0x39, 0xa9, 0x12, 0xcc, 0x60, 0x6d, 0x50,
	0x9e, 0x90, 0x0c, 0x17, 0x67, 0x6e, 0x62, 0x45, 0x5e, 0x69, 0x6e, 0x52, 0x6a, 0x91, 0x04, 0x0b,
	0x58, 0x0a, 0x21, 0x20, 0xa4, 0xc8, 0xc5, 0x9a, 0x5a, 0x54, 0x94, 0x5f, 0x24, 0xc1, 0xaa, 0xc0,
	0xa8, 0xc1, 0x07, 0x31, 0xd7, 0xb5, 0xa8, 0xc8, 0x39, 0x3f, 0x25, 0x35, 0x08, 0x22, 0xa3, 0x64,
	0xcc, 0xc5, 0x0e, 0xb5, 0x49, 0x48, 0x80, 0x8b, 0x39, 0x2f, 0xb3, 0x14, 0xea, 0x2e, 0x10, 0x13,
	0x64, 0x2b, 0xd4, 0x68, 0x26, 0x88, 0xad, 0x10, 0x9e, 0x92, 0x09, 0x17, 0x07, 0xcc, 0xdf, 0x58,
	0xbd, 0x23, 0xc2, 0xc5, 0x5a, 0x92, 0x99, 0x9b, 0x5a, 0x0c, 0xd5, 0x06, 0xe1, 0x28, 0x35, 0x32,
	0x72, 0x71, 0x04, 0x13, 0xd0, 0x96, 0x9c, 0x58, 0x94, 0x52, 0x0c, 0x0e, 0x06, 0xde, 0x20, 0x08,
	0x07, 0xe6, 0x2c, 0x66, 0x6c, 0xce, 0x62, 0x41, 0x76, 0x16, 0x11, 0xde, 0x75, 0xd2, 0xb9, 0xf0,
	0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x0f, 0x0f, 0xe5, 0x18, 0x1b, 0x1e, 0xc9, 0x31, 0xae,
	0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0xbe, 0x78, 0x24, 0xc7, 0xf0, 0xe1, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0xe0,
	0x68, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xcb, 0x97, 0x6e, 0x07, 0x02, 0x00, 0x00,
}
