// Code generated by protoc-gen-gogo.
// source: web.proto
// DO NOT EDIT!

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConfigType int32

const (
	CONFIG_OK      ConfigType = 0
	CONFIG_BOX     ConfigType = 1
	CONFIG_ENV     ConfigType = 2
	CONFIG_LOTTERY ConfigType = 3
	CONFIG_NOTICE  ConfigType = 4
	CONFIG_PRIZE   ConfigType = 5
	CONFIG_SHOP    ConfigType = 6
	CONFIG_VIP     ConfigType = 7
	CONFIG_CLASSIC ConfigType = 8
)

var ConfigType_name = map[int32]string{
	0: "CONFIG_OK",
	1: "CONFIG_BOX",
	2: "CONFIG_ENV",
	3: "CONFIG_LOTTERY",
	4: "CONFIG_NOTICE",
	5: "CONFIG_PRIZE",
	6: "CONFIG_SHOP",
	7: "CONFIG_VIP",
	8: "CONFIG_CLASSIC",
}
var ConfigType_value = map[string]int32{
	"CONFIG_OK":      0,
	"CONFIG_BOX":     1,
	"CONFIG_ENV":     2,
	"CONFIG_LOTTERY": 3,
	"CONFIG_NOTICE":  4,
	"CONFIG_PRIZE":   5,
	"CONFIG_SHOP":    6,
	"CONFIG_VIP":     7,
	"CONFIG_CLASSIC": 8,
}

func (ConfigType) EnumDescriptor() ([]byte, []int) { return fileDescriptorWeb, []int{0} }

type WebCode int32

const (
	WebReqMsg   WebCode = 0
	WebOnline   WebCode = 1
	WebNotice   WebCode = 2
	WebBuild    WebCode = 3
	WebGive     WebCode = 4
	WebSetHands WebCode = 5
	WebGetHands WebCode = 6
	WebShop     WebCode = 7
	WebSetEnv   WebCode = 8
	WebGetEnv   WebCode = 9
	WebDelEnv   WebCode = 10
	WebPrize    WebCode = 11
	WebBox      WebCode = 12
	WebRoom     WebCode = 13
	WebClassic  WebCode = 14
	WebVip      WebCode = 15
	WebDan      WebCode = 16
	WebTask     WebCode = 17
	WebBetting  WebCode = 18
	WebLottery  WebCode = 19
	WebChange   WebCode = 20
)

var WebCode_name = map[int32]string{
	0:  "WebReqMsg",
	1:  "WebOnline",
	2:  "WebNotice",
	3:  "WebBuild",
	4:  "WebGive",
	5:  "WebSetHands",
	6:  "WebGetHands",
	7:  "WebShop",
	8:  "WebSetEnv",
	9:  "WebGetEnv",
	10: "WebDelEnv",
	11: "WebPrize",
	12: "WebBox",
	13: "WebRoom",
	14: "WebClassic",
	15: "WebVip",
	16: "WebDan",
	17: "WebTask",
	18: "WebBetting",
	19: "WebLottery",
	20: "WebChange",
}
var WebCode_value = map[string]int32{
	"WebReqMsg":   0,
	"WebOnline":   1,
	"WebNotice":   2,
	"WebBuild":    3,
	"WebGive":     4,
	"WebSetHands": 5,
	"WebGetHands": 6,
	"WebShop":     7,
	"WebSetEnv":   8,
	"WebGetEnv":   9,
	"WebDelEnv":   10,
	"WebPrize":    11,
	"WebBox":      12,
	"WebRoom":     13,
	"WebClassic":  14,
	"WebVip":      15,
	"WebDan":      16,
	"WebTask":     17,
	"WebBetting":  18,
	"WebLottery":  19,
	"WebChange":   20,
}

func (WebCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorWeb, []int{1} }

// 同步数据
type SyncConfig struct {
	Type ConfigType `protobuf:"varint,1,opt,name=Type,proto3,enum=pb.ConfigType" json:"Type,omitempty"`
	Data []byte     `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *SyncConfig) Reset()                    { *m = SyncConfig{} }
func (*SyncConfig) ProtoMessage()               {}
func (*SyncConfig) Descriptor() ([]byte, []int) { return fileDescriptorWeb, []int{0} }

func (m *SyncConfig) GetType() ConfigType {
	if m != nil {
		return m.Type
	}
	return CONFIG_OK
}

func (m *SyncConfig) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 主动获取数据
type GetConfig struct {
	Type ConfigType `protobuf:"varint,1,opt,name=Type,proto3,enum=pb.ConfigType" json:"Type,omitempty"`
}

func (m *GetConfig) Reset()                    { *m = GetConfig{} }
func (*GetConfig) ProtoMessage()               {}
func (*GetConfig) Descriptor() ([]byte, []int) { return fileDescriptorWeb, []int{1} }

func (m *GetConfig) GetType() ConfigType {
	if m != nil {
		return m.Type
	}
	return CONFIG_OK
}

// web请求
type WebRequest struct {
	Code WebCode `protobuf:"varint,1,opt,name=Code,proto3,enum=pb.WebCode" json:"Code,omitempty"`
	Data []byte  `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *WebRequest) Reset()                    { *m = WebRequest{} }
func (*WebRequest) ProtoMessage()               {}
func (*WebRequest) Descriptor() ([]byte, []int) { return fileDescriptorWeb, []int{2} }

func (m *WebRequest) GetCode() WebCode {
	if m != nil {
		return m.Code
	}
	return WebReqMsg
}

func (m *WebRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WebResponse struct {
	Code    WebCode `protobuf:"varint,1,opt,name=Code,proto3,enum=pb.WebCode" json:"Code,omitempty"`
	ErrCode int32   `protobuf:"varint,2,opt,name=ErrCode,proto3" json:"ErrCode,omitempty"`
	ErrMsg  string  `protobuf:"bytes,3,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Result  []byte  `protobuf:"bytes,4,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (m *WebResponse) Reset()                    { *m = WebResponse{} }
func (*WebResponse) ProtoMessage()               {}
func (*WebResponse) Descriptor() ([]byte, []int) { return fileDescriptorWeb, []int{3} }

func (m *WebResponse) GetCode() WebCode {
	if m != nil {
		return m.Code
	}
	return WebReqMsg
}

func (m *WebResponse) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *WebResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *WebResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncConfig)(nil), "pb.SyncConfig")
	proto.RegisterType((*GetConfig)(nil), "pb.GetConfig")
	proto.RegisterType((*WebRequest)(nil), "pb.WebRequest")
	proto.RegisterType((*WebResponse)(nil), "pb.WebResponse")
	proto.RegisterEnum("pb.ConfigType", ConfigType_name, ConfigType_value)
	proto.RegisterEnum("pb.WebCode", WebCode_name, WebCode_value)
}
func (x ConfigType) String() string {
	s, ok := ConfigType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x WebCode) String() string {
	s, ok := WebCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *SyncConfig) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SyncConfig)
	if !ok {
		that2, ok := that.(SyncConfig)
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
	if this.Type != that1.Type {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	return true
}
func (this *GetConfig) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*GetConfig)
	if !ok {
		that2, ok := that.(GetConfig)
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
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (this *WebRequest) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*WebRequest)
	if !ok {
		that2, ok := that.(WebRequest)
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
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	return true
}
func (this *WebResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*WebResponse)
	if !ok {
		that2, ok := that.(WebResponse)
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
	if this.ErrCode != that1.ErrCode {
		return false
	}
	if this.ErrMsg != that1.ErrMsg {
		return false
	}
	if !bytes.Equal(this.Result, that1.Result) {
		return false
	}
	return true
}
func (this *SyncConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.SyncConfig{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&pb.GetConfig{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WebRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&pb.WebRequest{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WebResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&pb.WebResponse{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "ErrCode: "+fmt.Sprintf("%#v", this.ErrCode)+",\n")
	s = append(s, "ErrMsg: "+fmt.Sprintf("%#v", this.ErrMsg)+",\n")
	s = append(s, "Result: "+fmt.Sprintf("%#v", this.Result)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringWeb(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *SyncConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWeb(dAtA, i, uint64(m.Type))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWeb(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *GetConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWeb(dAtA, i, uint64(m.Type))
	}
	return i, nil
}

func (m *WebRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWeb(dAtA, i, uint64(m.Code))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintWeb(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	return i, nil
}

func (m *WebResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WebResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintWeb(dAtA, i, uint64(m.Code))
	}
	if m.ErrCode != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintWeb(dAtA, i, uint64(m.ErrCode))
	}
	if len(m.ErrMsg) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintWeb(dAtA, i, uint64(len(m.ErrMsg)))
		i += copy(dAtA[i:], m.ErrMsg)
	}
	if len(m.Result) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintWeb(dAtA, i, uint64(len(m.Result)))
		i += copy(dAtA[i:], m.Result)
	}
	return i, nil
}

func encodeFixed64Web(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Web(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintWeb(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SyncConfig) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovWeb(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovWeb(uint64(l))
	}
	return n
}

func (m *GetConfig) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovWeb(uint64(m.Type))
	}
	return n
}

func (m *WebRequest) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovWeb(uint64(m.Code))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovWeb(uint64(l))
	}
	return n
}

func (m *WebResponse) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovWeb(uint64(m.Code))
	}
	if m.ErrCode != 0 {
		n += 1 + sovWeb(uint64(m.ErrCode))
	}
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovWeb(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovWeb(uint64(l))
	}
	return n
}

func sovWeb(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozWeb(x uint64) (n int) {
	return sovWeb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *SyncConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&SyncConfig{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetConfig{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WebRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WebRequest{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WebResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WebResponse{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`ErrCode:` + fmt.Sprintf("%v", this.ErrCode) + `,`,
		`ErrMsg:` + fmt.Sprintf("%v", this.ErrMsg) + `,`,
		`Result:` + fmt.Sprintf("%v", this.Result) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringWeb(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *SyncConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWeb
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
			return fmt.Errorf("proto: SyncConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (ConfigType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWeb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWeb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWeb
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
func (m *GetConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWeb
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
			return fmt.Errorf("proto: GetConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (ConfigType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWeb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWeb
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
func (m *WebRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWeb
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
			return fmt.Errorf("proto: WebRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (WebCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWeb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWeb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWeb
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
func (m *WebResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWeb
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
			return fmt.Errorf("proto: WebResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WebResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (WebCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrCode", wireType)
			}
			m.ErrCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ErrCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWeb
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWeb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthWeb
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWeb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWeb
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
func skipWeb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWeb
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
					return 0, ErrIntOverflowWeb
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
					return 0, ErrIntOverflowWeb
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
				return 0, ErrInvalidLengthWeb
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWeb
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
				next, err := skipWeb(dAtA[start:])
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
	ErrInvalidLengthWeb = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWeb   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("web.proto", fileDescriptorWeb) }

var fileDescriptorWeb = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x8e, 0xd2, 0x50,
	0x14, 0xc6, 0x7b, 0xa1, 0xc3, 0x9f, 0xc3, 0x9f, 0x39, 0x73, 0x35, 0x86, 0x55, 0x25, 0xac, 0xc8,
	0xc4, 0x60, 0xa2, 0x4f, 0x30, 0x94, 0xca, 0x10, 0x19, 0x4a, 0x5a, 0x02, 0xea, 0xc6, 0xb4, 0x70,
	0x65, 0x1a, 0xb1, 0xb7, 0xb6, 0x97, 0x71, 0x70, 0xe5, 0x23, 0xf8, 0x18, 0xfa, 0x06, 0x3e, 0x82,
	0xcb, 0x59, 0xba, 0x94, 0xba, 0x71, 0x39, 0x8f, 0x60, 0x6e, 0xb9, 0x44, 0x16, 0x26, 0xce, 0xae,
	0xbf, 0xef, 0x9c, 0xf3, 0x9d, 0xef, 0xdc, 0xa4, 0x50, 0xfe, 0xc0, 0xfc, 0x4e, 0x14, 0x73, 0xc1,
	0x69, 0x2e, 0xf2, 0x5b, 0x3d, 0x00, 0x77, 0x13, 0xce, 0x4d, 0x1e, 0xbe, 0x09, 0x96, 0xb4, 0x05,
	0xfa, 0x64, 0x13, 0xb1, 0x06, 0x69, 0x92, 0x76, 0xfd, 0x49, 0xbd, 0x13, 0xf9, 0x9d, 0x5d, 0x45,
	0xaa, 0x4e, 0x56, 0xa3, 0x14, 0xf4, 0x9e, 0x27, 0xbc, 0x46, 0xae, 0x49, 0xda, 0x55, 0x27, 0xfb,
	0x6e, 0x3d, 0x86, 0x72, 0x9f, 0x89, 0xbb, 0x9b, 0xb4, 0xce, 0x00, 0x66, 0xcc, 0x77, 0xd8, 0xfb,
	0x35, 0x4b, 0x04, 0x7d, 0x08, 0xba, 0xc9, 0x17, 0xfb, 0x89, 0x8a, 0x9c, 0x98, 0x31, 0x5f, 0x4a,
	0x4e, 0x56, 0xf8, 0xe7, 0xce, 0x6b, 0xa8, 0x64, 0x16, 0x49, 0xc4, 0xc3, 0x84, 0xfd, 0xdf, 0xa3,
	0x01, 0x45, 0x2b, 0x8e, 0xb3, 0x1e, 0x69, 0x73, 0xe4, 0xec, 0x91, 0x3e, 0x80, 0x82, 0x15, 0xc7,
	0x17, 0xc9, 0xb2, 0x91, 0x6f, 0x92, 0x76, 0xd9, 0x51, 0x24, 0x75, 0x87, 0x25, 0xeb, 0x95, 0x68,
	0xe8, 0xd9, 0x5e, 0x45, 0xa7, 0x5f, 0x09, 0xc0, 0xdf, 0x8b, 0x68, 0x0d, 0xca, 0xa6, 0x3d, 0x7a,
	0x36, 0xe8, 0xbf, 0xb6, 0x9f, 0xa3, 0x46, 0xeb, 0x00, 0x0a, 0xbb, 0xf6, 0x0b, 0x24, 0x07, 0x6c,
	0x8d, 0xa6, 0x98, 0xa3, 0x14, 0xea, 0x8a, 0x87, 0xf6, 0x64, 0x62, 0x39, 0x2f, 0x31, 0x4f, 0x4f,
	0xa0, 0xa6, 0xb4, 0x91, 0x3d, 0x19, 0x98, 0x16, 0xea, 0x14, 0xa1, 0xaa, 0xa4, 0xb1, 0x33, 0x78,
	0x65, 0xe1, 0x11, 0x3d, 0x86, 0x8a, 0x52, 0xdc, 0x73, 0x7b, 0x8c, 0x85, 0x03, 0xe7, 0xe9, 0x60,
	0x8c, 0xc5, 0x03, 0x67, 0x73, 0x78, 0xe6, 0xba, 0x03, 0x13, 0x4b, 0xa7, 0xdf, 0x72, 0x50, 0x54,
	0xef, 0x20, 0x83, 0xee, 0x1e, 0xfd, 0x22, 0x59, 0xa2, 0xa6, 0xd0, 0x0e, 0x57, 0x41, 0xc8, 0x90,
	0x28, 0x1c, 0x71, 0x11, 0xcc, 0x19, 0xe6, 0x68, 0x15, 0x4a, 0x33, 0xe6, 0x77, 0xd7, 0xc1, 0x6a,
	0x81, 0x79, 0x5a, 0xc9, 0x5c, 0xfa, 0xc1, 0x15, 0x43, 0x5d, 0x06, 0x99, 0x31, 0xdf, 0x65, 0xe2,
	0xdc, 0x0b, 0x17, 0xc9, 0x2e, 0x99, 0xac, 0xee, 0x85, 0x82, 0x6a, 0x77, 0x2f, 0x79, 0x84, 0x45,
	0x65, 0xec, 0x32, 0x61, 0x85, 0x57, 0x58, 0x52, 0xd8, 0xdf, 0x61, 0x59, 0x61, 0x8f, 0xad, 0x24,
	0x82, 0x5a, 0x3b, 0x8e, 0x83, 0x8f, 0x0c, 0x2b, 0x14, 0xa0, 0x20, 0x43, 0xf0, 0x6b, 0xac, 0x2a,
	0x4f, 0x87, 0xf3, 0x77, 0x58, 0x93, 0xa7, 0xcb, 0xab, 0x56, 0x5e, 0x92, 0x04, 0x73, 0xac, 0xab,
	0xc6, 0x69, 0x10, 0xe1, 0xb1, 0xfa, 0xee, 0x79, 0x21, 0xa2, 0x1a, 0x9a, 0x78, 0xc9, 0x5b, 0x3c,
	0x51, 0x43, 0x5d, 0x26, 0x44, 0x10, 0x2e, 0x91, 0x2a, 0x1e, 0x72, 0x21, 0x58, 0xbc, 0xc1, 0x7b,
	0x2a, 0x8a, 0x79, 0xe9, 0x85, 0x4b, 0x86, 0xf7, 0xbb, 0x8f, 0x6e, 0xb6, 0x86, 0xf6, 0x63, 0x6b,
	0x68, 0xb7, 0x5b, 0x83, 0x7c, 0x4a, 0x0d, 0xf2, 0x25, 0x35, 0xc8, 0xf7, 0xd4, 0x20, 0x37, 0xa9,
	0x41, 0x7e, 0xa6, 0x06, 0xf9, 0x9d, 0x1a, 0xda, 0x6d, 0x6a, 0x90, 0xcf, 0xbf, 0x0c, 0xcd, 0x2f,
	0x64, 0xff, 0xd4, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xd8, 0xdf, 0x36, 0x60, 0x03,
	0x00, 0x00,
}
