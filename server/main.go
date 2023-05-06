package server

import (
	"github.com/ruifrodrigues/ecooda/challenge"
	"google.golang.org/grpc"
	"log"
	"net"

	gw "github.com/ruifrodrigues/ecooda/stubs/go/challenge/v1"
)

func main() {
	// create new gRPC server
	server := grpc.NewServer()

	// register the GreeterServerImpl on the gRPC server
	gw.RegisterChallengeServiceServer(server, challenge.NewChallengeService())

	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", ":5051"); err != nil {
		log.Fatal("error in listening on port :5051", err)
	} else {
		// the gRPC server
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}
