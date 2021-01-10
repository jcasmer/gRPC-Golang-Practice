package main

import (
	"context"
	"fmt"
	"grpc-go-course/blog/blogpb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	fmt.Println("hello I'm a client")

	certFile := "ssl/ca.crt" // certifcate Authority trust certificate
	creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
	if sslErr != nil {
		log.Fatalf("Failed while loading CA trust certificate: %v ", sslErr)
		return
	}
	opts := grpc.WithTransportCredentials(creds)
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("Could not connect: %v ", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	fmt.Printf("Created client \n")

	// create blog
	blog := &blogpb.Blog{
		AuthorId: "juan",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	createBlog, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Could not create blog %v ", err)
	}
	fmt.Printf("Blog has been created: %v", createBlog)
}
