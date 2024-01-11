package handlers

import (
	"context"
	"errors"
	pb "github.com/raedmajeed/booking-service/pkg/pb"
	"github.com/raedmajeed/booking-service/pkg/utils"
	"log"
	"time"
)

func (handler *BookingHandler) RegisterLoginRequest(ctx context.Context, p *pb.LoginRequest) (*pb.LoginResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	token, err := handler.svc.UserLogin(p)
	if err != nil {
		log.Printf("Unable to login %v of email == %v, err: %v", p.Role, p.Email, err.Error())
		return nil, err
	}
	if token == "" {
		log.Printf("Unable to login %v of email + %v, err: %v", p.Role, p.Email, err.Error())
		return nil, err
	}
	return utils.ConvertLoginRequestToResponse(token, p), nil
}

func (handler *BookingHandler) RegisterUser(ctx context.Context, p *pb.UserRequest) (*pb.UserResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := handler.svc.RegisterUserSvc(p)
	if err != nil {
		log.Printf("Unable to register of email == %v, err: %v", p.Email, err.Error())
		return nil, err
	}
	return &pb.UserResponse{
		Email: resp.Email,
	}, nil
}
func (handler *BookingHandler) VerifyUser(ctx context.Context, p *pb.OTPRequest) (*pb.UserResponse, error) {
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(time.Now()) {
		log.Println("deadline passed, aborting gRPC call")
		return nil, errors.New("deadline passed, aborting gRPC call")
	}

	resp, err := handler.svc.VerifyUserRequest(p)
	if err != nil {
		log.Printf("Unable to verify %v of email == %v, err: %v", p.Email, err.Error())
		return nil, err
	}
	return &pb.UserResponse{
		Email: resp.Email,
	}, nil
}
