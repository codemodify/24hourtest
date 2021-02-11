package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/codemodify/24hourtest/demo-workspace/contracts"
	"google.golang.org/grpc"
)

func TestConnection(t *testing.T) {
	// 1. connect to server
	netConnection, err := grpc.Dial(contracts.ServerAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to connect: %v", err))
		t.FailNow()
	}
	defer netConnection.Close()

	// 2. call method
	endpointClient := contracts.NewHelloWorldServiceClient(netConnection)

	ctx, cancel := context.WithTimeout(context.Background(), contracts.ConnectionDefaultTimeout)
	defer cancel()

	endpointReply, err := endpointClient.HelloWorld(ctx, &contracts.HelloWorldRequest{})
	if err != nil {
		fmt.Println(fmt.Sprintf("error on making RPC call: %v", err))
		t.FailNow()
	}

	// 3. consume reply
	fmt.Println(fmt.Sprintf("DEBUG: got: %s", endpointReply.HelloWorldPayload))

	if endpointReply.HelloWorldPayload != contracts.HelloWorldConstantMessage {
		t.FailNow()
	}
}
