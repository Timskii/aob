#include "node_visitor.h"
#include "gsp_base.h"
#include "gsp_node.h"
#include "gsp_list.h"
#include "gsp_sourcetoken.h"
#include "gsp_sqlparser.h"
#include "parser.h"
#include <stdlib.h>
#include <stddef.h>
#include <string.h>
#include <stdio.h>

Data data;
int sizeTableColumns;

static gsp_objectname* _getTableName(gsp_node* table){
	if(table->nodeType == t_gsp_table){
		gsp_table *normalTable = (gsp_table*)table;
		return normalTable->tableName;
	}
	else if(table->nodeType == t_gsp_fromTable){
		gsp_fromTable *fromTable = (gsp_fromTable*)table;
		return fromTable->tableName;
	}
	return NULL;
}

static void _process_table(gsp_node *node, struct gsp_visitor *visitor){
	
	char *str;
	SqlTraverser *traverser = (SqlTraverser *)visitor->context;
	if(traverser->isBaseTable(traverser, (gsp_node*)node)){
		str = gsp_node_text((gsp_node *)_getTableName(node));
		
		{
			List *fields = (List *)traverser->getTableObjectNameReferences(traverser, node);
			if(fields!=NULL){
				Iterator fieldIter = fields->getIterator(fields);
				while(fields->hasNext(fields, &fieldIter)){
					gsp_objectname *field = (gsp_objectname *)fields->next(&fieldIter);
					char* fieldName = (char *)malloc((strlen(str) + field->partToken->nStrLen + 2 + 24)*sizeof(char));
					memset(fieldName,'\0', (strlen(str) + field->partToken->nStrLen + 2 + 24)*sizeof(char));
					strcat(fieldName,str);
					strcat(fieldName,".");
					strncat(fieldName,field->partToken->pStr, field->partToken->nStrLen);
					data.tableColumns[sizeTableColumns] = (char*)malloc(sizeof(char));
					data.tableColumns[sizeTableColumns] = fieldName;
					printf("tableColumns = %s\n",data.tableColumns[sizeTableColumns]);
					sizeTableColumns++;
				}
				
			}
		}
	}
}

static struct gsp_visitor* _createVisitor(SqlTraverser *traverser){
	struct gsp_visitor *visitor = (struct gsp_visitor *)malloc(sizeof(struct gsp_visitor)+t_gsp_exceptionClause*sizeof(&visitor->handle_node));
	visitor->context = traverser;
	visitor->handle_node[t_gsp_table] = _process_table;
	visitor->handle_node[t_gsp_fromTable] = _process_table;
	return visitor;
}

static void _disposeVisitor(struct gsp_visitor* visitor){
	free(visitor);
}

Data parserData(char *sqlText)
{
	int rc, index;
	gsp_sqlparser *parser;
	List *nodeList;
	Iterator iter;
	struct gsp_visitor *visitor;
	SqlTraverser* traverser;
	

	//char *sqlText = "select m.revision, m.esttim, t.data from mailstore m join testtable t on t.id = m.id";

	gsp_dbvendor vendor = dbvoracle;

	rc = gsp_parser_create(vendor,&parser);
	if (rc){
		printf("create parser error");
		//todo return
	}

	printf("Origin SQL:\n");
	printf("%s\n",sqlText);
	printf("\n\n");
	rc= gsp_check_syntax(parser, sqlText);
	if (rc != 0){
		printf("parser error:%s\n",gsp_errmsg(parser));
		gsp_parser_free(parser);
		//todo return
	}
	
	traverser = createSqlTraverser();
	visitor = _createVisitor(traverser);

	for(index=0;index<parser->nStatement;index++ ){
		nodeList = traverser->traverseSQL(traverser, &parser->pStatement[index]);
		iter = nodeList->getIterator(nodeList);
		while(nodeList->hasNext(nodeList,&iter))
		{
			gsp_node *node = (gsp_node*)nodeList->next(&iter);
			if(node->nodeType == t_gsp_table ){
				visitor->handle_node[t_gsp_table](node, visitor);
			}else if(node->nodeType == t_gsp_fromTable ){
				visitor->handle_node[t_gsp_fromTable](node, visitor);
			}
		}
	}
	printf("sizeTableColumns=%d\n",sizeTableColumns);
	data.size = sizeTableColumns;
	sizeTableColumns=0;
	traverser->dispose(traverser);
	_disposeVisitor(visitor);
	gsp_parser_free(parser);
	return data;
}
