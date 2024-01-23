package main

import (
	"context"
	"fmt"
	pb "github.com/ShamimEnamul/grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
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
func DoGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("DoGreetManyTimes was invoked!")
	req := &pb.GreetRequest{FirstName: "Shamim"}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatal("Error while calling GreetManyTimes %v \n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error while getting data from GreetManyTimes %v \n", err)
		}
		log.Printf("GreetingManyTimes %s \n", msg.Result)
	}
}

func main() {
	fmt.Println("Client...")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect! ", err)
	}
	defer conn.Close()

	greetClient := pb.NewGreetServiceClient(conn)
	DoGreetManyTimes(greetClient)

}
