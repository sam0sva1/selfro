package router

import (
	"net/http"
)

type Router struct {
	status   string
	routes   map[string]map[string]http.HandlerFunc
	routeReg []string
}

func Start() Router {
	return Router{
		routes:   make(map[string]map[string]http.HandlerFunc),
		routeReg: make([]string, 0),
	}
}

func (r *Router) Get(route string, handler http.HandlerFunc) {
	r.routeReg = append(r.routeReg, route)
	r.routes[route] = make(map[string]http.HandlerFunc)
	r.routes[route][http.MethodGet] = handler
}

func buildRoute(pattern string, pack map[string]http.HandlerFunc) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if _, ok := pack[r.Method]; ok {
			pack[r.Method](w, r)
		}
	})
}

func (router Router) RegisterRoutes() {
	for _, val := range router.routeReg {
		if _, ok := router.routes[val]; ok {
			buildRoute(val, router.routes[val])
		}
	}
}

func (r Router) Run() {
	r.RegisterRoutes()
	http.ListenAndServe(":8080", nil)
}
