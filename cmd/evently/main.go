package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"buf.build/go/protovalidate"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soner3/evently/db"
	"github.com/soner3/evently/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {
	db.InitDB()

	endpoint := "localhost:50051"

	go runGrpcServer(endpoint)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := routes.InitRestRoutes(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatalln("Could initialize REST endpoint:", err)
	}

	log.Println("Started REST server")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalln("Could not start web server:", err)
	}

}

func runGrpcServer(endpoint string) {
	lis, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalln("Could not listen to port:", err)
	}

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalln("Could not create interceptor:", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(protovalidate_middleware.UnaryServerInterceptor(validator)))
	routes.InitGrpcRoutes(grpcServer)

	// Only in dev mode
	reflection.Register(grpcServer)

	log.Println("Started gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Could not start gRPC server:", err)
	}
}
