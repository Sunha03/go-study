package user

import (
	handler "go-study/user/handler"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

var r *mux.Router

func New() *mux.Router {
	r = mux.NewRouter()

	return r
}

func SwaggerHandler() {
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

func UserHandler() {
	h := r.PathPrefix("/user").Subrouter()
	h.HandleFunc("", handler.GetUserList).Methods("GET")
	h.HandleFunc("", handler.CreateUser).Methods("POST")
	h.HandleFunc("/{userId}", handler.GetUser).Methods("GET")
	h.HandleFunc("/{userId}", handler.UpdateUser).Methods("PUT")
	h.HandleFunc("/{userId}", handler.DeleteUser).Methods("DELETE")
}
