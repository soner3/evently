package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/soner3/evently/model"
	eventv1 "github.com/soner3/evently/proto/gen/event/v1"
	userv1 "github.com/soner3/evently/proto/gen/user/v1"
	"github.com/soner3/evently/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (uc *UserController) GetUser(ctx context.Context, req *userv1.GetUserRequest) (*userv1.GetUserResponse, error) {
	token, _ := util.ExtractTokenFromHeader(&ctx)
	_, claims, _ := util.ValidateToken(*token)
	userId := claims.Subject
	user := model.User{UserId: uuid.MustParse(userId)}
	if err := user.FindUserByIdWithEvents(); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not find user: %v", err)
	}

	events := make([]*eventv1.Event, len(*user.Events))

	for i, e := range *user.Events {
		events[i] = &eventv1.Event{
			EventId:     e.EventId.String(),
			Name:        e.Name,
			Description: e.Description,
			Location:    e.Location,
			DateTime:    timestamppb.New(e.DateTime),
			UserId:      userId,
		}

	}

	return &userv1.GetUserResponse{
		UserId: userId,
		Email:  user.Email,
		Events: events,
	}, nil

}
