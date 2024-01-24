package main

import (
	pb "github.com/ShamimEnamul/grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var address string = "0.0.0.0:5521"

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect! ", err)
	}
	defer conn.Close()

	//greetClient := pb.NewGreetServiceClient(conn)
	//doGreet(greetClient)
	calculatorClient := pb.NewCalculatorServiceClient(conn)
	Calculate(calculatorClient)
}
