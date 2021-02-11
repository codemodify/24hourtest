package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/codemodify/24hourtest/demo-workspace/contracts"
	"google.golang.org/grpc"
)

type helloWorldGrpcServer struct {
	contracts.UnimplementedHelloWorldServiceServer
}

func (s *helloWorldGrpcServer) HelloWorld(ctx context.Context, in *contracts.HelloWorldRequest) (*contracts.HelloWorldReply, error) {
	fmt.Println(fmt.Sprintf("DEBUG: sending %s", contracts.HelloWorldConstantMessage))

	return &contracts.HelloWorldReply{
		HelloWorldPayload: contracts.HelloWorldConstantMessage,
	}, nil
}

func main() {

	// 1. create server
	s := grpc.NewServer()
	contracts.RegisterHelloWorldServiceServer(s, &helloWorldGrpcServer{})

	// 2. serve endpoint
	netListener, err := net.Listen("tcp", contracts.ServerAddress)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to bind to port: %v", err))
		os.Exit(1)
	}

	err = s.Serve(netListener)
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(1)
	}
}
