// Code generated by protoc-gen-gogo.
// source: code.proto
// DO NOT EDIT!

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ErrCode int32

const (
	OK                     ErrCode = 0
	NotEnoughDiamond       ErrCode = 1
	NotEnoughCoin          ErrCode = 2
	NotInRoom              ErrCode = 3
	UsernameOrPwdError     ErrCode = 4
	PhoneNumberError       ErrCode = 5
	LoginError             ErrCode = 6
	UsernameEmpty          ErrCode = 7
	NameTooLong            ErrCode = 8
	PhoneNumberEnpty       ErrCode = 9
	PwdEmpty               ErrCode = 10
	PwdFormatError         ErrCode = 11
	PhoneRegisted          ErrCode = 12
	RegistError            ErrCode = 13
	UserDataNotExist       ErrCode = 14
	WechatLoingFailReAuth  ErrCode = 15
	GetWechatUserInfoFail  ErrCode = 16
	PayOrderFail           ErrCode = 17
	PayOrderError          ErrCode = 18
	RoomNotExist           ErrCode = 19
	RoomFull               ErrCode = 20
	CreateRoomFail         ErrCode = 21
	OperateError           ErrCode = 22
	NiuCardError           ErrCode = 23
	NiuValueError          ErrCode = 24
	BetValueError          ErrCode = 25
	GameStarted            ErrCode = 26
	NotInRoomCannotLeave   ErrCode = 27
	GameStartedCannotLeave ErrCode = 28
	StartedNotKick         ErrCode = 29
	RunningNotVote         ErrCode = 30
	VotingCantLaunchVote   ErrCode = 31
	NotVoteTime            ErrCode = 32
	NotInPrivateRoom       ErrCode = 33
	OtherLoginThisAccount  ErrCode = 34
	BeDealerNotEnough      ErrCode = 35
	SitNotEnough           ErrCode = 36
	SitDownFailed          ErrCode = 37
	BetDealerFailed        ErrCode = 38
	BetNotSeat             ErrCode = 39
	BetTopLimit            ErrCode = 40
	GameNotStart           ErrCode = 41
	StandUpFailed          ErrCode = 42
	DealerSitFailed        ErrCode = 43
	BeDealerAlreadySit     ErrCode = 44
	BeDealerAlready        ErrCode = 45
	DepositNumberError     ErrCode = 46
	DrawMoneyNumberError   ErrCode = 47
	GiveNumberError        ErrCode = 48
	GiveUseridError        ErrCode = 49
	GiveTooMuch            ErrCode = 50
	NotBankrupt            ErrCode = 51
	NotRelieves            ErrCode = 52
	NotPrizeDraw           ErrCode = 53
	NotGotPrizeDraw        ErrCode = 54
	BoxNotYet              ErrCode = 55
	NotBox                 ErrCode = 56
	NotTimes               ErrCode = 57
	AppleOrderFail         ErrCode = 58
	MatchClassicFail       ErrCode = 59
	EnterClassicNotEnough  ErrCode = 60
	NotWinning             ErrCode = 61
	AlreadyWinning         ErrCode = 62
	NotVip                 ErrCode = 63
	NotVipTimes            ErrCode = 64
	AlreadyInRoom          ErrCode = 65
	NotYourTurn            ErrCode = 66
	ErrorOperateValue      ErrCode = 67
)

var ErrCode_name = map[int32]string{
	0:  "OK",
	1:  "NotEnoughDiamond",
	2:  "NotEnoughCoin",
	3:  "NotInRoom",
	4:  "UsernameOrPwdError",
	5:  "PhoneNumberError",
	6:  "LoginError",
	7:  "UsernameEmpty",
	8:  "NameTooLong",
	9:  "PhoneNumberEnpty",
	10: "PwdEmpty",
	11: "PwdFormatError",
	12: "PhoneRegisted",
	13: "RegistError",
	14: "UserDataNotExist",
	15: "WechatLoingFailReAuth",
	16: "GetWechatUserInfoFail",
	17: "PayOrderFail",
	18: "PayOrderError",
	19: "RoomNotExist",
	20: "RoomFull",
	21: "CreateRoomFail",
	22: "OperateError",
	23: "NiuCardError",
	24: "NiuValueError",
	25: "BetValueError",
	26: "GameStarted",
	27: "NotInRoomCannotLeave",
	28: "GameStartedCannotLeave",
	29: "StartedNotKick",
	30: "RunningNotVote",
	31: "VotingCantLaunchVote",
	32: "NotVoteTime",
	33: "NotInPrivateRoom",
	34: "OtherLoginThisAccount",
	35: "BeDealerNotEnough",
	36: "SitNotEnough",
	37: "SitDownFailed",
	38: "BetDealerFailed",
	39: "BetNotSeat",
	40: "BetTopLimit",
	41: "GameNotStart",
	42: "StandUpFailed",
	43: "DealerSitFailed",
	44: "BeDealerAlreadySit",
	45: "BeDealerAlready",
	46: "DepositNumberError",
	47: "DrawMoneyNumberError",
	48: "GiveNumberError",
	49: "GiveUseridError",
	50: "GiveTooMuch",
	51: "NotBankrupt",
	52: "NotRelieves",
	53: "NotPrizeDraw",
	54: "NotGotPrizeDraw",
	55: "BoxNotYet",
	56: "NotBox",
	57: "NotTimes",
	58: "AppleOrderFail",
	59: "MatchClassicFail",
	60: "EnterClassicNotEnough",
	61: "NotWinning",
	62: "AlreadyWinning",
	63: "NotVip",
	64: "NotVipTimes",
	65: "AlreadyInRoom",
	66: "NotYourTurn",
	67: "ErrorOperateValue",
}
var ErrCode_value = map[string]int32{
	"OK":                     0,
	"NotEnoughDiamond":       1,
	"NotEnoughCoin":          2,
	"NotInRoom":              3,
	"UsernameOrPwdError":     4,
	"PhoneNumberError":       5,
	"LoginError":             6,
	"UsernameEmpty":          7,
	"NameTooLong":            8,
	"PhoneNumberEnpty":       9,
	"PwdEmpty":               10,
	"PwdFormatError":         11,
	"PhoneRegisted":          12,
	"RegistError":            13,
	"UserDataNotExist":       14,
	"WechatLoingFailReAuth":  15,
	"GetWechatUserInfoFail":  16,
	"PayOrderFail":           17,
	"PayOrderError":          18,
	"RoomNotExist":           19,
	"RoomFull":               20,
	"CreateRoomFail":         21,
	"OperateError":           22,
	"NiuCardError":           23,
	"NiuValueError":          24,
	"BetValueError":          25,
	"GameStarted":            26,
	"NotInRoomCannotLeave":   27,
	"GameStartedCannotLeave": 28,
	"StartedNotKick":         29,
	"RunningNotVote":         30,
	"VotingCantLaunchVote":   31,
	"NotVoteTime":            32,
	"NotInPrivateRoom":       33,
	"OtherLoginThisAccount":  34,
	"BeDealerNotEnough":      35,
	"SitNotEnough":           36,
	"SitDownFailed":          37,
	"BetDealerFailed":        38,
	"BetNotSeat":             39,
	"BetTopLimit":            40,
	"GameNotStart":           41,
	"StandUpFailed":          42,
	"DealerSitFailed":        43,
	"BeDealerAlreadySit":     44,
	"BeDealerAlready":        45,
	"DepositNumberError":     46,
	"DrawMoneyNumberError":   47,
	"GiveNumberError":        48,
	"GiveUseridError":        49,
	"GiveTooMuch":            50,
	"NotBankrupt":            51,
	"NotRelieves":            52,
	"NotPrizeDraw":           53,
	"NotGotPrizeDraw":        54,
	"BoxNotYet":              55,
	"NotBox":                 56,
	"NotTimes":               57,
	"AppleOrderFail":         58,
	"MatchClassicFail":       59,
	"EnterClassicNotEnough":  60,
	"NotWinning":             61,
	"AlreadyWinning":         62,
	"NotVip":                 63,
	"NotVipTimes":            64,
	"AlreadyInRoom":          65,
	"NotYourTurn":            66,
	"ErrorOperateValue":      67,
}

func (ErrCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorCode, []int{0} }

func init() {
	proto.RegisterEnum("pb.ErrCode", ErrCode_name, ErrCode_value)
}
func (x ErrCode) String() string {
	s, ok := ErrCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("code.proto", fileDescriptorCode) }

var fileDescriptorCode = []byte{
	// 843 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x54, 0xcb, 0x76, 0x1b, 0x45,
	0x10, 0x95, 0x02, 0x38, 0x49, 0xc7, 0x8f, 0x72, 0xc7, 0x36, 0x49, 0x80, 0xe1, 0xfd, 0x0a, 0x21,
	0x3c, 0xc2, 0xfb, 0xad, 0x97, 0x7d, 0x7c, 0x22, 0x8f, 0x74, 0x24, 0xc5, 0x39, 0x59, 0xb6, 0x35,
	0x85, 0xd4, 0x27, 0x9a, 0xae, 0x39, 0xad, 0x1a, 0xd9, 0x66, 0xc5, 0x27, 0xf0, 0x19, 0x7c, 0x03,
	0x5f, 0xc0, 0x32, 0x4b, 0x96, 0x58, 0x6c, 0x58, 0xe6, 0x13, 0x38, 0x35, 0xd3, 0xb2, 0x15, 0x96,
	0x73, 0xbb, 0xea, 0x76, 0xf5, 0xbd, 0xb7, 0x46, 0xa9, 0x21, 0x25, 0x78, 0x37, 0xf3, 0xc4, 0xa4,
	0x2f, 0x65, 0x47, 0xb7, 0xff, 0x50, 0xea, 0x72, 0xcb, 0xfb, 0x06, 0x25, 0xa8, 0x57, 0xd4, 0xa5,
	0xce, 0x7d, 0xa8, 0xe8, 0x2d, 0x05, 0x31, 0x71, 0xcb, 0x51, 0x3e, 0x1a, 0x37, 0xad, 0x49, 0xc9,
	0x25, 0x50, 0xd5, 0x9b, 0x6a, 0xed, 0x1c, 0x6d, 0x90, 0x75, 0x70, 0x49, 0xaf, 0xa9, 0xab, 0x31,
	0xf1, 0xbe, 0xeb, 0x11, 0xa5, 0xf0, 0x9c, 0xde, 0x51, 0xfa, 0xc1, 0x14, 0xbd, 0x33, 0x29, 0x76,
	0x7c, 0xf7, 0x38, 0x69, 0x79, 0x4f, 0x1e, 0x9e, 0x17, 0xbe, 0xee, 0x98, 0x1c, 0xc6, 0x79, 0x7a,
	0x84, 0xbe, 0x44, 0x5f, 0xd0, 0xeb, 0x4a, 0xb5, 0x69, 0x64, 0x5d, 0xf9, 0xbd, 0x22, 0xfc, 0x8b,
	0xee, 0x56, 0x9a, 0xf1, 0x29, 0x5c, 0xd6, 0x1b, 0xea, 0x5a, 0x6c, 0x52, 0x1c, 0x10, 0xb5, 0xc9,
	0x8d, 0xe0, 0xca, 0xff, 0x99, 0x9c, 0x94, 0x5d, 0xd5, 0xab, 0xea, 0x8a, 0xdc, 0x56, 0x34, 0x29,
	0xad, 0xd5, 0x7a, 0xf7, 0x38, 0xd9, 0x25, 0x9f, 0x1a, 0x2e, 0xb9, 0xaf, 0x09, 0x77, 0xd1, 0xd7,
	0xc3, 0x91, 0x9d, 0x32, 0x26, 0xb0, 0x2a, 0xdc, 0xe5, 0x57, 0x59, 0xb3, 0x26, 0xdc, 0x72, 0x7f,
	0xd3, 0xb0, 0x91, 0x77, 0x9e, 0xd8, 0x29, 0xc3, 0xba, 0xbe, 0xa9, 0xb6, 0x1f, 0xe2, 0x70, 0x6c,
	0xb8, 0x4d, 0xd6, 0x8d, 0x76, 0x8d, 0x9d, 0xf4, 0xb0, 0x96, 0xf3, 0x18, 0x36, 0xe4, 0x68, 0x0f,
	0xb9, 0x3c, 0x95, 0xce, 0x7d, 0xf7, 0x33, 0x49, 0x01, 0x80, 0x06, 0xb5, 0xda, 0x35, 0xa7, 0x1d,
	0x9f, 0xa0, 0x2f, 0x90, 0xcd, 0x62, 0x82, 0x80, 0x94, 0x17, 0x6a, 0x29, 0x12, 0xe1, 0xce, 0x2f,
	0xbb, 0x2e, 0x0f, 0x11, 0x64, 0x37, 0x9f, 0x4c, 0x60, 0x4b, 0x1e, 0xd2, 0xf0, 0x68, 0x18, 0x0b,
	0x4c, 0x68, 0xb6, 0xa5, 0xa7, 0x93, 0xa1, 0x37, 0x8c, 0x25, 0xcb, 0x8e, 0x20, 0xb1, 0xcd, 0x1b,
	0xc6, 0x07, 0xb9, 0x5f, 0x2c, 0x8c, 0xb2, 0xf9, 0xa1, 0x99, 0xe4, 0xa1, 0xe8, 0x86, 0x40, 0x75,
	0xe4, 0x25, 0xe8, 0xa6, 0xbc, 0x7f, 0xcf, 0xa4, 0xd8, 0x67, 0xe3, 0x45, 0x90, 0x5b, 0xfa, 0x86,
	0xda, 0x3a, 0x37, 0xb3, 0x61, 0x9c, 0x23, 0x6e, 0xa3, 0x99, 0x21, 0xbc, 0xa4, 0x6f, 0xa9, 0x9d,
	0xa5, 0xd2, 0xe5, 0xb3, 0x97, 0x65, 0xc8, 0x80, 0xc7, 0xc4, 0xf7, 0xed, 0xf0, 0x31, 0xbc, 0x22,
	0x58, 0x2f, 0x77, 0xce, 0xba, 0x51, 0x4c, 0x7c, 0x48, 0x8c, 0x10, 0x09, 0xfb, 0x21, 0xb1, 0x75,
	0xa3, 0x86, 0x71, 0xdc, 0x36, 0xb9, 0x1b, 0x8e, 0x8b, 0x93, 0x57, 0x0b, 0x93, 0xcb, 0xb2, 0x81,
	0x4d, 0x11, 0x5e, 0x0b, 0xf1, 0xdb, 0x77, 0x5d, 0x6f, 0x67, 0xe1, 0xf5, 0xf0, 0xba, 0xa8, 0xdd,
	0xe1, 0x31, 0xfa, 0x22, 0x33, 0x83, 0xb1, 0x9d, 0xd6, 0x86, 0x43, 0xca, 0x1d, 0xc3, 0x1b, 0x7a,
	0x5b, 0x6d, 0xd6, 0xb1, 0x89, 0x66, 0x82, 0xfe, 0x3c, 0xa1, 0xf0, 0xa6, 0x28, 0xd3, 0xb7, 0x7c,
	0x81, 0xbc, 0x25, 0x32, 0xf4, 0x2d, 0x37, 0xe9, 0xd8, 0x89, 0x9c, 0x98, 0xc0, 0xdb, 0xfa, 0xba,
	0xda, 0xa8, 0x23, 0x97, 0xcd, 0x01, 0x7c, 0x47, 0xa2, 0x59, 0x47, 0xe9, 0xec, 0xa3, 0x61, 0x78,
	0x57, 0x46, 0xac, 0x23, 0x0f, 0x28, 0x6b, 0xdb, 0xd4, 0x32, 0xbc, 0x27, 0xd4, 0xa2, 0x88, 0x54,
	0xc8, 0xe3, 0xe1, 0xfd, 0x82, 0x9a, 0x8d, 0x4b, 0x1e, 0x64, 0x81, 0xe5, 0xb6, 0x50, 0x97, 0xbc,
	0x7d, 0xcb, 0x01, 0xfc, 0x40, 0x76, 0x64, 0x31, 0x6b, 0x6d, 0xe2, 0xd1, 0x24, 0xa7, 0x7d, 0xcb,
	0x70, 0xa7, 0x9c, 0xe3, 0x19, 0x1c, 0x3e, 0x94, 0xe2, 0x26, 0x66, 0x34, 0xb5, 0xbc, 0xbc, 0x3a,
	0x77, 0x45, 0xcc, 0xa6, 0x37, 0xc7, 0x07, 0xe4, 0xf0, 0x74, 0xf9, 0xe4, 0x23, 0xa1, 0xd9, 0xb3,
	0xb3, 0x67, 0x36, 0xed, 0xe3, 0x05, 0x28, 0x19, 0xb5, 0x21, 0x25, 0x9f, 0x14, 0xfe, 0xdb, 0x99,
	0xec, 0xd6, 0x41, 0x3e, 0x1c, 0xc3, 0xa7, 0xc1, 0x87, 0xba, 0x71, 0x8f, 0x7d, 0x9e, 0x31, 0xdc,
	0x0b, 0x40, 0x0f, 0x27, 0x16, 0x67, 0x38, 0x85, 0xcf, 0x8a, 0xa8, 0x11, 0x77, 0xbd, 0xfd, 0x05,
	0xe5, 0x7a, 0xf8, 0x5c, 0x98, 0x63, 0xe2, 0xbd, 0x65, 0xf0, 0x0b, 0xf9, 0x2b, 0xd4, 0xe9, 0x24,
	0x26, 0x7e, 0x84, 0x0c, 0x5f, 0x6a, 0xa5, 0x56, 0x84, 0x97, 0x4e, 0xe0, 0x2b, 0x09, 0x78, 0x4c,
	0x2c, 0x3e, 0x4f, 0xe1, 0x6b, 0xc9, 0x49, 0x2d, 0xcb, 0x26, 0x78, 0xb1, 0x27, 0xdf, 0x88, 0xf9,
	0x07, 0x86, 0x87, 0xe3, 0xc6, 0xc4, 0x4c, 0xa7, 0x76, 0x58, 0xa0, 0xdf, 0x8a, 0xf9, 0x2d, 0xc7,
	0xe8, 0x03, 0x7a, 0xe1, 0xe9, 0x77, 0xe2, 0x55, 0x4c, 0xfc, 0xd0, 0x16, 0x79, 0x83, 0xef, 0x0b,
	0xd2, 0x52, 0xc0, 0x05, 0xf6, 0x43, 0x18, 0xe1, 0xd0, 0x66, 0xf0, 0xe3, 0x22, 0x6e, 0x36, 0x2b,
	0xa7, 0xf8, 0x49, 0x9c, 0x0b, 0x0d, 0xe1, 0x47, 0x56, 0x0b, 0x35, 0x8f, 0x28, 0xf7, 0x83, 0xdc,
	0x3b, 0xa8, 0x4b, 0xc2, 0x0a, 0xdd, 0xc2, 0xee, 0x15, 0x8b, 0x04, 0x8d, 0xfa, 0x9d, 0x27, 0x67,
	0x51, 0xe5, 0xaf, 0xb3, 0xa8, 0xf2, 0xf4, 0x2c, 0xaa, 0xfe, 0x3a, 0x8f, 0xaa, 0xbf, 0xcf, 0xa3,
	0xea, 0x9f, 0xf3, 0xa8, 0xfa, 0x64, 0x1e, 0x55, 0xff, 0x9e, 0x47, 0xd5, 0x7f, 0xe7, 0x51, 0xe5,
	0xe9, 0x3c, 0xaa, 0xfe, 0xf6, 0x4f, 0x54, 0x39, 0x5a, 0x29, 0xfe, 0xba, 0xf7, 0xfe, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0x4d, 0x64, 0x76, 0x60, 0x83, 0x05, 0x00, 0x00,
}