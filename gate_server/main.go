package main

import (
	"time"

	core "github.com/kermitbu/gant-core"
	log "github.com/kermitbu/gant-log"
)

type GateServer struct {
	core.CoreServer
}

func (svr *GateServer) InitTcpSvr(port string) {

	// 有节点加入
	svr.Handle(1, func(req *core.GRequest, res *core.GResponse) {
	})

	// 有节点离开
	svr.Handle(2, func(req *core.GRequest, res *core.GResponse) {

	})

	err := svr.InitConnectAsServer(port)
	if err != nil {
		log.Error("%s", err.Error())
	}
}

func (svr *GateServer) InitTcpClient(masterAddr, serveAddr string) {
	var complete = make(chan int, 1)

	err := svr.InitConnectAsClient(masterAddr, complete)
	if err != nil {
		log.Warn("无法连接到Master服务器，稍后发起重连。。。")

		// 初始化失败，10秒后再次初始化TCP连接
		time.AfterFunc(time.Second*10, func() {
			svr.InitTcpClient(masterAddr, serveAddr)
		})
		return
	}

	for {
		// 等待连接成功。
		result := <-complete
		if result > 0 {
			log.Warn("无法连接到Master服务器，稍后发起重连。。。")
			time.AfterFunc(time.Second*10, func() {
				svr.InitTcpClient(masterAddr, serveAddr)
			})
		}
	}

}

func main() {
	svr := new(GateServer)

	svr.InitTcpClient("127.0.0.1", "9666")

	svr.InitTcpSvr("9687")
}
