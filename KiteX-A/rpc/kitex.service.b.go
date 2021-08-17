package rpc

import (
	"time"

	"github.com/Duslia997/KiteX-A/KiteX-B/kitex_gen/api/serviceb"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
)

var ServerBClient serviceb.Client

func init() {
	var err error
	options := []client.Option{}
	options = append(options, client.WithLongConnection(connpool.IdleConfig{
		MaxIdlePerAddress: 100,
		MaxIdleGlobal:     1000,
		MaxIdleTimeout:    60 * time.Second,
	}))
	options = append(options, client.WithRPCTimeout(time.Second*5))
	options = append(options, client.WithConnectTimeout(time.Millisecond*50))
	options = append(options, client.WithHostPorts("0.0.0.0:8887"))

	ServerBClient, err = serviceb.NewClient("serviceb", options...)
	if err != nil {
		panic(err)
	}
}
