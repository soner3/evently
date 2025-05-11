package routes

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soner3/evently/controller"
	eventv1 "github.com/soner3/evently/proto/gen/event/v1"
	"google.golang.org/grpc"
)

func InitGrpcRoutes(grpcServer *grpc.Server) {
	eventv1.RegisterEventServiceServer(grpcServer, &controller.EventController{})
}

func InitRestRoutes(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	if err := eventv1.RegisterEventServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}
	return nil

}
