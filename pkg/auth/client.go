package client

import (
	"fmt"

	"github.com/nathannlu/microservice/pkg/auth/pb"
	"github.com/nathannlu/microservice/pkg/config"
	"google.golang.org/grpc"
)


type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {

	// Start grpc connection
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.withInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	
	return pb.NewAuthServiceClient(cc)
}
