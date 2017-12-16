// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//打包消息
func Rpacket(msg interface{}) (uint32, []byte, error) {
	switch msg.(type) {
	case *CConfig:
		msg.(*CConfig).Code = 1020
		b, err := msg.(*CConfig).Marshal()
		return 1020, b, err
	case *CNiu:
		msg.(*CNiu).Code = 4011
		b, err := msg.(*CNiu).Marshal()
		return 4011, b, err
	case *CWxLogin:
		msg.(*CWxLogin).Code = 1004
		b, err := msg.(*CWxLogin).Marshal()
		return 1004, b, err
	case *CBuildAgent:
		msg.(*CBuildAgent).Code = 1026
		b, err := msg.(*CBuildAgent).Marshal()
		return 1026, b, err
	case *CPrizeDraw:
		msg.(*CPrizeDraw).Code = 1054
		b, err := msg.(*CPrizeDraw).Marshal()
		return 1054, b, err
	case *CQualifying:
		msg.(*CQualifying).Code = 1081
		b, err := msg.(*CQualifying).Marshal()
		return 1081, b, err
	case *CBettingRecord:
		msg.(*CBettingRecord).Code = 4205
		b, err := msg.(*CBettingRecord).Marshal()
		return 4205, b, err
	case *CUserData:
		msg.(*CUserData).Code = 1022
		b, err := msg.(*CUserData).Marshal()
		return 1022, b, err
	case *CDanInfo:
		msg.(*CDanInfo).Code = 1080
		b, err := msg.(*CDanInfo).Marshal()
		return 1080, b, err
	case *CDealerList:
		msg.(*CDealerList).Code = 4043
		b, err := msg.(*CDealerList).Marshal()
		return 4043, b, err
	case *CFreeTrend:
		msg.(*CFreeTrend).Code = 4051
		b, err := msg.(*CFreeTrend).Marshal()
		return 4051, b, err
	case *CEnterZiRoom:
		msg.(*CEnterZiRoom).Code = 4100
		b, err := msg.(*CEnterZiRoom).Marshal()
		return 4100, b, err
	case *CShop:
		msg.(*CShop).Code = 3010
		b, err := msg.(*CShop).Marshal()
		return 3010, b, err
	case *CBank:
		msg.(*CBank).Code = 1030
		b, err := msg.(*CBank).Marshal()
		return 1030, b, err
	case *CPrizeBox:
		msg.(*CPrizeBox).Code = 1055
		b, err := msg.(*CPrizeBox).Marshal()
		return 1055, b, err
	case *CKick:
		msg.(*CKick).Code = 4005
		b, err := msg.(*CKick).Marshal()
		return 4005, b, err
	case *CWxpayOrder:
		msg.(*CWxpayOrder).Code = 3002
		b, err := msg.(*CWxpayOrder).Marshal()
		return 3002, b, err
	case *CPing:
		msg.(*CPing).Code = 1050
		b, err := msg.(*CPing).Marshal()
		return 1050, b, err
	case *CBankrupts:
		msg.(*CBankrupts).Code = 1052
		b, err := msg.(*CBankrupts).Marshal()
		return 1052, b, err
	case *CFreeDealer:
		msg.(*CFreeDealer).Code = 4042
		b, err := msg.(*CFreeDealer).Marshal()
		return 4042, b, err
	case *COperate:
		msg.(*COperate).Code = 4108
		b, err := msg.(*COperate).Marshal()
		return 4108, b, err
	case *CWxpayQuery:
		msg.(*CWxpayQuery).Code = 3004
		b, err := msg.(*CWxpayQuery).Marshal()
		return 3004, b, err
	case *CChatText:
		msg.(*CChatText).Code = 2003
		b, err := msg.(*CChatText).Marshal()
		return 2003, b, err
	case *CEnterFreeRoom:
		msg.(*CEnterFreeRoom).Code = 4040
		b, err := msg.(*CEnterFreeRoom).Marshal()
		return 4040, b, err
	case *CCreateZiRoom:
		msg.(*CCreateZiRoom).Code = 4101
		b, err := msg.(*CCreateZiRoom).Marshal()
		return 4101, b, err
	case *CFreeBet:
		msg.(*CFreeBet).Code = 4046
		b, err := msg.(*CFreeBet).Marshal()
		return 4046, b, err
	case *CLogin:
		msg.(*CLogin).Code = 1000
		b, err := msg.(*CLogin).Marshal()
		return 1000, b, err
	case *CMailList:
		msg.(*CMailList).Code = 1071
		b, err := msg.(*CMailList).Marshal()
		return 1071, b, err
	case *CReady:
		msg.(*CReady).Code = 4006
		b, err := msg.(*CReady).Marshal()
		return 4006, b, err
	case *CVote:
		msg.(*CVote).Code = 4017
		b, err := msg.(*CVote).Marshal()
		return 4017, b, err
	case *CApplePay:
		msg.(*CApplePay).Code = 3006
		b, err := msg.(*CApplePay).Marshal()
		return 3006, b, err
	case *CPrizeList:
		msg.(*CPrizeList).Code = 1053
		b, err := msg.(*CPrizeList).Marshal()
		return 1053, b, err
	case *CVipList:
		msg.(*CVipList).Code = 1062
		b, err := msg.(*CVipList).Marshal()
		return 1062, b, err
	case *CClassicList:
		msg.(*CClassicList).Code = 1060
		b, err := msg.(*CClassicList).Marshal()
		return 1060, b, err
	case *CGetMailItem:
		msg.(*CGetMailItem).Code = 1073
		b, err := msg.(*CGetMailItem).Marshal()
		return 1073, b, err
	case *CCreateRoom:
		msg.(*CCreateRoom).Code = 4001
		b, err := msg.(*CCreateRoom).Marshal()
		return 4001, b, err
	case *CGameRecord:
		msg.(*CGameRecord).Code = 4020
		b, err := msg.(*CGameRecord).Marshal()
		return 4020, b, err
	case *CFreeSit:
		msg.(*CFreeSit).Code = 4044
		b, err := msg.(*CFreeSit).Marshal()
		return 4044, b, err
	case *CZiGameRecord:
		msg.(*CZiGameRecord).Code = 4103
		b, err := msg.(*CZiGameRecord).Marshal()
		return 4103, b, err
	case *CRoomList:
		msg.(*CRoomList).Code = 4115
		b, err := msg.(*CRoomList).Marshal()
		return 4115, b, err
	case *CBuy:
		msg.(*CBuy).Code = 3000
		b, err := msg.(*CBuy).Marshal()
		return 3000, b, err
	case *CChatVoice:
		msg.(*CChatVoice).Code = 2004
		b, err := msg.(*CChatVoice).Marshal()
		return 2004, b, err
	case *CNotice:
		msg.(*CNotice).Code = 2008
		b, err := msg.(*CNotice).Marshal()
		return 2008, b, err
	case *CDanRanking:
		msg.(*CDanRanking).Code = 1082
		b, err := msg.(*CDanRanking).Marshal()
		return 1082, b, err
	case *CEnterRoom:
		msg.(*CEnterRoom).Code = 4000
		b, err := msg.(*CEnterRoom).Marshal()
		return 4000, b, err
	case *CGetCurrency:
		msg.(*CGetCurrency).Code = 1024
		b, err := msg.(*CGetCurrency).Marshal()
		return 1024, b, err
	case *CLotteryInfo:
		msg.(*CLotteryInfo).Code = 4220
		b, err := msg.(*CLotteryInfo).Marshal()
		return 4220, b, err
	case *CPrizeCards:
		msg.(*CPrizeCards).Code = 4082
		b, err := msg.(*CPrizeCards).Marshal()
		return 4082, b, err
	case *CRegist:
		msg.(*CRegist).Code = 1002
		b, err := msg.(*CRegist).Marshal()
		return 1002, b, err
	case *CDealer:
		msg.(*CDealer).Code = 4008
		b, err := msg.(*CDealer).Marshal()
		return 4008, b, err
	case *CLottery:
		msg.(*CLottery).Code = 4221
		b, err := msg.(*CLottery).Marshal()
		return 4221, b, err
	case *CDeleteMail:
		msg.(*CDeleteMail).Code = 1072
		b, err := msg.(*CDeleteMail).Marshal()
		return 1072, b, err
	case *CEnterClassicRoom:
		msg.(*CEnterClassicRoom).Code = 4060
		b, err := msg.(*CEnterClassicRoom).Marshal()
		return 4060, b, err
	case *CGetPrize:
		msg.(*CGetPrize).Code = 4081
		b, err := msg.(*CGetPrize).Marshal()
		return 4081, b, err
	case *CPushDiscard:
		msg.(*CPushDiscard).Code = 4106
		b, err := msg.(*CPushDiscard).Marshal()
		return 4106, b, err
	case *CBetting:
		msg.(*CBetting).Code = 4201
		b, err := msg.(*CBetting).Marshal()
		return 4201, b, err
	case *CLeave:
		msg.(*CLeave).Code = 4004
		b, err := msg.(*CLeave).Marshal()
		return 4004, b, err
	case *CBet:
		msg.(*CBet).Code = 4010
		b, err := msg.(*CBet).Marshal()
		return 4010, b, err
	case *CLaunchVote:
		msg.(*CLaunchVote).Code = 4016
		b, err := msg.(*CLaunchVote).Marshal()
		return 4016, b, err
	case *CBettingInfo:
		msg.(*CBettingInfo).Code = 4200
		b, err := msg.(*CBettingInfo).Marshal()
		return 4200, b, err
	default:
		return 0, []byte{}, errors.New("unknown message")
	}
}