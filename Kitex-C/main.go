package main

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api"
	"github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api/servicea"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
)

var (
	serverAClient servicea.Client
	count         uint64
	errCount      uint64
)

const (
	Concurrent = 100
)

func sendReq(waitGroup *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Test failed %v, %s.", err, string(debug.Stack()))
		}
		waitGroup.Done()
	}()

	for {
		req := api.NewRequest()
		req.SetMessage("test")
		resp, err := serverAClient.ServiceA(context.Background(), req)
		if err != nil {
			log.Println("resp = ", resp, " err = ", err)
			atomic.AddUint64(&errCount, 1)
		}

		atomic.AddUint64(&count, 1)
		time.Sleep(time.Millisecond * 50)
	}
}

func run() {
	var wg sync.WaitGroup
	wg.Add(Concurrent)
	for i := 0; i < Concurrent; i++ {
		go sendReq(&wg)
	}
	wg.Wait()
}

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
	options = append(options, client.WithHostPorts("0.0.0.0:8888"))

	serverAClient, err = servicea.NewClient("servicea", options...)
	if err != nil {
		panic(err)
	}

	go func() {
		lastCount := count
		errLastCount := count
		for range time.Tick(time.Second) {
			curCount := atomic.LoadUint64(&count)
			log.Println("qps = ", curCount-lastCount)
			lastCount = curCount

			errCurCount := atomic.LoadUint64(&errCount)
			log.Println("err qps = ", errCurCount-errLastCount)
			errLastCount = errCurCount
		}
	}()
}

func main() {
	fmt.Println("run")
	run()
	fmt.Println("run exit")
}
