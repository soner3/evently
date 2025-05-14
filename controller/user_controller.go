package controller

import (
	"context"

	"github.com/soner3/evently/model"
	userv1 "github.com/soner3/evently/proto/gen/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserController struct {
	userv1.UnimplementedUserServiceServer
}

func (us *UserController) CreateUser(ctx context.Context, req *userv1.CreateUserRequest) (*userv1.CreateUserResponse, error) {
	user := model.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
	if err := user.Save(); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not save user: %v", err)
	}
	return &userv1.CreateUserResponse{
		UserId: user.UserId.String(),
		Email:  user.Email,
	}, nil
}

// func (uc *UserController) DeleteUser(ctx context.Context, req *userv1.DeleteUserRequest) (*userv1.DeleteUserResponse, error) {
// 	user := model.User{
// 		UserId: uuid.MustParse(req.),
// 	}
// 	if err := user.DeleteUserById(); err != nil {
// 		return nil, status.Errorf(codes.Internal, "Could not delete user: %v", err)
// 	}
// 	return &userv1.DeleteUserResponse{
// 		Message: "Deleted",
// 	}, nil
// }

// func (uc *UserController) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
// 	user := model.User{}

// }
