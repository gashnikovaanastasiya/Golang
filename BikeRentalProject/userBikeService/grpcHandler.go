package main

import (
	"context"
	"google.golang.org/grpc/metadata"
	pb "userBikeService/user"
)

type GrpcError struct {
	Message string
}

func (a *GrpcError) Error() string {
	return a.Message
}

type Handler struct {
	authorization
	withHashing
	tokenMaker
	repository
	pb.UnimplementedUserServiceServer
}

func (h *Handler) CreateUser(ctx context.Context, in *pb.User) (*pb.Response, error) {
	str, err := h.withHashing.CreateUser(&User{Name: in.Name, Password: in.Password, Balance: int(in.Balance)})
	return &pb.Response{Response: str}, err
}
func (h *Handler) SignIn(ctx context.Context, in *pb.SignInInfo) (*pb.Token, error) {
	role, id, err := h.authorization.Authorize(in.Name, in.Password)
	if err != nil {
		return nil, err
	}

	token, err := h.tokenMaker.CreateToken(id, role)
	if err != nil {

		return nil, err
	}
	return &pb.Token{Token: token}, nil
}
func (h *Handler) SetBalance(ctx context.Context, in *pb.SetBalanceInfo) (*pb.Response, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, &GrpcError{Message: "parsing context error"}
	}

	role := meta.Get("Role")[0]
	if role != "user" {
		return nil, &GrpcError{Message: "Access denied"}
	}
	str, err := h.repository.SetBalance(int(in.Id), int(in.Money))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: str}, nil
}
func (h *Handler) GetAll(ctx context.Context, in *pb.Empty) (*pb.Users, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, &GrpcError{Message: "parsing context error"}
	}
	role := meta.Get("Role")[0]
	if role != "admin" {
		return nil, &GrpcError{Message: "Access denied"}
	}
	users, err := h.repository.GetAll()
	if err != nil {
		return nil, err
	}
	protoUsers := make([]*pb.User, 0)
	for _, user := range users {
		protoUsers = append(protoUsers, &pb.User{Id: int32(user.UserId), Name: user.Name,
			Password: user.Password, Balance: int32(user.Balance)})
	}
	return &pb.Users{User: protoUsers}, nil
}
func (h *Handler) DeleteUser(ctx context.Context, in *pb.Id) (*pb.Response, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, &GrpcError{Message: "parsing context error"}
	}
	role := meta.Get("Role")[0]
	if role != "admin" {
		return nil, &GrpcError{Message: "Access denied"}
	}
	str, err := h.repository.DeleteUser(int(in.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: str}, nil
}
