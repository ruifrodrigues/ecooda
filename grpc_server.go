package main

import (
	"github.com/ruifrodrigues/ecooda/services/challenge"
	challengev1 "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

func runGrpcServer() {
	// create new gRPC server
	server := grpc.NewServer()

	// register the GreeterServerImpl on the gRPC server
	challengev1.RegisterChallengeServiceServer(server, challenge.NewChallengeService())

	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", ":8080"); err != nil {
		log.Fatal("error in listening on port :8080", err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
