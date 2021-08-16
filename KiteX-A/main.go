package main

import (
	api "github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api/servicea"
	"log"
)

func main() {
	svr := api.NewServer(new(ServiceAImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
