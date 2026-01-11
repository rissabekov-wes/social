package api

// implement http.Handler that returns 200 with application/json content type and ok

import (
	"net/http"

	"github.com/Wesfarmers-Digital/pkg/one_http"
)

const (
	httpMethod = "GET"
	httpPath   = "/example"
)

// type ApiHandlerExample struct{}

// func NewApiHandlerExample() *ApiHandlerExample {
// 	return &ApiHandlerExample{}
// }

// func (h *ApiHandlerExample) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"status":"ok"}`))
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func ConfigRoute() one_http.Route {
	return one_http.Route{
		Method:  httpMethod,
		Path:    httpPath,
		Handler: http.HandlerFunc(Handler),
	}
}
