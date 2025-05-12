package controller

import (
	"context"

	"github.com/soner3/evently/model"
	eventv1 "github.com/soner3/evently/proto/gen/event/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventController struct {
	eventv1.UnimplementedEventServiceServer
}

func (e *EventController) CreateEvent(_ context.Context, req *eventv1.CreateEventRequest) (*eventv1.CreateEventResponse, error) {
	event := model.NewEvent(req.GetName(), req.GetDescription(), req.GetLocation(), req.GetDateTime().AsTime())
	return &eventv1.CreateEventResponse{
		EventId:     event.EventId.String(),
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    timestamppb.New(event.DateTime),
		UserId:      event.UserId.String(),
	}, nil
}
