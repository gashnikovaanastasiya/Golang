package main

import (
	pb3 "BikeServiceClient/auth"
	pb2 "BikeServiceClient/bike"
	pb1 "BikeServiceClient/user"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
	"net/http"
	"strconv"
	"time"
)

type EchoServer struct {
	*echo.Echo
	pb1.UserServiceClient
	pb2.BikeRentalServiceClient
	pb3.AuthenticationServiceClient
}

func (e *EchoServer) Register() {
	e.GET("/createUser", e.createUser)
	e.GET("/users/getAll", e.GetAllUsers)
	e.GET("/setBalance", e.setBalance)
	e.GET("/signIn", e.signIn)
	e.GET("/deleteUser", e.deleteUser)
	e.GET("/bikes/getAll", e.getAllBikes)
	e.GET("/rentBike", e.rentBike)
	e.GET("/returnBike", e.returnBike)
	e.GET("/createBike", e.createBike)
}
func (e *EchoServer) createUser(c echo.Context) error {
	balance, err := strconv.Atoi(c.QueryParam("balance"))
	if err != nil {
		return err
	}
	res, err := e.UserServiceClient.CreateUser(context.Background(), &pb1.User{Name: c.QueryParam("name"),
		Password: c.QueryParam("password"), Balance: int32(balance)})
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, res.Response)
}
func (e *EchoServer) GetAllUsers(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	authResp, err := e.AuthenticationServiceClient.ParseToken(context.Background(), &pb3.Token{
		Token: token,
	})
	if err != nil {
		return err
	}
	header := metadata.New(map[string]string{"Role": authResp.Role})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	res, err := e.UserServiceClient.GetAll(ctx, &pb1.Empty{})
	if err != nil {
		return c.String(http.StatusExpectationFailed, err.Error())
	}
	return c.String(http.StatusOK, res.String())
}
func (e *EchoServer) setBalance(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	authResp, err := e.AuthenticationServiceClient.ParseToken(context.Background(), &pb3.Token{
		Token: token,
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(authResp)
	header := metadata.New(map[string]string{"Role": authResp.Role})
	fmt.Println(header)
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	money, err := strconv.Atoi(c.QueryParam("money"))
	if err != nil {
		return err
	}
	res, err := e.UserServiceClient.SetBalance(ctx, &pb1.SetBalanceInfo{Id: authResp.Id,
		Money: int32(money)})
	if err != nil {
		return c.String(http.StatusExpectationFailed, err.Error())
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) signIn(c echo.Context) error {
	res, err := e.UserServiceClient.SignIn(context.Background(), &pb1.SignInInfo{Name: c.QueryParam("name"),
		Password: c.QueryParam("password")})
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "token", Value: "", HttpOnly: true, Expires: time.Now().Add(10 * time.Minute)})
		return c.String(http.StatusExpectationFailed, err.Error())
	}
	c.SetCookie(&http.Cookie{Name: "token", Value: res.Token, HttpOnly: true, Expires: time.Now().Add(10 * time.Minute)})
	return c.String(http.StatusAccepted, "You entered successfully")
}
func (e *EchoServer) deleteUser(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	authResp, err := e.AuthenticationServiceClient.ParseToken(context.Background(), &pb3.Token{
		Token: token,
	})
	if err != nil {
		return err
	}
	header := metadata.New(map[string]string{"Role": authResp.Role})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	Id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}
	res, err := e.UserServiceClient.DeleteUser(ctx, &pb1.Id{Id: int32(Id)})
	if err != nil {
		return c.String(http.StatusExpectationFailed, err.Error())
	}
	return c.String(http.StatusOK, res.Response)
}

func (e *EchoServer) createBike(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	authResp, err := e.AuthenticationServiceClient.ParseToken(context.Background(), &pb3.Token{
		Token: token,
	})
	if err != nil {
		return err
	}
	header := metadata.New(map[string]string{"Role": authResp.Role})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	res, err := e.BikeRentalServiceClient.CreateBike(ctx, &pb2.Address{Address: c.QueryParam("address")})
	if err != nil {
		return c.String(http.StatusExpectationFailed, err.Error())
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) getAllBikes(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	authResp, err := e.AuthenticationServiceClient.ParseToken(context.Background(), &pb3.Token{
		Token: token,
	})
	if err != nil {
		return err
	}
	header := metadata.New(map[string]string{"Role": authResp.Role})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	res, err := e.BikeRentalServiceClient.GetAll(ctx, &pb2.Empty{})
	if err != nil {
		return c.String(http.StatusExpectationFailed, err.Error())
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) rentBike(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	authResp, err := e.AuthenticationServiceClient.ParseToken(context.Background(), &pb3.Token{
		Token: token,
	})
	if err != nil {
		return err
	}
	header := metadata.New(map[string]string{"Role": authResp.Role,
		"Id": strconv.Itoa(int(authResp.Id))})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	Id, err := strconv.Atoi(c.QueryParam("id"))
	res, err := e.BikeRentalServiceClient.RentBike(ctx, &pb2.BikeRequest{BikeId: int32(Id),
		Address: c.QueryParam("address")})
	if err != nil {
		return c.String(http.StatusExpectationFailed, err.Error())
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) returnBike(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use service")
	}
	token := cookies.Value
	authResp, err := e.AuthenticationServiceClient.ParseToken(context.Background(), &pb3.Token{
		Token: token,
	})
	if err != nil {
		return err
	}
	header := metadata.New(map[string]string{"Role": authResp.Role,
		"Id": strconv.Itoa(int(authResp.Id))})
	ctx := metadata.NewOutgoingContext(context.Background(), header)
	Id, err := strconv.Atoi(c.QueryParam("id"))
	res, err := e.BikeRentalServiceClient.ReturnBike(ctx, &pb2.BikeRequest{BikeId: int32(Id),
		Address: c.QueryParam("address")})
	if err != nil {
		return c.String(http.StatusExpectationFailed, err.Error())
	}
	return c.String(http.StatusOK, res.Response)
}
