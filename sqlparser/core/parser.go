package main

// #cgo CFLAGS: -I.../gparser/core 
// #cgo LDFLAGS: -L. -lgsp_sqlparser
// #include <gsp_sqlparser.h>
import "C"

import "fmt"

var vendor C.gsp_dbvendor = C.dbvoracle
var parser *C.gsp_sqlparser

func main() {
	rc := C.gsp_parser_create(vendor,&parser)
	fmt.println(rc)
		
}
