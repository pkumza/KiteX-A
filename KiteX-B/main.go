package main

import (
	api "github.com/Duslia997/KiteX-B/KiteX-B/kitex_gen/api/serviceb"
	"log"
)

func main() {
	svr := api.NewServer(new(ServiceBImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
