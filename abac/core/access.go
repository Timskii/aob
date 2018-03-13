package core

import (
	"log"

	//cb "github.com/casbin/casbin"
	pb "ps.kz/aob/abac/protos"
)

func checkAccess(userSql *pb.UserSql) {

	//e := cb.NewEnforcer("")
	for _, tableColumn := range userSql.Obj.TableColumns {
		log.Printf("/n tableColumn=%v", tableColumn)
	}
}
