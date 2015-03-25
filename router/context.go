package router

import(
    "net/http"
    "log"
)

type Context struct{
    http.ResponseWriter
    Request *http.Request
}

var(
    params = map[string]string{}
)

func NexContext(rw http.ResponseWriter,req *http.Request)*Context{
    return &Context{rw,req}
}

func (this *Context)Params(key string)string{
    err := this.Request.ParseForm()
    if nil != err{
        log.Println(err)
        return ""
    }

    return ""
}
