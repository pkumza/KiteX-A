package main

import (
	api "github.com/Duslia997/KiteX-A/KiteX-B/kitex_gen/api/serviceb"
	"github.com/cloudwego/kitex/server"
	"log"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:8887")

	svr := api.NewServer(new(ServiceBImpl), server.WithServiceAddr(addr))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
