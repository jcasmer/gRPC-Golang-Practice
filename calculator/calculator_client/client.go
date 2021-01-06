package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("hello I'm a calculator client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v ", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	fmt.Printf("Created calculator client \n")
	// doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBiDiStreaming(c)
	doErrorUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC..")
	req := &calculatorpb.SumRequest{
		FirstNum:  3,
		SecondNum: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("erro while calling Calculator RPC: %v ", err)
	}
	log.Printf("Response from Calculator: %v", res.Result)
}

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC..")

	req := &calculatorpb.PrimeNumberManyTimesRequest{
		Number: 120,
	}

	resStream, err := c.PrimeNumberManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling PrimeNumberManyTimes RPC: %v ", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from PrimeNumberManyTimes: %v", msg.GetResult())
	}

}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC..")

	numbers := []int32{1, 2, 3, 4}
	stream, err := c.ComputeAvg(context.Background())
	if err != nil {
		log.Fatalf("error calling ComputeAvg: %v", err)
	}
	// we iterate over our slice and send each message individally
	for _, req := range numbers {
		fmt.Printf("Sending req: %v \n", req)
		stream.Send(&calculatorpb.ComputeAvgRequest{
			Number: req,
		})
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from  ComputeAvg: %v", err)
	}
	fmt.Printf("ComputeAvg response: %v \n ", res)
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC..")
	stream, err := c.FindMaximun(context.Background())
	if err != nil {
		log.Fatalf("error while opening stream and calling FindMaximum: %v", err)
	}

	waitCh := make(chan struct{})
	// send go routine
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}
		for _, number := range numbers {
			fmt.Printf("Sending number: %v\n", number)
			stream.Send(&calculatorpb.FindMaximunRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	// receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				// we've reached the end of the stream
				break
			}
			if err != nil {
				log.Fatalf("error while reading server stream: %v", err)
				break
			}
			maximum := res.GetMaximum()
			fmt.Printf("Received a new maximum of ... %v \n", maximum)

		}
		close(waitCh)
	}()

	// block until everything is done
	<-waitCh
}

func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a SquareRoot Unary RPC..")

	// correct call
	doErrorCall(c, 10)

	// error call
	doErrorCall(c, -2)
}
func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	res, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{
		Number: n,
	})

	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			fmt.Println(respErr.Message(), respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative umber! ")
			}
		} else {
			log.Fatalf("Big error calling SquareRoot: %v ", err)
		}
	}
	fmt.Printf("Result of square root of %v:  %v \n", n, res.GetDoubleRoot())

}
