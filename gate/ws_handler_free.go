package main

import (
	"encoding/json"

	"goplay/data"
	"goplay/game/config"
	"goplay/glog"
	"goplay/pb"
	"utils"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家百人场请求处理
func (ws *WSConn) HandlerFree(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CEnterFreeRoom:
		arg := msg.(*pb.CEnterFreeRoom)
		glog.Debugf("CEnterFreeRoom %#v", arg)
		ws.freeEnter(arg, ctx)
	case *pb.CFreeDealer:
		arg := msg.(*pb.CFreeDealer)
		glog.Debugf("CFreeDealer %#v", arg)
	case *pb.CDealerList:
		arg := msg.(*pb.CDealerList)
		glog.Debugf("CDealerList %#v", arg)
		//ws.rolePid.Request(arg, ctx.Self())
	case *pb.CFreeSit:
		arg := msg.(*pb.CFreeSit)
		glog.Debugf("CFreeSit %#v", arg)
		//ws.rolePid.Request(msg2, ctx.Self())
	case *pb.CFreeBet:
		arg := msg.(*pb.CFreeBet)
		glog.Debugf("CFreeBet %#v", arg)
		//ws.Send(arg)
	case *pb.CFreeTrend:
		arg := msg.(*pb.CFreeTrend)
		glog.Debugf("CFreeTrend %#v", arg)
		//ws.rolePid.Request(msg2, ctx.Self())
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//进入自由场
func (ws *WSConn) freeEnter(arg *pb.CEnterFreeRoom, ctx actor.Context) {
	if ws.gamePid != nil {
		//进入房间
		ws.entryRoom()
		return
	}
	stoc := new(pb.SEnterFreeRoom)
	//匹配可以进入的房间
	response1 := ws.matchRoom(data.ROOM_FREE)
	if response1 == nil {
		stoc.Error = pb.RoomNotExist
		ws.Send(stoc)
		return
	} else if response1.Desk != nil {
		ws.gamePid = response1.Desk
		//进入房间
		ws.entryRoom()
		return
	}
	//节点不存在
	if response1.Node == nil {
		stoc.Error = pb.RoomNotExist
		ws.Send(stoc)
		return
	}
	//创建新房间
	response2 := ws.createRoom(data.ROOM_FREE)
	if response1 == nil {
		stoc.Error = pb.RoomNotExist
		ws.Send(stoc)
		return
	}
	//新桌子
	response3 := ws.spawnRoom(response1.Node, response2.Data)
	if response3 == nil {
		stoc.Error = pb.RoomNotExist
		ws.Send(stoc)
		return
	}
	ws.gamePid = response3.Desk
	ws.entryRoom()
}

//自由场数据
func (ws *WSConn) freeData() string {
	now := uint32(utils.Timestamp())
	count := uint32(config.GetEnv(data.ENV11))
	deskData := data.NewDeskData(0, 0, data.ROOM_FREE, 1,
		0, 0, 0, 0, count, now, 0, 0, 0, "", "", "", "")
	result, err := json.Marshal(deskData)
	if err != nil {
		glog.Errorf("deskData Marshal err %v", err)
		return ""
	}
	return string(result)
}
