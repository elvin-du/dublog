package controller

import(
    "dublog/router"
    "log"
)

func Handler(fun Action)router.HandlerFunc{
    return commonHandler(fun)
}

func AuthHandler(fun Action)router.HandlerFunc{
    return commonHandler(fun,_AuthHelper.AuthToken)
}

func commonHandler(fun Action,preHandlers ...PreHandler)router.HandlerFunc{
    return func(ctx *router.Context){
        for _,preHandler := range preHandlers{
            err := preHandler(ctx)
            if nil != err{
                log.Println(err)
                return
            }
        }

        fun(ctx)
    }
}
