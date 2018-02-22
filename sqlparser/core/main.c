#include <stdio.h>
#include <stdlib.h>
#include "parser.h"

int main()
{

	Data data = parserData("select 1 from dual");
	int size = sizeof(data.tables)/sizeof(Table*);

	
	printf("size = %d\n",size);

	for (int i=0; i < data.size; i++){
		printf("data.tables[%d] = %s\n",i,data.tables[i]->tableName);
			
			for (int j=0; j<data.tables[i]->size; j++){
				printf("data.tables[%d]->columns[%d] = %s\n",i,j,data.tables[i]->columns[j]);
			}
	}
	return 0;
}