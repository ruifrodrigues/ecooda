package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ruifrodrigues/ecooda/config"
	pb "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func runHttpProxy(conf config.Config) {
	ctx := context.Background()

	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux()

	endpoint := "localhost" + conf.Values.GrpcPort
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// setting up a dail up for gRPC service by specifying endpoint/target url
	if err := pb.RegisterChallengeServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatal(err)
	}

	if err := pb.RegisterLocationServiceHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		log.Fatal(err)
	}

	// Creating a normal HTTP server
	server := http.Server{
		Handler: mux,
	}

	// creating a listener for server
	l, err := net.Listen("tcp", conf.Values.HttpPort)
	if err != nil {
		log.Fatal(err)
	}

	// start server
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}

	// close DB connection
	_ = conf.Database.CloseConnection()
}
