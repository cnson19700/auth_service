package grpc

import (
	"context"

	"github.com/cnson19700/auth_service/proto"
	"github.com/cnson19700/auth_service/usecase/auth"
)

// AuthService .
type AuthService struct {
	proto.UnimplementedAuthServiceServer

	AuthUsecase auth.IUsecase
}

func (as *AuthService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {

	token, err := as.AuthUsecase.Login(ctx, auth.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	return &proto.LoginResponse{
		Token: token,
	}, err
}
