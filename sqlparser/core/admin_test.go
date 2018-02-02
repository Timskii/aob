package core

import "testing"
import "fmt"
import pb "ps.kz/aob/protos"
import "golang.org/x/net/context"

func TestParseSql(t *testing.T) {
	
	script := &pb.Script{"select sysdate from dual"}

	s := NewAdminServer()

	objsql,_ := s.ParseSql(context.Background(),script)
	fmt.Println(objsql.ObjScript)
}