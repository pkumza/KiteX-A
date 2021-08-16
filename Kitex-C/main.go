package main

import (
	"context"
	"fmt"
	"github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api"
	"github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api/servicea"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"log"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
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
	options := client.WithLongConnection(connpool.IdleConfig{
		MaxIdlePerAddress: 100,
		MaxIdleGlobal:     1000,
		MaxIdleTimeout:    60 * time.Second,
	})
	serverAClient, err = servicea.NewClient("servicea", client.WithHostPorts("0.0.0.0:8888"), options)
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
