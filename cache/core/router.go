package core

import (
	"log"
	

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "ps.kz/aob/cache/protos"
)

const (
	address = "localhost:4001"
	
)

func cache(*pb.UserSql) {
	log.Println("cache")
}