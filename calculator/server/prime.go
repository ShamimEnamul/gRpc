package main

import pb "github.com/ShamimEnamul/grpc/calculator/proto"

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {

	number := req.Val1
	k := 2
	for number > 1 {
		if int(number)%k == 0 {
			err := stream.Send(&pb.PrimeResponse{
				Result: int32(k),
			})
			if err != nil {
				return err
			}
			number = int32(int(number) / k)
		} else {
			k += 1
		}
	}
	return nil
}
