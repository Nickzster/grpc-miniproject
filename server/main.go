package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Nickzster/grpc-miniproject/todo"
	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Starting TODO Server!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := todo.Server{}

	grpcServer := grpc.NewServer()

	todo.RegisterTodoServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}