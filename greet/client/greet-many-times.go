package main

//func DoGreetManyTimes(c pb.GreetServiceClient) {
//	log.Println("DoGreetManyTimes was invoked!")
//	req := &pb.GreetRequest{FirstName: "Shamim"}
//
//	stream, err := c.GreetManyTimes(context.Background(), req)
//
//	if err != nil {
//		log.Fatal("Error while calling GreetManyTimes %v \n", err)
//	}
//
//	for {
//		msg, err := stream.Recv()
//		if err == io.EOF {
//			break
//		}
//
//		if err != nil {
//			log.Fatal("Error while getting data from GreetManyTimes %v \n", err)
//		}
//		log.Printf("GreetingManyTimes %s \n", msg.Result)
//	}

//}
