package main

import (
	"context"
	"fmt"
	"grpc-go-course/blog/blogpb"
	"io"
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

	// doCreateblog(c)
	// doReadblog(c)
	// doUpdateblog(c)
	// doDeleteblog(c)
	doListblog(c)

}

func doCreateblog(c blogpb.BlogServiceClient) {
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

func doReadblog(c blogpb.BlogServiceClient) {

	// just for tutorial, this won't read a blog
	_, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5ffb93b80562f1b6ea9da2cc"})
	if err != nil {
		fmt.Printf("Error happened while reading: %v \n", err)
	}

	readBlog, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5ffb93b80562f1b6ea9da2cf"})
	if err != nil {
		fmt.Printf("Error happened while reading: %v", err)
	}
	fmt.Printf("Blog was read: %v", readBlog)
}

func doUpdateblog(c blogpb.BlogServiceClient) {

	// create blog
	blog := &blogpb.Blog{
		Id:       "5ffb93b80562f1b6ea9da2cf",
		AuthorId: "juan esteban",
		Title:    "My first blog (edited",
		Content:  "Content of the first blog, with addtions",
	}

	updatedBlog, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog})
	if err != nil {
		fmt.Printf("Error happened while updating: %v", err)
		return
	}
	fmt.Printf("Blog was updated: %v", updatedBlog)
}

func doDeleteblog(c blogpb.BlogServiceClient) {

	deleteBlog, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: "5ffb8b9c3e8b888b06cd6626"})
	if err != nil {
		fmt.Printf("Error happened while deleting: %v", err)
		return
	}
	fmt.Printf("Blog was deleted: %v", deleteBlog)
}

func doListblog(c blogpb.BlogServiceClient) {

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		fmt.Printf("Error happened while calling ListBlog: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		fmt.Printf("Blog: %v \n", res.GetBlog())
	}

}
