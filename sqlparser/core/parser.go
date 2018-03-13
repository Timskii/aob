package core

/*
 #cgo CFLAGS: -I. -I../gparser/core -I../gparser/ext/collection/includes -I../gparser/ext/expr_traverse -I../gparser/ext/node_visitor
 #include <parser.h>
*/
import "C"

import (
	"fmt"
	pb "ps.kz/aob/sqlparser/protos"
)

func parserData(sqlText string) (*pb.ObjScript, error) {

	/*sqlText = `CREATE TABLE test_set AS
	SELECT al.object_name, u.name
	FROM all_objects al, users u
	WHERE u.id = al.user
	AND SUBSTR(al.object_name,1,1) BETWEEN 'A' AND 'N';`*/
	var objScripts *pb.ObjScript = &pb.ObjScript{nil}
	fmt.Println("<------------beg go C.parserData() -------------")
	data := C.parserData(C.CString(sqlText))
	fmt.Println("------------end go C.parserData() -------------> \n")

	fmt.Printf("int(data.size) = %+v\n", int(data.size))

	for i := 0; i < int(data.size); i++ {
		fmt.Printf("\ndata.tables[%d] = %v\n", i, C.GoString(data.tableColumns[i]))
		objScripts.TableColumns = append(objScripts.TableColumns, C.GoString(data.tableColumns[i]))
	}
	return objScripts, nil
}
