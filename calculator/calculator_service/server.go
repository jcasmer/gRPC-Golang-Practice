package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

func (*server) FindMaximun(stream calculatorpb.CalculatorService_FindMaximunServer) error {
	fmt.Printf("FindMaximun function was invoked with a stream request \n")
	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error while reading client stream: %v", err)
			return err
		}
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			sendErr := stream.Send(&calculatorpb.FindMaximunResponse{
				Maximum: maximum,
			})
			if sendErr != nil {
				log.Fatalf("error while sending data to client: %v", sendErr)
				return sendErr
			}

		}
	}

}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Printf("Received SquareRoot RPC \n")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	fmt.Println("=========== number ", number)
	return &calculatorpb.SquareRootResponse{
		DoubleRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("calculator server")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	// resgister reflection services on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v ", err)
	}
}
