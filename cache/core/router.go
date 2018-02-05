package core

import (
	"log"
	

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ps.kz/aob/cache/protos"
)

const (
	address = "localhost:4003"
	
)

func cache(userSql *pb.UserSql) {
	log.Println("cache")
	conn, err := grpc.Dial(addressCA, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAOBServiceClient(conn)

	r, err := c.CheckAccess(context.Background(), userSql)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}