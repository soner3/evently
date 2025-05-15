package controller

import (
	"context"

	"github.com/soner3/evently/model"
	authv1 "github.com/soner3/evently/proto/gen/auth/v1"
	"github.com/soner3/evently/util"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthController struct {
	authv1.UnimplementedAuthServiceServer
}

func (ac *AuthController) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	user := model.User{Email: req.Email}
	if err := user.FindUserByEmail(true); err != nil {
		return nil, status.Error(codes.NotFound, "No user found for this email")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword())); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Wrong password")
	}
	token, err := util.GenerateToken(user.Email, user.UserId.String())
	if err != nil {
		return nil, status.Error(codes.Internal, "Could not create token")
	}

	return &authv1.LoginResponse{
		AccessToken: token,
	}, nil

}
