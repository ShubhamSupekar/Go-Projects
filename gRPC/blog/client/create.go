package main

import (
	"context"
	"log"

	pb "grpcapp.com/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog function running")

	blog := &pb.Blog{
		AuthorId: "Shubham",
		Title: "My First Blog",
		Content: "This is the content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(),blog)
	if err !=nil{
		log.Fatalf("Unexpexted err: %v\n",err)
	}

	log.Printf("Blog has been created: %s",res.Id)
	return res.Id
}
