package server

import (
	"awesomeProject/task1/ayushpb/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Start(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	ayushpb.RegisterCalculatorServiceServer(s, &CalculatorHandler{})

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
