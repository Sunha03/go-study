package docker

import (
	"fmt"
	"log"
	"net/http"
)

func DockerMain() {
	// HTTP 요청에 대해 "Hello World"라는 응답 보내기
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request") // 클라이언트로부터 요청 받으면 해당 메세지를 표준으로 출력
		fmt.Fprintf(w, "Hello World")
	})

	log.Println("start server")
	// 8000 포트로 요청받음
	server := &http.Server{
		Addr: ":8000",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
