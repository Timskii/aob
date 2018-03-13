package core

import (
	"log"

	cb "github.com/casbin/casbin"
	pb "ps.kz/aob/abac/protos"
)

func enforce(e *cb.Enforcer, sub string, obj string, act string) {
	e.Enforce(sub, obj, act)
}

func checkAccess(userSql *pb.UserSql) {

	e := cb.NewEnforcer("/opt/gopath/src/ps.kz/aob/abac/core/abac_model_policy.conf", "/opt/gopath/src/ps.kz/aob/abac/core/policy.csv")
	for _, tableColumn := range userSql.Obj.TableColumns {
		log.Printf("\ntableColumn=%v", tableColumn)

		enforce(e, userSql.GetUser(), tableColumn, "read")
	}
}
