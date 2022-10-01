package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	userService "github.com/onemgvv/grpc-contract/gen/go/users/service/v1"
	v1 "github.com/onemgvv/grpc-service/internal/delivery/grpc/v1"
	"github.com/onemgvv/grpc-service/pkg/database"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

const grpcHostPort = "0.0.0.0:5001"

func main() {
	grpcServer := grpc.NewServer()
	storage := database.NewStorage()
	listen, err := net.Listen("tcp", grpcHostPort)
	if err != nil {
		log.Fatal(err.Error())
	}

	userService.RegisterUserServiceServer(grpcServer, v1.NewUserServer(
		userService.UnimplementedUserServiceServer{},
		storage,
	))

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err = userService.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, grpcHostPort, opts)
	if err != nil {
		log.Fatal(err.Error())
	}

	g, _ := errgroup.WithContext(context.Background())
	g.Go(func() (err error) {
		return grpcServer.Serve(listen)
	})

	g.Go(func() (err error) {
		return http.ListenAndServe(":5000", mux)
	})

	log.Println("App started!")

	if err = g.Wait(); err != nil {
		log.Fatal(err.Error())
	}
}
