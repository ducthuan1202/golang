package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"controller"
)

type Method string

const (
	GET Method = "GET"
	POST Method = "POST"
	PUT Method = "PUT"
	PATCH Method = "PATCH"
	DELETE Method = "DELETE"
)

type Route struct {
	handler http.HandlerFunc
	method []Method
	name string
}

type MyRoute struct {
	RootRoute *mux.Router
	Routers map[string]Route
}

var AppRoute MyRoute 

func (mr *MyRoute)Register() {	
	for key, val := range mr.Routers {		
		var methods []string
		for _, m := range val.method {
			methods = append(methods, string(m))
		}		
		mr.RootRoute.HandleFunc(key, val.handler).Methods(methods...).Name(val.name)
	}
}

func (mr *MyRoute)Initial(){

	mr.RootRoute = mux.NewRouter()

	mr.Routers = map[string] Route {
		"/": Route{ 
			method: []Method{ GET }, 
			handler: controller.HomepageHandler,
			name: "home",
		},
		"/products": Route{ 
			method: []Method{ GET }, 
			handler: controller.IndexProductHandler,
			name: "products.index",
		},
		"/products/create": Route{ 
			method: []Method{ POST }, 
			handler: controller.AddProductHandler,
			name: "products.add",
		},
		"/products/{product_id:[0-9]+}": Route{ 
			method: []Method{ GET }, 
			handler: controller.DetailProductHandler,
			name: "products.detail",
		},
		"/articles": Route{ 
			method: []Method{ GET }, 
			handler: controller.IndexArticleHandler,
			name: "articles.index",
		},
		"/articles/{article_id:[0-9]+}": Route{ 
			method: []Method{ GET }, 
			handler: controller.DetailArticleHandler,
			name: "articles.detail",
		},
	}	

	mr.Register()
}

func (mr *MyRoute)IsEmpty()bool  {
	if len(mr.Routers) > 0 {
		return false
	}
	return true
}