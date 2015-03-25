package router

import (
	"log"
	"net/http"
)

type (
	HandlerFunc func(*Context)
	router      struct{}
)

var (
	handlers = map[string]HandlerFunc{}
	_router  = &router{}
)

func Router() *router {
	return _router
}

func (this *router) Register(path string, h HandlerFunc) {
	handlers[path] = h
}

func (this *router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	ctx := NexContext(rw, req)
	path := req.RequestURI
	h, ok := handlers[path]
	if !ok {
		log.Println("uri not found")
		rw.WriteHeader(404)
		rw.Write([]byte("uri not found"))
		return
	}

	h(ctx)
}
