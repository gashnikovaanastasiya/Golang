package main

import (
	pb1 "BikeRentaService1/bike"
	pb2 "BikeRentaService1/user"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	userAddress = "localhost:30033"
	port        = ":30031"
)

func main() {
	s := grpc.NewServer()
	r := NewRepository()
	conn, err := grpc.Dial(userAddress, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	cl := pb2.NewUserServiceClient(conn)
	handler := &Handler{pb1.UnimplementedBikeRentalServiceServer{}, cl, r}
	pb1.RegisterBikeRentalServiceServer(s, handler)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println(err)
		return
	}
}
