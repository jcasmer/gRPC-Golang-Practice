package main

import (
	"context"
	"fmt"
	"grpc-go-course/calculator/calculatorpb"
	"io"
	"log"

	"google.golang.org/grpc"
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
	doServerStreaming(c)
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
	log.Printf("Response from Greet: %v", res.Result)
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
