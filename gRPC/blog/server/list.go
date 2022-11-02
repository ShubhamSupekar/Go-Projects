package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpcapp.com/blog/proto"
)

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("ListBlog function running")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Unknown internal error",
		)
	}
	defer cur.Close(stream.Context())

	for cur.Next(stream.Context()) {
		data := &BlogIteam{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				"Error while decoding data from MongoDB",
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			"Unknown internal error",
		)
	}

	return nil
}
