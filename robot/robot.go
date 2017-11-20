/**********************************************************
 * Author        : piaohua
 * Email         : 814004090@qq.com
 * Last modified : 2017-11-19 11:32:23
 * Filename      : robot.go
 * Description   : 机器人
 * *******************************************************/
package main

import (
	"os"
	"os/signal"
	"runtime"
	"utils"

	"goplay/glog"

	ini "gopkg.in/ini.v1"
)

const (
	ROBOTS_NAME = "RobotMsg"
	VERSION     = "0.0.1"
)

var (
	cfg *ini.File
	sec *ini.Section
	err error

	rbs *RobotServer

	aesEnc *utils.AesEncrypt
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	defer glog.Flush()
	//日志定义
	glog.Init()
	//加载配置
	cfg, err = ini.Load("conf.ini")
	if err != nil {
		panic(err)
	}
	cfg.BlockMode = false //只读
	//初始化
	aesInit()
	//
	bind := cfg.Section("robot").Key("bind").Value()
	name := cfg.Section("cookie").Key("name").Value()
	phone := cfg.Section("robot").Key("phone").Value()
	//启动服务
	rbs = new(RobotServer)
	rbs.phone = phone
	if rbs != nil {
		rbs.Start() //启动服务
		rbs.NewRemote(bind, name)
	}
	//启动服务
	signalListen()
	//关闭服务
	//关闭websocket连接
	if rbs != nil {
		rbs.Close() //关闭服务
	}
	//延迟等待
}

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c)
	//signal.Stop(c)
	for {
		s := <-c
		glog.Fatal("get signal:", s)
	}
}

//加密初始化
func aesInit() {
	aesEnc = new(utils.AesEncrypt)
	key := cfg.Section("login").Key("key").Value()
	aesEnc.SetKey([]byte(key))
}

//加密
func aesEn(doc string) (arrEncrypt []byte) {
	arrEncrypt, err = aesEnc.Encrypt([]byte(doc))
	if err != nil {
		glog.Errorf("arrEncrypt: %s", doc)
	}
	return
}

//解密
func aesDe(arrEncrypt []byte) (strMsg string) {
	bMsg, err := aesEnc.Decrypt(arrEncrypt)
	if err != nil {
		glog.Errorf("arrEncrypt: %s", string(arrEncrypt))
	}
	strMsg = string(bMsg)
	return
}