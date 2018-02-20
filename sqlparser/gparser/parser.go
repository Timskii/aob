package main

/*
#cgo CFLAGS: -c -I ext/collection/includes -I ext/expr_traverse -I ext/node_visitor -I core -I ext/modifysql
#cgo LDFLAGS: -L lib -lgspcore -lgspcollection
#include <gsp_sqlparser.h>
#include <gsp_base.h>
#include <node_visitor.h>
*/
import "C"

import "fmt"
import "unsafe"

type Table struct{
	Name string
	Columns []string
}


/*func process_table(node *C.gsp_node, traverser *SqlTraverser) {
	var tables []Table

	str := C.gsp_node_text(C._getTableName(node));
}*/

func main() {

	var vendor C.gsp_dbvendor = C.dbvoracle
	var parser *C.gsp_sqlparser
	var sqlScript string = "select 2 from dual"

	rc := C.gsp_parser_create(vendor,&parser)
	defer C.gsp_parser_free(parser)
	
	fmt.Println(rc)

	
	script := C.CString(sqlScript)
	defer C.free(unsafe.Pointer(script))
	fmt.Println(sqlScript)
	fmt.Printf("rc = |%-8T|\n",rc)
	fmt.Printf("sqlScript = |%-8T|\n",sqlScript)

	rc= C.gsp_check_syntax(parser, script)
	trave := C.createSqlTraverser()
	fmt.Println(rc)

	/*for index:0;index<gsp_parser_freeer.nStatement;index++ {
		nodeList := traverser.traverseSQL(traverser, &parser.pStatement[index])
		iter := nodeList.getIterator(nodeList)
		/*while(nodeList.hasNext(nodeList,&iter))
		{
			node := *C.gsp_node.nodeList.next(&iter)
			if(node.nodeList.next(&iter).nodeType == C.t_gsp_table ){
				fmt.Println("process_table t_gsp_table")
				//process_table(node,traverser);
			}else if(node.nodeType == C.t_gsp_fromTable ){
				fmt.Println("process_table t_gsp_fromTable")
				//process_table(node,traverser);
			}
		}
	}*/
	
	
		
}
