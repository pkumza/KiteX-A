package rpc

import (
	"github.com/Duslia997/KiteX-A/KiteX-B/kitex_gen/api/serviceb"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"time"
)

var ServerBClient serviceb.Client

func init() {
	var err error
	options := client.WithLongConnection(connpool.IdleConfig{
		MaxIdlePerAddress: 100,
		MaxIdleGlobal:     1000,
		MaxIdleTimeout:    60 * time.Second,
	})
	ServerBClient, err = serviceb.NewClient("serviceb", client.WithHostPorts("0.0.0.0:8887"), options)
	if err != nil {
		panic(err)
	}
}
