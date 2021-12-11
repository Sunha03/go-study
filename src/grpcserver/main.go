package main

import (
	"log"
	"net"
	"context"

	"grpcserver/data"
	userpb "grpcserver/pb"
	"google.golang.org/grpc"
)

const portNumber = "8000" // port 번호

type userServer struct {
	userpb.UserServer
}

func main() {
	listen, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	grpcServer := grpc.NewServer()		// User service를 GRPC server에 등록
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("Start GRPC server on %s port", portNumber)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to Serve : %s", err)
	}
}

// GetUser : user 상세 조회
func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userId := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.UserData {
		if u.UserId != userId {
			continue
		} else {
			userMessage = u
			break
		}
	}

	return &userpb.GetUserResponse {
		UserMessage: userMessage,
	}, nil
}

// ListUsers : user 리스트 조회
func (s *userServer) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessages := make([]*userpb.UserMessage, len(data.UserData))
	for i, user := range data.UserData {
		userMessages[i] = user
	}

	return &userpb.ListUsersResponse {
		UserMessages: userMessages,
	}, nil
}