
package main

import (
	"log"
	

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ps.kz/aob/sqlrelay/protos"
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
	
	r, err := c.ParseAndCheckAccess(context.Background(), &pb.UserSql{User:"tim",Script:"sElect t.data,t.username, t.place, ad.name, ad.house fRom /*people p*/ test t, adresses ad"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %+v", r)
}