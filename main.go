package main

import (
	"log"
	"net/http"
    "dublog/router"
    "dublog/service"
)

func main() {
    service.Start()
	addr := ":9999"
	err := http.ListenAndServe(addr, router.Router())
	if nil != err {
		log.Fatalln(err)
	}
}


