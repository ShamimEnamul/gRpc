package main

import (
	"context"
	pb "github.com/ShamimEnamul/grpc/calculator/proto"
	"log"
)

func (s *Server) Add(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Calculate function is invoked with %v\n", req)
	return &pb.SumResponse{
		Result: req.Val1 + req.Val2}, nil
}
