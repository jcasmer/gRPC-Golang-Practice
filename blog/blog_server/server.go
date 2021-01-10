package main

import (
	"context"
	"fmt"
	"grpc-go-course/blog/blogpb"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

type server struct{}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	fmt.Println("Creating blog")
	blog := req.GetBlog()

	data := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}
	res, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}
	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func main() {
	// if we crash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// mongodb server
	fmt.Println("Connecting to MongoDB")
	client, dbErr := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if dbErr != nil {
		log.Fatalf("Failed to listen: %v ", dbErr)
	}
	dbErr = client.Connect(context.TODO())
	if dbErr != nil {
		log.Fatalf("Failed to listen: %v ", dbErr)
	}
	collection = client.Database("mydb").Collection("blog")

	// blog service
	fmt.Println("Blog Service Started")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v ", err)
	}
	certFile := "ssl/server.crt"
	keyFile := "ssl/server.pem"
	creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
	if sslErr != nil {
		log.Fatalf("Failed loading certificates: %v ", sslErr)
		return
	}
	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {

		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v ", err)
		}

	}()

	// Wait for control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stopping the listener")
	lis.Close()
	fmt.Println("Closing MongoDB Connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of program")

}
