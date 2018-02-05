package core

import (
		"log"

		"golang.org/x/net/context"
		pb "ps.kz/aob/main/protos"
		)

func NewAdminServer() *ServerAdmin {
	s := new(ServerAdmin)
	return s
}

type ServerAdmin struct{

}

func (*ServerAdmin) ParseAndCheckAccess (ctx context.Context, userSql *pb.UserSql) (*pb.UserSql, error){
	log.Println(userSql)
	userSql.Script = "select * from test"
	parser(userSql)
	cache(userSql)
	return userSql,nil
}

func (*ServerAdmin) ParseSql (ctx context.Context, script *pb.Script) (*pb.ObjScript, error){
	return nil, nil
}
func (*ServerAdmin) CheckAccess(ctx context.Context, userSql *pb.UserSql) (*pb.UserSql, error) {
	return nil, nil
}