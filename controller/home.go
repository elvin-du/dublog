package controller

import(
    "dublog/router"
)

func init(){
    router.Router().Register("/home",Handler(_homeController.Home))
}

type HomeController struct{}

var(
    _homeController = &HomeController{}
)

func (*HomeController)Home(ctx *router.Context){
    ctx.WriteHeader(200)
    ctx.Write([]byte("I love you!!"))
}


