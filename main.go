package main

import (
	"go-study/grpcserver"
	"go-study/study"
	"go-study/test"
)

func main() {
	// test
	test.TestMain()

	// study
	study.StudyMain()

	// grpcserver
	grpcserver.GrpcServerMain()
}
