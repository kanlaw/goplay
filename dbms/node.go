package main

import (
	"time"

	"goplay/glog"
	"goplay/pb"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"github.com/gogo/protobuf/proto"
)

var (
	nodePid *actor.PID
)

//数据库操作服务
type DBMSActor struct {
	Name string
	//大厅服务
	hallPid *actor.PID
	//日志服务
	logger *actor.PID
	//网关节点
	gates map[string]*actor.PID
	//关闭通道
	stopCh chan struct{}
	//更新状态
	status bool
	//计时
	timer int
}

func (a *DBMSActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *pb.Request:
		ctx.Respond(&pb.Response{})
	case *actor.Started:
		glog.Notice("Starting, initialize actor here")
	case *actor.Stopping:
		glog.Notice("Stopping, actor is about to shut down")
	case *actor.Stopped:
		glog.Notice("Stopped, actor and its children are stopped")
	case *actor.Restarting:
		glog.Notice("Restarting, actor is about to restart")
	case *actor.ReceiveTimeout:
		glog.Infof("ReceiveTimeout: %v", ctx.Self().String())
	case proto.Message:
		a.Handler(msg, ctx)
	default:
		glog.Errorf("unknown message %v", msg)
	}
}

func newDBMSActor() actor.Actor {
	a := new(DBMSActor)
	a.Name = cfg.Section("dbms").Name()
	a.gates = make(map[string]*actor.PID)
	a.logger = NewLogger()
	a.stopCh = make(chan struct{})
	return a
}

func NewRemote(bind, name, room, role, mail, bets string) {
	remote.Start(bind)
	//
	//remote.Register(name, actor.FromProducer(newDBMSActor))
	msg := new(pb.ServeStart)
	dbmsProps := actor.
		FromInstance(newDBMSActor())
	remote.Register(name, dbmsProps)
	nodePid, err = actor.SpawnNamed(dbmsProps, name)
	if err != nil {
		glog.Fatalf("nodePid err %v", err)
	}
	glog.Infof("nodePid %s", nodePid.String())
	nodePid.Tell(msg)
	//
	//remote.Register(room, actor.FromProducer(newRoomActor))
	roomProps := actor.
		FromInstance(newRoomActor())
	remote.Register(room, roomProps)
	roomPid, err = actor.SpawnNamed(roomProps, room)
	if err != nil {
		glog.Fatalf("roomPid err %v", err)
	}
	glog.Infof("roomPid %s", roomPid.String())
	//roomPid.Tell(new(pb.HallConnect))
	roomPid.Tell(msg)
	//
	//remote.Register(role, actor.FromProducer(newRoleActor))
	roleProps := actor.
		FromInstance(newRoleActor())
	remote.Register(role, roleProps)
	rolePid, err = actor.SpawnNamed(roleProps, role)
	if err != nil {
		glog.Fatalf("rolePid err %v", err)
	}
	glog.Infof("rolePid %s", rolePid.String())
	//rolePid.Tell(new(pb.HallConnect))
	rolePid.Tell(msg)
	//
	//remote.Register(mail, actor.FromProducer(newMailActor))
	mailProps := actor.
		FromInstance(newMailActor())
	remote.Register(mail, mailProps)
	mailPid, err = actor.SpawnNamed(mailProps, mail)
	if err != nil {
		glog.Fatalf("mailPid err %v", err)
	}
	glog.Infof("mailPid %s", mailPid.String())
	//mailPid.Tell(new(pb.HallConnect))
	mailPid.Tell(msg)
	//
	//remote.Register(bets, actor.FromProducer(newBetsActor))
	betsProps := actor.
		FromInstance(newBetsActor())
	remote.Register(bets, betsProps)
	betsPid, err = actor.SpawnNamed(betsProps, bets)
	if err != nil {
		glog.Fatalf("betsPid err %v", err)
	}
	glog.Infof("betsPid %s", betsPid.String())
	//betsPid.Tell(new(pb.HallConnect))
	betsPid.Tell(msg)
}

//关闭
func Stop() {
	timeout := 3 * time.Second
	msg := new(pb.ServeStop)
	if mailPid != nil {
		res1, err1 := mailPid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("mailPid Stop err: %v", err1)
		}
		response1 := res1.(*pb.ServeStoped)
		glog.Debugf("response1: %#v", response1)
		mailPid.Stop()
	}
	if betsPid != nil {
		res1, err1 := betsPid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("betsPid Stop err: %v", err1)
		}
		response1 := res1.(*pb.ServeStoped)
		glog.Debugf("response1: %#v", response1)
		betsPid.Stop()
	}
	if rolePid != nil {
		res1, err1 := rolePid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("rolePid Stop err: %v", err1)
		}
		response1 := res1.(*pb.ServeStoped)
		glog.Debugf("response1: %#v", response1)
		rolePid.Stop()
	}
	if roomPid != nil {
		res1, err1 := roomPid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("roomPid Stop err: %v", err1)
		}
		response1 := res1.(*pb.ServeStoped)
		glog.Debugf("response1: %#v", response1)
		roomPid.Stop()
	}
	if nodePid != nil {
		res1, err1 := nodePid.RequestFuture(msg, timeout).Result()
		if err1 != nil {
			glog.Errorf("nodePid Stop err: %v", err1)
		}
		response1 := res1.(*pb.ServeStoped)
		glog.Debugf("response1: %#v", response1)
		nodePid.Stop()
	}
}
