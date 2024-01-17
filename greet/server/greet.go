package main

import (
	"context"
	pb "github.com/ShamimEnamul/grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function is invoked with %v\n", req)
	return &pb.GreetResponse{
		Result: "Response for" + req.FirstName}, nil
}
