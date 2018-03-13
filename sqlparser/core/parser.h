#include <stdlib.h>
#define MAX_SIZE_TAB 128
#define MAX_SIZE_COL 1024

/*typedef struct Table{
	char *tableName;
	char *columns[MAX_SIZE_COL];
	int size;
	
}Table;*/

typedef struct Data{
	char *tableColumns[MAX_SIZE_TAB];
	int size;
}Data;

Data parserData(char *sqlScript);