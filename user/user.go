package user

import (
	"log"
	"net/http"

	router "go-study/user/router"
)

const port = ":12000"

func UserMain() {
	r := router.New()

	router.UserHandler()

	log.Fatal(http.ListenAndServe(port, r))
}
