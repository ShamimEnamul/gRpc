package main

import pb "github.com/ShamimEnamul/grpc/calculator/proto"

func (s *Server) Primes(req *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {

	number := req.Val1

	return nil
}
