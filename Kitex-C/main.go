package main

import (
	"context"
	"fmt"
	"github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api"
	"github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api/servicea"
	"github.com/cloudwego/kitex/client"
	"log"
)

func main() {
	c, err := servicea.NewClient("servicea", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	req := api.NewRequest()
	req.SetMessage("hello")
	resp, err := c.ServiceA(context.Background(), req)
	fmt.Println("resp = ", resp, " err = ", err)
}
