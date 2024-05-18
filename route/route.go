package route

import (
	"net/http"

	"RegistrationSyetem/controller"
	"RegistrationSyetem/dto"
	"RegistrationSyetem/services"

	"github.com/gorilla/mux"
)

type RouteType struct {
	Method     string
	Pattern    string
	Handle     http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []RouteType

func init() {
	register("POST", "/user/createUser", controller.CreateUser, nil)
	register("POST", "/calendar/createClendar", controller.CreateUser, nil)
	register("GET", "/test", Test, nil)
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	for _, route := range routes {
		r.Methods(route.Method).Path(route.Pattern).Handler(route.Handle)

		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}

	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, RouteType{method, pattern, handler, middleware})
}

func Test(w http.ResponseWriter, r *http.Request) {

	response := dto.ApiResponse{ResultCode: "200", ResultMessage: "hello world"}
	services.ResponseWithJson(w, http.StatusOK, response)
}
