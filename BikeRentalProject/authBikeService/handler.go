package main

import (
	pb "authBikeService/auth"
	"context"
)

type Handler struct {
	pb.UnimplementedAuthenticationServiceServer
	authentication
}

func (h *Handler) ParseToken(ctx context.Context, in *pb.Token) (*pb.Response, error) {
	claims, err := h.decode(in.Token)
	if err != nil {
		return nil, err
	}
	return &pb.Response{Id: int32(claims.Id), Role: claims.Role}, nil
}
