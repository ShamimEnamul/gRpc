package main

import (
	"context"
	pb "github.com/ShamimEnamul/grpc/calculator/proto"
	"io"
	"log"
)

func Primes(c pb.CalculatorServiceClient) {
	reqBody := &pb.PrimeRequest{
		Val1: 120,
	}

	stream, err := c.Primes(context.Background(), reqBody)
	if err != nil {
		log.Fatal("Error occurred from primes")
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error occurred from primes while reading")
		}
		log.Println("result: ", msg.Result)
	}
	return
}
