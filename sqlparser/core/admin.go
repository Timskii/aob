package core

import (
		"log"

		"golang.org/x/net/context"
		pb "ps.kz/aob/sqlparser/protos"
		)

func NewAdminServer() *ServerAdmin {
	s := new(ServerAdmin)
	return s
}

type ServerAdmin struct{

}

func (*ServerAdmin) ParseSql (ctx context.Context, script *pb.Script) (*pb.ObjScript, error){
	log.Println("parseSql",script)
	objScript := &pb.ObjScript{`{`+
		`"select":"*",`+
		`"from":"test1"`+
		`}`}
	
	return objScript, nil
}

