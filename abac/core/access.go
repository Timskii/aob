package core

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ps.kz/aob/protos"
)

const (
	address = "localhost:4001"
)

func parser(userSql *pb.UserSql) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAOBServiceClient(conn)

	script := &pb.Script{userSql.Script}

	r, err := c.ParseSql(context.Background(), script)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}

func cache(*pb.UserSql) {
	log.Println("cache")
}
