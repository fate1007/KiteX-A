package main

import (
	api "kitex.service.a/kitex_gen/api/servicea"
	"log"
)

func main() {
	svr := api.NewServer(new(ServiceAImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
