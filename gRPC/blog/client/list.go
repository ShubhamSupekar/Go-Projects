package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpcapp.com/blog/proto"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("listBlog func running")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while calling listBlogs: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Something happened: %v\n", err)
		}

		log.Println(res)
	}
}
