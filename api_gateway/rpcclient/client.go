package rpcclient

import (
	"context"

	"github.com/siddhantprateek/opendesk/configs"
	pb "github.com/siddhantprateek/opendesk/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// AUTH_RPC_PORT is the environment variable used to
// specify the port on which the Auth service's RPC server is running.
var (
	AUTH_RPC_PORT = configs.GetEnv("AUTH_RPC_PORT")
)

/*
AuthClient establishes a gRPC connection with the
Auth service and returns the client and connection instances.

The client can be used to make RPC calls to the Auth service's defined methods.
Example usage:

	  client, conn, err := AuthClient()
	  if err != nil {
			Handle error
	  }
	  defer conn.Close()

Use the 'client' to make RPC calls to the Auth service.
Parameters:

	None

Returns:
  - pb.LoginSerivceClient: A client instance that implements
    the LoginSerivceClient interface to interact with the Auth service.
  - *grpc.ClientConn: A gRPC client connection to the Auth service.
    It should be closed after use.
  - error: An error, if any, encountered during the connection establishment.
    It will be nil if the connection is successful.

Note:
  - The 'AUTH_RPC_PORT' environment variable should be set
    to the correct port number where the Auth service's RPC server is running.
  - This function uses insecure credentials to connect to the Auth
    service. In production, it's essential to use secure credentials.
*/
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
