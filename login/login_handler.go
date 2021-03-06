package main

import (
	"fmt"

	"api/wxpay"
	"goplay/glog"
	"goplay/pb"

	"github.com/valyala/fasthttp"
)

func Start(addr string) {
	if err = fasthttp.ListenAndServe(addr, requestHandler); err != nil {
		glog.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func fooHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())
}

func barHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)
}

func logHandler(ctx *fasthttp.RequestCtx) {
	glog.Debugf("Request method is %q\n", ctx.Method())
	glog.Debugf("RequestURI is %q\n", ctx.RequestURI())
	glog.Debugf("Requested path is %q\n", ctx.Path())
	glog.Debugf("Host is %q\n", ctx.Host())
	glog.Debugf("Query string is %q\n", ctx.QueryArgs())
	glog.Debugf("User-Agent is %q\n", ctx.UserAgent())
	glog.Debugf("Connection has been established at %s\n", ctx.ConnTime())
	glog.Debugf("Request has been started at %s\n", ctx.Time())
	glog.Debugf("Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	glog.Debugf("Your ip is %q\n\n", ctx.RemoteIP())
	glog.Debugf("Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)
}

// 接收交易结果通知
func wxpayHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Content-type", "text/plain;;charset=UTF-8")
	switch string(ctx.Method()) {
	case "POST":
		result := ctx.PostBody()
		tradeResult, err := wxpay.ParseTradeResult(result)
		glog.Debugf("result %s, tradeResult %#v, err %v", string(result), tradeResult, err)
		if err == nil {
			//发货
			msg := new(pb.WxpayCallback)
			msg.Result = string(result)
			nodePid.Tell(msg)
		} else {
			glog.Errorf("trade result err: %v", err)
		}
	default:
		glog.Error("wxpay method err")
	}
	fmt.Fprintf(ctx, wxpay.TradeRespXml())
}

//TODO version rule
func gateHandler(ctx *fasthttp.RequestCtx) {
	r := "gate.node"
	r += string(ctx.QueryArgs().Peek("version"))
	glog.Debugf("gate node : %s", r)
	sec1, err1 := cfg.GetSection(r)
	if err1 != nil {
		glog.Error("Unknwon version ", err1)
		ctx.Error("Unknwon version", fasthttp.StatusBadRequest)
		return
	}
	key2, err2 := sec1.GetKey("host")
	if err2 != nil {
		glog.Error("Unknwon version ", err2)
		ctx.Error("Unknwon version", fasthttp.StatusBadRequest)
		return
	}
	host := key2.Value()
	logHandler(ctx)
	glog.Debugf("gate host %s", host)
	fmt.Fprintf(ctx, "%s", string(aesEn(host)))
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/gate":
		gateHandler(ctx)
	case "/wxpay":
		wxpayHandler(ctx)
	case "/foo":
		fooHandler(ctx)
	case "/bar":
		barHandler(ctx)
	default:
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	}
}
