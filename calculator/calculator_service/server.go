package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	result := req.GetFirstNum() + req.GetSecondNum()
	res := &calculatorpb.SumResponse{
		Result: result,
	}
	return res, nil
}

func (*server) PrimeNumberManyTimes(req *calculatorpb.PrimeNumberManyTimesRequest, stream calculatorpb.CalculatorService_PrimeNumberManyTimesServer) error {
	fmt.Printf("PrimeNumberManyTimes function was invoked with %v\n", req)
	N := req.GetNumber()
	k := int32(2)
	for N > 1 {
		if (N % k) == 0 { // if k evenly divides into N
			res := &calculatorpb.PrimeNumbertManyTimesResponse{
				Result: k,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			N = N / k // divide N by k so that we have the rest of the number left.
		} else {
			k++
		}

	}

	return nil

}

func (*server) ComputeAvg(stream calculatorpb.CalculatorService_ComputeAvgServer) error {
	fmt.Printf("ComputeAvg function was invoked with a stream request \n")
	sum := 0.0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			average := sum / float64(count)
			// we have finished reading the client stream
			return stream.SendAndClose(&calculatorpb.ComputeAvgResponse{
				Result: average,
			})
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
		}
		sum += float64(req.GetNumber())
		count++
	}
}

func main() {
	fmt.Println("calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v ", err)
	}
}
