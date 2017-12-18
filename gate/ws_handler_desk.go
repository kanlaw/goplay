package main

import (
	"time"

	"goplay/data"
	"goplay/glog"
	"goplay/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
)

//玩家桌子常用共有操作请求处理
func (ws *WSConn) HandlerDesk(msg interface{}, ctx actor.Context) {
	switch msg.(type) {
	case *pb.CLeave:
		arg := msg.(*pb.CLeave)
		glog.Debugf("CEnterFreeRoom %#v", arg)
		//rsp := handler.LotteryInfo(arg)
		//ws.Send(rsp)
	case *pb.CKick:
		arg := msg.(*pb.CKick)
		glog.Debugf("CKick %#v", arg)
	case *pb.CReady:
		arg := msg.(*pb.CReady)
		glog.Debugf("CReady %#v", arg)
		//ws.rolePid.Request(arg, ctx.Self())
	case *pb.CLaunchVote:
		arg := msg.(*pb.CLaunchVote)
		glog.Debugf("CLaunchVote %#v", arg)
		//ws.rolePid.Request(msg2, ctx.Self())
	case *pb.CVote:
		arg := msg.(*pb.CVote)
		glog.Debugf("CVote %#v", arg)
		//ws.Send(arg)
	case *pb.CRoomList:
		arg := msg.(*pb.CRoomList)
		glog.Debugf("CRoomList %#v", arg)
		//ws.rolePid.Request(msg2, ctx.Self())
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

//进入房间
func (ws *WSConn) entryRoom(ctx actor.Context) {
	if ws.gamePid == nil {
		glog.Errorf("not in the room: %s", ws.User.GetUserid())
		return
	}
	//TODO 进入房间
	//ws.gamePid.Tell(msg)
}

//大厅中匹配可用房间
func (ws *WSConn) matchRoom(rtype uint32) *pb.MatchedDesk {
	//匹配可以进入的房间
	msg1 := new(pb.MatchDesk)
	switch rtype {
	case data.ROOM_FREE:
		msg1.Rtype = data.ROOM_FREE
		//节点注册名称,TODO 多节点处理
		msg1.Name = cfg.Section("game.free").Name()
	}
	timeout := 3 * time.Second
	res1, err1 := ws.hallPid.RequestFuture(msg1, timeout).Result()
	if err1 != nil {
		glog.Errorf("matchRoom err: %v", err1)
		return nil
	}
	response1 := res1.(*pb.MatchedDesk)
	glog.Debugf("response1: %#v", response1)
	return response1
}

//数据中心创建房间
func (ws *WSConn) createRoom(rtype uint32) *pb.CreatedDesk {
	msg2 := new(pb.CreateDesk)
	switch rtype {
	case data.ROOM_FREE:
		msg2.Data = ws.freeData()
	}
	if msg2.Data == "" {
		return nil
	}
	timeout := 3 * time.Second
	res2, err2 := ws.roomPid.RequestFuture(msg2, timeout).Result()
	if err2 != nil {
		glog.Errorf("createRoom err: %v", err2)
		return nil
	}
	response2 := res2.(*pb.CreatedDesk)
	glog.Debugf("response2: %#v", response2)
	if response2.Error != pb.OK {
		return nil
	}
	return response2
}

//创建新桌子
func (ws *WSConn) spawnRoom(node *actor.PID, rdata string) *pb.SpawnedDesk {
	if rdata == "" || node == nil {
		return nil
	}
	msg2 := new(pb.SpawnDesk)
	msg2.Data = rdata
	timeout := 3 * time.Second
	res2, err2 := node.RequestFuture(msg2, timeout).Result()
	if err2 != nil {
		glog.Errorf("spawnRoom err: %v", err2)
		return nil
	}
	response2 := res2.(*pb.SpawnedDesk)
	glog.Debugf("response2: %#v", response2)
	if response2.Desk == nil {
		return nil
	}
	return response2
}
