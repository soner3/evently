package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/soner3/evently/model"
	eventv1 "github.com/soner3/evently/proto/gen/event/v1"
	"github.com/soner3/evently/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventController struct {
	eventv1.UnimplementedEventServiceServer
}

func (e *EventController) CreateEvent(ctx context.Context, req *eventv1.CreateEventRequest) (*eventv1.CreateEventResponse, error) {
	token, _ := util.ExtractTokenFromHeader(&ctx)
	_, claims, _ := util.ValidateToken(*token)
	userId := claims.Subject

	event := model.Event{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Location:    req.GetLocation(),
		DateTime:    req.GetDateTime().AsTime(),
		User:        model.User{UserId: uuid.MustParse(userId)},
	}
	if err := event.Save(); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not save event: %v", err)
	}
	return &eventv1.CreateEventResponse{
		EventId:     event.EventId.String(),
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    timestamppb.New(event.DateTime),
		UserId:      event.User.UserId.String(),
	}, nil
}

func (e *EventController) DeleteEvent(ctx context.Context, req *eventv1.DeleteEventRequest) (*eventv1.DeleteEventResponse, error) {
	event := model.Event{}
	event.EventId = uuid.MustParse(req.GetEventId())
	if err := event.DeleteById(); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not delete event: %v", err)
	}
	return &eventv1.DeleteEventResponse{
		Message: "Deleted",
	}, nil
}

func (e *EventController) GetEvent(ctx context.Context, req *eventv1.GetEventRequest) (*eventv1.GetEventResponse, error) {
	eventModel := model.Event{}
	event, err := eventModel.FindById(uuid.MustParse(req.GetEventId()))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "no event found: %v", err)
	}
	return &eventv1.GetEventResponse{
		EventId:     event.EventId.String(),
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    timestamppb.New(event.DateTime),
		UserId:      event.User.UserId.String(),
	}, nil

}

func (e *EventController) UpdateEvent(ctx context.Context, req *eventv1.UpdateEventRequest) (*eventv1.UpdateEventResponse, error) {
	token, _ := util.ExtractTokenFromHeader(&ctx)
	_, claims, _ := util.ValidateToken(*token)
	userId := claims.Subject

	event := model.Event{}
	event.EventId = uuid.MustParse(req.GetEventId())
	event.Name = req.GetName()
	event.Description = req.GetDescription()
	event.Location = req.GetDescription()
	event.DateTime = req.GetDateTime().AsTime()
	event.User.UserId = uuid.MustParse(userId)
	if err := event.Save(); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not save event: %v", err)
	}

	return &eventv1.UpdateEventResponse{
		EventId:     event.EventId.String(),
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    timestamppb.New(event.DateTime),
		UserId:      event.User.UserId.String(),
	}, nil

}
