package main

/*
 #cgo CFLAGS: -I.
 #cgo LDFLAGS: -L. -lparser
 #cgo LDFLAGS: -L../gparser/lib -lgspcore
 #cgo LDFLAGS: -L../gparser/lib -lgspcollection
 #include <parser.h>
 */
import "C"

import "fmt"

func main() {
	sqlText := `CREATE TABLE test_set AS
SELECT al.object_name, u.name
FROM all_objects al, users u
WHERE u.id = al.user
AND SUBSTR(al.object_name,1,1) BETWEEN 'A' AND 'N';`

	fmt.Println("<------------beg go C.parserData() -------------")
	data := C.parserData(C.CString(sqlText))
	fmt.Println("------------end go C.parserData() -------------> \n")
	//fmt.Printf("table format = %v\n",table1.tables)
		
	//fmt.Printf("data = %v\n",data)
	for i:=0; i<int(data.size); i++ {

		fmt.Printf("data.tables[%d] = %v\n",i,C.GoString(data.tables[i].tableName))
		for j:=0;j<int(data.tables[i].size);j++{
			fmt.Printf("data.tables[%d].columns[%d] = %v\n",i,j,C.GoString(data.tables[i].columns[j]))
		} 
	}

}
