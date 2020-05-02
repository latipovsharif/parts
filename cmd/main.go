package main

import (
	"fmt"
	"net"
	"parts/pkg/services"

	"google.golang.org/grpc"
)

func main() {
	partService := services.Part{}
	server := grpc.NewServer()
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(fmt.Sprintf("cannot listen on 9090: %s", err))
	}

	services.RegisterPartServer(server, &partService)

	if err := server.Serve(listener); err != nil {
		panic(fmt.Sprintf("cannot serve on localhost 9090"))
	}
}
