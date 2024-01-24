package main

import (
	"fmt"
	pb "github.com/ShamimEnamul/grpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function is invoked with %v\n")

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello, number %d", req.FirstName, i)

		err := stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
