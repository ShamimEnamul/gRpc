package main

import (
	"context"
	"fmt"
	pb "github.com/ShamimEnamul/grpc/greet/proto"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"net"
)

const address string = "0.0.0.0:5521"

type Server struct {
	pb.GreetServiceServer
}

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function is invoked with %v\n", req)
	return &pb.GreetResponse{
		Result: "Response for" + req.FirstName}, nil
}

func main() {
	fmt.Println("SERVER...")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Error occurred!")
	}
	log.Println("Listening on", address)

	fmt.Println(lis.Addr())
	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed")
	}

}
