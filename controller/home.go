package controller

import(
    "dublog/router"
    "log"
//    "dublog/model"
    "text/template"
)

func init(){
    router.Router().Register("/home",Handler(_homeController.Home))
}

type HomeController struct{}

var(
    _homeController = &HomeController{}
)

func (*HomeController)Home(ctx *router.Context){
    tpl,err := template.ParseFiles("./static/home.html")
    if nil != err{
        log.Println(err)
        return
    }

    err = tpl.Execute(ctx,nil)
    if nil != err{
        log.Println(err)
        return
    }

//    ctx.WriteHeader(200)
//    res,err := model.HomeModel().Index()
//    if nil != err{
//        log.Println(err)
//        ctx.WriteHeader(500)
//        ctx.Write([]byte(err.Error()))
//        return
//    }
//    ctx.Write([]byte(res["user"]))
}


