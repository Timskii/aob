package core

import (
	"log"
	

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ps.kz/aob/main/protos"
)

const (
	addressP = "localhost:4001"
	addressCA = "localhost:4002"
	
)

func parser(userSql *pb.UserSql){
	conn, err := grpc.Dial(addressP, grpc.WithInsecure())
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
	userSql.Script = r.ObjScript
	log.Printf("Greeting: %s", userSql.Script)
}

func cache(userSql *pb.UserSql) {
	conn, err := grpc.Dial(addressCA, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAOBServiceClient(conn)

	userSql, err = c.CheckAccess(context.Background(), userSql)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", userSql)
}