package main

import (
	"context"
	pb "github.com/ShamimEnamul/grpc/calculator/proto"
	"log"
)

func Calculate(c pb.CalculatorServiceClient) {
	log.Println("calculate was invoked")
	res, err := c.Add(context.Background(), &pb.SumRequest{
		Val1: 2,
		Val2: 5,
	})

	if err != nil {
		{
			log.Fatalf("Error occurred from calculate function %v\n", err)
		}
	}

	log.Println(res)

}
