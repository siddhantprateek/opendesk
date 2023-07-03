package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	AuthGrpc "github.com/siddhantprateek/opendesk/api_gateway/rpcclient"
	"golang.org/x/crypto/bcrypt"

	"github.com/siddhantprateek/opendesk/pb"
)

// Authentication Handlers
func AuthInit(e echo.Context) error {
	return e.JSON(http.StatusOK, echo.Map{
		"status":  "200",
		"message": "Authentication gRPC calls routes.",
	})
}

func CreateUser(e echo.Context) error {
	user := new(pb.CreateUserRequest)
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "Unable to create User",
		})
	}

	authClient, conn, err := AuthGrpc.AuthClient()
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": "connection to auth gRPC service failed",
		})
	}
	defer conn.Close()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Unable to hash the password.")
	}
	req := &pb.CreateUserRequest{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	res, err := authClient.CreateUser(context.Background(), req)
	if err != nil {
		return e.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusCreated, res)
}

func UpdateUser(e echo.Context) error {
	return nil
}

func GetUser(e echo.Context) error {
	return nil
}

func GetAllUser(e echo.Context) error {
	return nil
}

func Login(e echo.Context) error {
	return nil
}

func DeleteUser(e echo.Context) error {
	return nil
}

func ForgetPassword(e echo.Context) error {
	return nil
}
