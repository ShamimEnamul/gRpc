package main

import (
	"context"
	"fmt"
	pb "github.com/ShamimEnamul/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var address string = "0.0.0.0:5521"

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "shamim enamul",
	})
	if err != nil {
		log.Fatalf("Error occurred from doGreet function %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)

}

func calculate(c pb.CalculatorServiceClient) {
	log.Println("calculate was invoked")
	res, err := c.Add(context.Background(), &pb.CalculateRequest{
		Val1: 2,
		Val2: 5,
	})

	if err != nil {
		{
			log.Fatalf("Error occurred from calculate function %v\n", err)
		}
	}

	log.Println(res)
	log.Println("Sum: ", res.Sum)

}

func main() {
	fmt.Println("Client...")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect! ", err)
	}
	defer conn.Close()

	//greetClient := pb.NewGreetServiceClient(conn)
	//doGreet(greetClient)
	calculatorClient := pb.NewCalculatorServiceClient(conn)
	calculate(calculatorClient)
}
