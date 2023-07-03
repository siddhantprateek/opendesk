package rpcclient

import (
	"context"

	"github.com/siddhantprateek/opendesk/configs"
	pb "github.com/siddhantprateek/opendesk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	AUTH_RPC_PORT = configs.GetEnv("AUTH_RPC_PORT")
)

func AuthClient() (pb.LoginSerivceClient, *grpc.ClientConn, error) {
	conn, err := grpc.DialContext(context.Background(),
		AUTH_RPC_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	authClient := pb.NewLoginSerivceClient(conn)

	return authClient, conn, nil
}
