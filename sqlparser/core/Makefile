LIB_DIR =  ../gparser/lib
BUILD_DIR = .
DEMO_SRC_DIR = .
CORE_DIR = ../gparser/core
COLLECTION_DIR = ../gparser/ext/collection
EXPRTRAVERSE_DIR = ../gparser/ext/expr_traverse
NODEVISITOR_DIR = ../gparser/ext/node_visitor
GSP_CORE_LIB = gspcore
GSP_COLLECTION_LIB = gspcollection
LIB_NAME = libparser.a

AR = ar
ARFLAGS = rcs
CC = gcc
CFLAGS = -c -I $(COLLECTION_DIR)/includes -I $(EXPRTRAVERSE_DIR) -I $(CORE_DIR) -I $(NODEVISITOR_DIR)


#DEMO_EXE = $(BUILD_DIR)/parser

SOURCE = $(DEMO_SRC_DIR)/parser.c \
	$(EXPRTRAVERSE_DIR)/expr_traverse.c \
	$(NODEVISITOR_DIR)/node_visitor.c 
	

OBJECTS = $(BUILD_DIR)/parser.o \
	$(BUILD_DIR)/expr_traverse.o \
	$(BUILD_DIR)/node_visitor.o


parser: create_build_dir \
	$(OBJECTS) \
	create_parser \
	clean_obj
	
create_parser:
	$(AR)  $(ARFLAGS)  $(LIB_NAME) $(OBJECTS)
#	$(CC)  $(OBJECTS)  -L$(LIB_DIR) -l$(GSP_CORE_LIB) -l$(GSP_COLLECTION_LIB) -o $(DEMO_EXE)
	
create_build_dir:
	@if [ ! -d "$(BUILD_DIR)" ]; \
	then \
	    mkdir $(BUILD_DIR); \
	fi \
	 	
$(BUILD_DIR)/parser.o: $(DEMO_SRC_DIR)/parser.c
	$(CC) $(CFLAGS) $< -o $@
	
$(BUILD_DIR)/expr_traverse.o: $(EXPRTRAVERSE_DIR)/expr_traverse.c
	$(CC) $(CFLAGS) $< -o $@

$(BUILD_DIR)/node_visitor.o: $(NODEVISITOR_DIR)/node_visitor.c
	$(CC) $(CFLAGS) $< -o $@

clean_obj:
	rm -f $(BUILD_DIR)/*.o 
	
clean:
	rm -rf $(BUILD_DIR)
