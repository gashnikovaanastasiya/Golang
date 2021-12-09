package main

import (
	pb3 "BikeServiceClient/auth"
	pb2 "BikeServiceClient/bike"
	pb1 "BikeServiceClient/user"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
)

const (
	port          = ":30030"
	userAddress   = "localhost:30033"
	authAddress   = "localhost:30032"
	rentalAddress = "localhost:30031"
)

func main() {
	e := echo.New()
	conn1, err := grpc.Dial(userAddress, grpc.WithInsecure())
	conn2, err := grpc.Dial(rentalAddress, grpc.WithInsecure())
	conn3, err := grpc.Dial(authAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	cl1 := pb1.NewUserServiceClient(conn1)
	cl2 := pb2.NewBikeRentalServiceClient(conn2)
	cl3 := pb3.NewAuthenticationServiceClient(conn3)
	es := EchoServer{e, cl1, cl2, cl3}
	es.Register()
	if err := es.Start(port); err != nil {
		fmt.Println(err)
		return
	}
}
