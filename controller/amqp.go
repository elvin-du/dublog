package controller

import (
	"dublog/router"
	"log"
)

func init() {
	router.Router().Register("/amqp/publish", Handler(_amqpController.Publish))
}

type AmqpController struct{}

var (
	_amqpController = &AmqpController{}
)

func (*AmqpController) Publish(ctx *router.Context) {
	err := ctx.Request.ParseForm()
	if nil != err {
		log.Println(err)
		return
	}
    ctx.ResponseWriter.WriteHeader(200)
    ctx.ResponseWriter.Write([]byte(ctx.Request.FormValue("msg")))
}
