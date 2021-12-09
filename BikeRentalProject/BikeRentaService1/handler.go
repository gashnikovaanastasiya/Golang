package main

import (
	pb "BikeRentaService1/bike"
	pb1 "BikeRentaService1/user"
	"context"
	"google.golang.org/grpc/metadata"
	"strconv"
)

const (
	costPerHour = 2
)

type GrpcError struct {
	Message string
}

func (a *GrpcError) Error() string {
	return a.Message
}

type Handler struct {
	pb.UnimplementedBikeRentalServiceServer
	pb1.UserServiceClient
	repository
}

func (h *Handler) CreateBike(ctx context.Context, in *pb.Address) (*pb.Response, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, &GrpcError{Message: "parsing context error"}
	}
	role := meta.Get("Role")[0]
	if role != "admin" {
		return nil, &GrpcError{Message: "Access denied."}
	}
	res, err := h.repository.CreateBike(in.Address)
	return &pb.Response{Response: res}, err
}
func (h *Handler) GetAll(ctx context.Context, empty *pb.Empty) (*pb.Response, error) {
	bikes, err := h.repository.GetAll()
	if err != nil {
		return nil, err
	}
	response := "All bikes:\n"
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, &GrpcError{Message: "parsing context error"}
	}
	role := meta.Get("Role")[0]
	switch role {
	case "admin":
		for _, bike := range bikes {
			response += bike.String()
		}
	case "user":
		for _, bike := range bikes {
			response += bike.ShowAvailable()
		}
	default:
		response = "Please sign in to use the app"
	}
	return &pb.Response{Response: response}, nil
}
func (h *Handler) RentBike(ctx context.Context, in *pb.BikeRequest) (*pb.Response, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, &GrpcError{Message: "parsing context error"}
	}
	role := meta.Get("Role")[0]
	if role != "user" {
		return nil, &GrpcError{Message: "Access denied"}
	}
	id := meta.Get("Id")[0]
	Id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	res, err := h.repository.RentBike(Id, int(in.BikeId))
	return &pb.Response{Response: res}, err
}
func (h *Handler) ReturnBike(ctx context.Context, in *pb.BikeRequest) (*pb.Response, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, &GrpcError{Message: "parsing context error"}
	}
	role := meta.Get("Role")[0]
	if role != "user" {
		return nil, &GrpcError{Message: "Access denied"}
	}
	id := meta.Get("Id")[0]
	Id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	hours, err := h.repository.ReturnBike(Id, int(in.BikeId), in.Address)
	if err != nil {
		return nil, err
	}
	header := metadata.New(map[string]string{"Role": "user"})
	ctx = metadata.NewOutgoingContext(context.Background(), header)
	_, err = h.UserServiceClient.SetBalance(ctx, &pb1.SetBalanceInfo{Id: int32(Id),
		Money: int32(-1 * costPerHour * hours)})
	if err != nil {
		return nil, err
	}
	return &pb.Response{Response: "Bike was returned"}, nil
}
