
package main

import (
	"log"
	

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ps.kz/aob/protos"
)

const (
	address = "localhost:4000"
	
)

type User struct {
	name string
	age int32
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAOBServiceClient(conn)
	
	r, err := c.ParseAndCheckAccess(context.Background(), &pb.UserSql{User:"tim",Script:"select * frim dual"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}