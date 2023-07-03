package main

import (
	"context"
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

	_ = &pb.CreateUserRequest{
		Name:     name,
		Username: username,
		Email:    email,
		Password: password,
	}

	return &pb.CreateUserResponse{Status: 1, Message: "User Created"}, nil
}

func (s *server) UserUpdate(ctx context.Context,
	req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {

	return &pb.UpdateUserResponse{
		Success: true,
		Message: "User Data has been Updated."}, nil
}

func (s *server) ForgetPassword(ctx context.Context,
	req *pb.ForgetPasswordRequest) (*pb.ForgetPasswordResponse, error) {

	return &pb.ForgetPasswordResponse{
		Success:     true,
		Message:     "Password has been Updated.",
		NewPassword: "NewPassword."}, nil
}

func GetSingleUser(ctx context.Context,
	req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	return &pb.GetUserResponse{}, nil
}

func GetAllUsers(ctx context.Context,
	req *pb.GetAllUserRequest) (*pb.GetAllUsersResponse, error) {

	return &pb.GetAllUsersResponse{}, nil
}

func DeleteUser(ctx context.Context,
	req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {

	return &pb.DeleteUserResponse{}, nil
}

func Login(ctx context.Context,
	req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return &pb.UserLoginResponse{}, nil
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
