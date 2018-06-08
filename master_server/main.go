package main

import (
	"github.com/kermitbu/gant-log"

	"github.com/kermitbu/gant-core"
)

type MasterServer struct {
	core.CoreServer
}

func InitTcpSvr(port string) {
	svr := new(MasterServer)

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

func main() {
	InitTcpSvr("3900")
}
