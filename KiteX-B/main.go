package main

import (
	api "kitex.service.b/kitex_gen/api/serviceb"
	"log"
)

func main() {
	svr := api.NewServer(new(ServiceBImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
