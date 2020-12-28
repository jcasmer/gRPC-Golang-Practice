package main

import (
	"context"
	"fmt"
	"grpc-go-course/greet/greetpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v ", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client \n")
	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC..")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Juan",
			LastName:  "Mera",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greeet RPC: %v ", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC..")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Juan",
			LastName:  "Mera",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v ", err)
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
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Client Streaming RPC..")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Juan",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Esteban",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Felipe",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Jhon",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error calling LongGreet: %v", err)
	}
	// we iterate over our slice and send each message individally
	for _, req := range requests {
		fmt.Printf("Sending req: %v \n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from  LongGreet: %v", err)
	}
	fmt.Printf("LongGreet response: %v \n ", res)
}
