package main

import (
	"context"
	"log"

	pb "grpcapp.com/blog/proto"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog func running")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Soham",
		Title:    "A new title",
		Content:  "Updated content of the blog",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happen while updating")
	}

	log.Println("Blog was updated")
}
