package main

import (
	"context"
	"fmt"
	"log"
	"net"

	configs "github.com/siddhantprateek/opendesk/configs"
	"github.com/siddhantprateek/opendesk/pb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	authGRPCservice = configs.GetEnv("AUTH_RPC_PORT")
)

type server struct {
	pb.UnimplementedLoginSerivceServer
}

func (s *server) CreateUser(ctx context.Context,
	req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	name := req.Name
	username := req.Username
	email := req.Email
	password := req.Password

	fmt.Println("\nResquest Recieved.")
	fmt.Println("name: ", name)
	fmt.Println("username: ", username)
	fmt.Println("email: ", email)
	fmt.Println("password: ", password)

	return &pb.CreateUserResponse{Status: 1, Message: "User Created"}, nil
}

func AuthRPCserver() {
	rpcServer := grpc.NewServer()
	reflection.Register(rpcServer)
	listener, err := net.Listen("tcp", authGRPCservice)
	if err != nil {
		log.Fatal("Cannot create listener.")
	}

	logrus.Infof("Starting gRPC server at %s", listener.Addr().String())
	pb.RegisterLoginSerivceServer(rpcServer, new(server))
	if err := rpcServer.Serve(listener); err != nil {
		log.Fatal("Error: ", err)
	}

}

func main() {
	AuthRPCserver()
}
