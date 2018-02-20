go run -compiler=gccgo -gccgoflags="-L lib -lgspcore -lgspcollection" -gcflags="-I ext/collection/includes -I ext/expr_traverse -I ext/node_visitor -I core -I ext/modifysql" parser.go
