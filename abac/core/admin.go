package core

import (
		"log"

		"golang.org/x/net/context"
		pb "ps.kz/aob/abac/protos"
		)

func NewAdminServer() *ServerAdmin {
	s := new(ServerAdmin)
	return s
}

type ServerAdmin struct{

}

func (*ServerAdmin) CheckAccess (ctx context.Context, userSql *pb.UserSql) (*pb.UserSql, error){
	log.Println("abac")
	return userSql,nil
}
