package main

import (
	"context"
	"go_bibit_test/movie/delivery/grpc/movie_grpc"
	"log"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect")
	}
	defer conn.Close()

	c := movie_grpc.NewMovieHandlerClient(conn)
	message := movie_grpc.SingleRequest{
		Id: "tt0372784",
	}

	response, err := c.GetMovie(context.Background(), &message)
	if err != nil {
		log.Fatalf(err.Error())
	}

	log.Println(response)
}
