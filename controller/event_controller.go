package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/soner3/evently/model"
	eventv1 "github.com/soner3/evently/proto/gen/event/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventController struct {
	eventv1.UnimplementedEventServiceServer
}

func (e *EventController) CreateEvent(_ context.Context, req *eventv1.CreateEventRequest) (*eventv1.CreateEventResponse, error) {
	event := model.Event{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Location:    req.GetLocation(),
		DateTime:    req.GetDateTime().AsTime(),
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
		UserId:      event.UserId.String(),
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

func (e *EventController) GetAllEvents(ctx context.Context, req *eventv1.GetAllEventsRequest) (*eventv1.GetAllEventsResponse, error) {
	eventModel := model.Event{}
	events, err := eventModel.ListEvents()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not list events: %v", err)
	}
	resEvents := make([]*eventv1.Event, len(*events))
	for i, e := range *events {
		resEvents[i] = &eventv1.Event{
			EventId:     e.EventId.String(),
			Name:        e.Name,
			Description: e.Description,
			Location:    e.Location,
			DateTime:    timestamppb.New(e.DateTime),
			UserId:      e.UserId.String(),
		}
	}
	return &eventv1.GetAllEventsResponse{
		Events: resEvents,
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
		UserId:      event.UserId.String(),
	}, nil

}

func (e *EventController) UpdateEvent(ctx context.Context, req *eventv1.UpdateEventRequest) (*eventv1.UpdateEventResponse, error) {
	event := model.Event{}
	event.EventId = uuid.MustParse(req.GetEventId())
	event.Name = req.GetName()
	event.Description = req.GetDescription()
	event.Location = req.GetDescription()
	event.DateTime = req.GetDateTime().AsTime()
	event.UserId = uuid.MustParse(req.GetUserId())
	if err := event.Save(); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not save event: %v", err)
	}

	return &eventv1.UpdateEventResponse{
		EventId:     event.EventId.String(),
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    timestamppb.New(event.DateTime),
		UserId:      event.UserId.String(),
	}, nil

}
