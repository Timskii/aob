#!/bin/sh

f_o="parser.o"
f_a="libparser.a"
LIB_DIR="../gparser/lib"
BUILD_DIR="."
DEMO_SRC_DIR="."
CORE_DIR="../gparser/core"
COLLECTION_DIR="../gparser/ext/collection"
EXPRTRAVERSE_DIR="../gparser/ext/expr_traverse"
NODEVISITOR_DIR="../gparser/ext/node_visitor"
CORE_LIB="gspcore"
COLLECTION_LIB="gspcollection"

echo "-- check binary $f_o and $f_a"
if [ -f "$f_o" ]; then
	rm $f_o
	echo "--file $f_o deleted! --"
fi
if [ -f "$f_a" ]; then
	rm $f_a
	echo "--file $f_a deleted! --"
fi
gcc parser.c "$NODEVISITOR_DIR"/node_visitor.c -static -Wall -c -I "$COLLECTION_DIR)"/includes -I "$EXPRTRAVERSE_DIR" -I "$CORE_DIR" -I "$NODEVISITOR_DIR" -L"$LIB_DIR" -l"$CORE_LIB" -l"$COLLECTION_LIB" -lgspcollection64
ar rcs "$f_a" "$f_o"

