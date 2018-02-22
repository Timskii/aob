#!/bin/sh

CORE_DIR="../gparser/core"
COLLECTION_DIR="../gparser/ext/collection"
EXPRTRAVERSE_DIR="../gparser/ext/expr_traverse"
NODEVISITOR_DIR="../gparser/ext/node_visitor"
GSP_CORE_LIB="gspcore"
GSP_COLLECTION_LIB="gspcollection"
LIB_DIR="../gparser/lib"
build="build"

file="build/main"
echo "check directory $build"
if [ ! -d "$build" ]; then
	mkdir "$build"
	echo "-- directory $build created --"
fi
echo "-- check binary $file"
if [ -f "$file" ]; then
	rm $file
	echo "--file $file deleted! --"
fi
gcc main.c parser.c -o "$file" -I "$COLLECTION_DIR)/includes" -I "$EXPRTRAVERSE_DIR" -I "$CORE_DIR" -I "$NODEVISITOR_DIR" -L"$LIB_DIR" -l"$GSP_CORE_LIB" -l"$GSP_COLLECTION_LIB"
./build/main