package user

import (
	"log"
	"net/http"

	router "go-study/user/router"
)

const port = ":12000"

func UserMain() {
	r := router.New()

	router.SwaggerHandler()
	router.UserHandler()

	log.Println("===================== user server start (", port, ") =====================")
	log.Fatal(http.ListenAndServe(port, r))
}
