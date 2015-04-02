package main

import (
	"dublog/config"
	"dublog/router"
	"dublog/service"
	"log"
	"net/http"
)

var addr string

func init() {
	err := config.Get("http:addr", &addr)
	if nil != err {
		log.Fatalln(err)
	}
	log.Println(addr)
}

func main() {
	service.Start()
	//	addr := ":9999"
	err := http.ListenAndServe(addr, router.Router())
	if nil != err {
		log.Fatalln(err)
	}
}
