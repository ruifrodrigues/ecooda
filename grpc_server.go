package main

import (
	"github.com/ruifrodrigues/ecooda/config"
	"github.com/ruifrodrigues/ecooda/services/challenge"
	"github.com/ruifrodrigues/ecooda/services/location"
	ecoodav1 "github.com/ruifrodrigues/ecooda/stubs/go/ecooda/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func runGrpcServer(conf config.Config) {
	// create new gRPC server
	server := grpc.NewServer()

	// register services from gRPC server
	ecoodav1.RegisterChallengeServiceServer(server, challenge.NewChallengeService(conf))
	ecoodav1.RegisterLocationServiceServer(server, location.NewLocationService(conf))

	// start listening for a tcp connection
	if l, err := net.Listen("tcp", conf.Values.GrpcPort); err != nil {
		log.Fatal("error in listening on port "+conf.Values.GrpcPort, err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
