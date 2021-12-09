package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	pb "userBikeService/user"
)

const port = ":30033"

func main() {
	s := grpc.NewServer()
	r := NewRepository()
	h := &Handler{&AuthorizationService{r}, &UserCreatorService{r},
		&TokenMaker{r}, r, pb.UnimplementedUserServiceServer{}}

	pb.RegisterUserServiceServer(s, h)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve", err)
	}
}
