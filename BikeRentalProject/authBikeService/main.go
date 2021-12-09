package main

import (
	pb "authBikeService/auth"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	port = ":30032"
)

func main() {
	s := grpc.NewServer()
	h := &Handler{pb.UnimplementedAuthenticationServiceServer{},
		&AuthenticationService{}}
	pb.RegisterAuthenticationServiceServer(s, h)
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
