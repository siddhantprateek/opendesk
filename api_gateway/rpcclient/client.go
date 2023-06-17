package rpcclient

import (
	"context"
	"errors"

	"github.com/siddhantprateek/opendesk/configs"
	pb "github.com/siddhantprateek/opendesk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	authGRPCservice = configs.GetEnv("AUTH_RPC_PORT")
	authGRPCclient  pb.LoginSerivceClient
)

func GrpcAuthClient(ctx *context.Context) error {

	conn, err := grpc.DialContext(*ctx, authGRPCservice, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		authGRPCclient = nil
		return errors.New("connection to auth gRPC service failed")
	}

	if authGRPCclient != nil {
		conn.Close()
		return nil
	}

	authGRPCclient = pb.NewLoginSerivceClient(conn)
	return nil
}
