package main

import (
	"context"
	"log"

	pb "grpcapp.com/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog func running")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("eror while deleting: %v\n", err)
	}

	log.Println("Blog was deleted.")
}
