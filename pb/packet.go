// Code generated by tool/gen.go.
// DO NOT EDIT!

package pb

import (
	"errors"
)

//打包消息
func Packet(msg interface{}) (uint32, []byte, error) {
	switch msg.(type) {
	case *SBettingInfo:
		msg.(*SBettingInfo).Code = 4200
		b, err := msg.(*SBettingInfo).Marshal()
		return 4200, b, err
	case *SShop:
		msg.(*SShop).Code = 3010
		b, err := msg.(*SShop).Marshal()
		return 3010, b, err
	case *SMailList:
		msg.(*SMailList).Code = 1071
		b, err := msg.(*SMailList).Marshal()
		return 1071, b, err
	case *SGetMailItem:
		msg.(*SGetMailItem).Code = 1073
		b, err := msg.(*SGetMailItem).Marshal()
		return 1073, b, err
	case *SDanRanking:
		msg.(*SDanRanking).Code = 1082
		b, err := msg.(*SDanRanking).Marshal()
		return 1082, b, err
	case *SPrizeCards:
		msg.(*SPrizeCards).Code = 4082
		b, err := msg.(*SPrizeCards).Marshal()
		return 4082, b, err
	case *SNotice:
		msg.(*SNotice).Code = 2008
		b, err := msg.(*SNotice).Marshal()
		return 2008, b, err
	case *SPrizeBox:
		msg.(*SPrizeBox).Code = 1055
		b, err := msg.(*SPrizeBox).Marshal()
		return 1055, b, err
	case *SDeleteMail:
		msg.(*SDeleteMail).Code = 1072
		b, err := msg.(*SDeleteMail).Marshal()
		return 1072, b, err
	case *SBetting:
		msg.(*SBetting).Code = 4201
		b, err := msg.(*SBetting).Marshal()
		return 4201, b, err
	case *SEnterClassicRoom:
		msg.(*SEnterClassicRoom).Code = 4060
		b, err := msg.(*SEnterClassicRoom).Marshal()
		return 4060, b, err
	case *SPushDiscard:
		msg.(*SPushDiscard).Code = 4106
		b, err := msg.(*SPushDiscard).Marshal()
		return 4106, b, err
	case *SPushDealerDeal:
		msg.(*SPushDealerDeal).Code = 4110
		b, err := msg.(*SPushDealerDeal).Marshal()
		return 4110, b, err
	case *SEnterRoom:
		msg.(*SEnterRoom).Code = 4000
		b, err := msg.(*SEnterRoom).Marshal()
		return 4000, b, err
	case *SKick:
		msg.(*SKick).Code = 4005
		b, err := msg.(*SKick).Marshal()
		return 4005, b, err
	case *SPushDealer:
		msg.(*SPushDealer).Code = 4009
		b, err := msg.(*SPushDealer).Marshal()
		return 4009, b, err
	case *SVoteResult:
		msg.(*SVoteResult).Code = 4018
		b, err := msg.(*SVoteResult).Marshal()
		return 4018, b, err
	case *SGameRecord:
		msg.(*SGameRecord).Code = 4020
		b, err := msg.(*SGameRecord).Marshal()
		return 4020, b, err
	case *SPushDealerBu:
		msg.(*SPushDealerBu).Code = 4112
		b, err := msg.(*SPushDealerBu).Marshal()
		return 4112, b, err
	case *SWxpayQuery:
		msg.(*SWxpayQuery).Code = 3004
		b, err := msg.(*SWxpayQuery).Marshal()
		return 3004, b, err
	case *SBank:
		msg.(*SBank).Code = 1030
		b, err := msg.(*SBank).Marshal()
		return 1030, b, err
	case *SFreeGamestart:
		msg.(*SFreeGamestart).Code = 4048
		b, err := msg.(*SFreeGamestart).Marshal()
		return 4048, b, err
	case *SClassicGameover:
		msg.(*SClassicGameover).Code = 4062
		b, err := msg.(*SClassicGameover).Marshal()
		return 4062, b, err
	case *SZiGameover:
		msg.(*SZiGameover).Code = 4102
		b, err := msg.(*SZiGameover).Marshal()
		return 4102, b, err
	case *SPushDeal:
		msg.(*SPushDeal).Code = 4109
		b, err := msg.(*SPushDeal).Marshal()
		return 4109, b, err
	case *SPushNewBetting:
		msg.(*SPushNewBetting).Code = 4204
		b, err := msg.(*SPushNewBetting).Marshal()
		return 4204, b, err
	case *SLottery:
		msg.(*SLottery).Code = 4221
		b, err := msg.(*SLottery).Marshal()
		return 4221, b, err
	case *SWxLogin:
		msg.(*SWxLogin).Code = 1004
		b, err := msg.(*SWxLogin).Marshal()
		return 1004, b, err
	case *SPrizeDraw:
		msg.(*SPrizeDraw).Code = 1054
		b, err := msg.(*SPrizeDraw).Marshal()
		return 1054, b, err
	case *SDanInfo:
		msg.(*SDanInfo).Code = 1080
		b, err := msg.(*SDanInfo).Marshal()
		return 1080, b, err
	case *SQualifying:
		msg.(*SQualifying).Code = 1081
		b, err := msg.(*SQualifying).Marshal()
		return 1081, b, err
	case *SZiCamein:
		msg.(*SZiCamein).Code = 4104
		b, err := msg.(*SZiCamein).Marshal()
		return 4104, b, err
	case *SDealerList:
		msg.(*SDealerList).Code = 4043
		b, err := msg.(*SDealerList).Marshal()
		return 4043, b, err
	case *SFreeSit:
		msg.(*SFreeSit).Code = 4044
		b, err := msg.(*SFreeSit).Marshal()
		return 4044, b, err
	case *SBettingRecord:
		msg.(*SBettingRecord).Code = 4205
		b, err := msg.(*SBettingRecord).Marshal()
		return 4205, b, err
	case *SMailNotice:
		msg.(*SMailNotice).Code = 1070
		b, err := msg.(*SMailNotice).Marshal()
		return 1070, b, err
	case *SLeave:
		msg.(*SLeave).Code = 4004
		b, err := msg.(*SLeave).Marshal()
		return 4004, b, err
	case *SReady:
		msg.(*SReady).Code = 4006
		b, err := msg.(*SReady).Marshal()
		return 4006, b, err
	case *SDraw:
		msg.(*SDraw).Code = 4007
		b, err := msg.(*SDraw).Marshal()
		return 4007, b, err
	case *SFreeDealer:
		msg.(*SFreeDealer).Code = 4042
		b, err := msg.(*SFreeDealer).Marshal()
		return 4042, b, err
	case *SPushCurrency:
		msg.(*SPushCurrency).Code = 1028
		b, err := msg.(*SPushCurrency).Marshal()
		return 1028, b, err
	case *SNiu:
		msg.(*SNiu).Code = 4011
		b, err := msg.(*SNiu).Marshal()
		return 4011, b, err
	case *SPushJackpot:
		msg.(*SPushJackpot).Code = 4202
		b, err := msg.(*SPushJackpot).Marshal()
		return 4202, b, err
	case *SLogin:
		msg.(*SLogin).Code = 1000
		b, err := msg.(*SLogin).Marshal()
		return 1000, b, err
	case *SRegist:
		msg.(*SRegist).Code = 1002
		b, err := msg.(*SRegist).Marshal()
		return 1002, b, err
	case *SConfig:
		msg.(*SConfig).Code = 1020
		b, err := msg.(*SConfig).Marshal()
		return 1020, b, err
	case *SBankrupts:
		msg.(*SBankrupts).Code = 1052
		b, err := msg.(*SBankrupts).Marshal()
		return 1052, b, err
	case *SVipList:
		msg.(*SVipList).Code = 1062
		b, err := msg.(*SVipList).Marshal()
		return 1062, b, err
	case *SEnterZiRoom:
		msg.(*SEnterZiRoom).Code = 4100
		b, err := msg.(*SEnterZiRoom).Marshal()
		return 4100, b, err
	case *SDanNotice:
		msg.(*SDanNotice).Code = 1084
		b, err := msg.(*SDanNotice).Marshal()
		return 1084, b, err
	case *SLaunchVote:
		msg.(*SLaunchVote).Code = 4016
		b, err := msg.(*SLaunchVote).Marshal()
		return 4016, b, err
	case *SVote:
		msg.(*SVote).Code = 4017
		b, err := msg.(*SVote).Marshal()
		return 4017, b, err
	case *SEnterFreeRoom:
		msg.(*SEnterFreeRoom).Code = 4040
		b, err := msg.(*SEnterFreeRoom).Marshal()
		return 4040, b, err
	case *SFreeCamein:
		msg.(*SFreeCamein).Code = 4041
		b, err := msg.(*SFreeCamein).Marshal()
		return 4041, b, err
	case *SGameover:
		msg.(*SGameover).Code = 4012
		b, err := msg.(*SGameover).Marshal()
		return 4012, b, err
	case *SPubDraw:
		msg.(*SPubDraw).Code = 4080
		b, err := msg.(*SPubDraw).Marshal()
		return 4080, b, err
	case *SPushStatus:
		msg.(*SPushStatus).Code = 4111
		b, err := msg.(*SPushStatus).Marshal()
		return 4111, b, err
	case *SBuy:
		msg.(*SBuy).Code = 3000
		b, err := msg.(*SBuy).Marshal()
		return 3000, b, err
	case *SChatText:
		msg.(*SChatText).Code = 2003
		b, err := msg.(*SChatText).Marshal()
		return 2003, b, err
	case *SBuildAgent:
		msg.(*SBuildAgent).Code = 1026
		b, err := msg.(*SBuildAgent).Marshal()
		return 1026, b, err
	case *SPushVip:
		msg.(*SPushVip).Code = 1063
		b, err := msg.(*SPushVip).Marshal()
		return 1063, b, err
	case *SCreateRoom:
		msg.(*SCreateRoom).Code = 4001
		b, err := msg.(*SCreateRoom).Marshal()
		return 4001, b, err
	case *SBroadcast:
		msg.(*SBroadcast).Code = 2006
		b, err := msg.(*SBroadcast).Marshal()
		return 2006, b, err
	case *SGetCurrency:
		msg.(*SGetCurrency).Code = 1024
		b, err := msg.(*SGetCurrency).Marshal()
		return 1024, b, err
	case *SPushDraw:
		msg.(*SPushDraw).Code = 4105
		b, err := msg.(*SPushDraw).Marshal()
		return 4105, b, err
	case *SChatVoice:
		msg.(*SChatVoice).Code = 2004
		b, err := msg.(*SChatVoice).Marshal()
		return 2004, b, err
	case *SGetPrize:
		msg.(*SGetPrize).Code = 4081
		b, err := msg.(*SGetPrize).Marshal()
		return 4081, b, err
	case *SPushAuto:
		msg.(*SPushAuto).Code = 4107
		b, err := msg.(*SPushAuto).Marshal()
		return 4107, b, err
	case *SPushPaoHu:
		msg.(*SPushPaoHu).Code = 4113
		b, err := msg.(*SPushPaoHu).Marshal()
		return 4113, b, err
	case *SPushBetting:
		msg.(*SPushBetting).Code = 4203
		b, err := msg.(*SPushBetting).Marshal()
		return 4203, b, err
	case *SPing:
		msg.(*SPing).Code = 1050
		b, err := msg.(*SPing).Marshal()
		return 1050, b, err
	case *SClassicList:
		msg.(*SClassicList).Code = 1060
		b, err := msg.(*SClassicList).Marshal()
		return 1060, b, err
	case *SFreeGameover:
		msg.(*SFreeGameover).Code = 4050
		b, err := msg.(*SFreeGameover).Marshal()
		return 4050, b, err
	case *SPrizeList:
		msg.(*SPrizeList).Code = 1053
		b, err := msg.(*SPrizeList).Marshal()
		return 1053, b, err
	case *SFreeTrend:
		msg.(*SFreeTrend).Code = 4051
		b, err := msg.(*SFreeTrend).Marshal()
		return 4051, b, err
	case *SCreateZiRoom:
		msg.(*SCreateZiRoom).Code = 4101
		b, err := msg.(*SCreateZiRoom).Marshal()
		return 4101, b, err
	case *SZiGameRecord:
		msg.(*SZiGameRecord).Code = 4103
		b, err := msg.(*SZiGameRecord).Marshal()
		return 4103, b, err
	case *SLotteryInfo:
		msg.(*SLotteryInfo).Code = 4220
		b, err := msg.(*SLotteryInfo).Marshal()
		return 4220, b, err
	case *SWxpayOrder:
		msg.(*SWxpayOrder).Code = 3002
		b, err := msg.(*SWxpayOrder).Marshal()
		return 3002, b, err
	case *SApplePay:
		msg.(*SApplePay).Code = 3006
		b, err := msg.(*SApplePay).Marshal()
		return 3006, b, err
	case *SLoginOut:
		msg.(*SLoginOut).Code = 1006
		b, err := msg.(*SLoginOut).Marshal()
		return 1006, b, err
	case *SDealer:
		msg.(*SDealer).Code = 4008
		b, err := msg.(*SDealer).Marshal()
		return 4008, b, err
	case *SOperate:
		msg.(*SOperate).Code = 4108
		b, err := msg.(*SOperate).Marshal()
		return 4108, b, err
	case *SUserData:
		msg.(*SUserData).Code = 1022
		b, err := msg.(*SUserData).Marshal()
		return 1022, b, err
	case *SCamein:
		msg.(*SCamein).Code = 4003
		b, err := msg.(*SCamein).Marshal()
		return 4003, b, err
	case *SBet:
		msg.(*SBet).Code = 4010
		b, err := msg.(*SBet).Marshal()
		return 4010, b, err
	case *SFreeBet:
		msg.(*SFreeBet).Code = 4046
		b, err := msg.(*SFreeBet).Marshal()
		return 4046, b, err
	case *SRoomList:
		msg.(*SRoomList).Code = 4115
		b, err := msg.(*SRoomList).Marshal()
		return 4115, b, err
	default:
		return 0, []byte{}, errors.New("unknown message")
	}
}