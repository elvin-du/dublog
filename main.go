package main

import (
	"dublog/config"
	"dublog/router"
	"dublog/service"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
    "errors"
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
    go StartWeb()
	err := http.ListenAndServe(addr, router.Router())
	if nil != err {
		log.Fatalln(err)
	}
}

func StartWeb() {
	time.Sleep(time.Second * 1)

	err := Open("http://"+addr + "/home")
	if nil != err {
		log.Println(err)
	}
    log.Println("open ok")
}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func Open(uri string) error {
	run, ok := commands[runtime.GOOS]
	if !ok {
		log.Printf("don't know how to open things on %s platform", runtime.GOOS)
		return errors.New("cannot find os platform")
	}

	cmd := exec.Command(run, uri)
	return cmd.Start()
}
