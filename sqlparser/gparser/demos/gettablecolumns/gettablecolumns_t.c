
#include "node_visitor.h"
#include "gsp_base.h"
#include "gsp_node.h"
#include "gsp_list.h"
#include "gsp_sourcetoken.h"
#include "gsp_sqlparser.h"
#include <stdlib.h>
#include <stddef.h>
#include <string.h>
#include <stdio.h>

static List* tableInfoList;
static List* fieldInfoList;

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

typedef struct Table{
	char *tableName;
	char *columns[5];
	
}Table;

static void _process_table(gsp_node *node, struct gsp_visitor *visitor){
	
	char *str;
	char *test;
	int index;
	SqlTraverser *traverser = (SqlTraverser *)visitor->context;
	if(traverser->isBaseTable(traverser, (gsp_node*)node)){
		str = gsp_node_text((gsp_node *)_getTableName(node));
		test = gsp_node_text(node);
		printf("\nstr = %s\n",str);
		printf("\ntest = %s\n",test);
		printf("\nt-gasp = %d\n",t_gsp_table);
		if(!tableInfoList->contains(tableInfoList, str)){
			tableInfoList->add(tableInfoList, str);
		}
		{
			List *fields = (List *)traverser->getTableObjectNameReferences(traverser, node);
			if(fields!=NULL){
				Iterator fieldIter = fields->getIterator(fields);
				while(fields->hasNext(fields, &fieldIter)){
					gsp_objectname *field = (gsp_objectname *)fields->next(&fieldIter);
					char* fieldName = (char *)malloc((strlen(str) + field->partToken->nStrLen + 2 + 24)*sizeof(char));
					memset(fieldName,'\0', (strlen(str) + field->partToken->nStrLen + 2 + 24)*sizeof(char));
					strcat(fieldName,str);
					printf("fieldname = %s\n",fieldName);
					strcat(fieldName,".");
					strncat(fieldName,field->partToken->pStr, field->partToken->nStrLen);
					printf("fieldname2 = %s\n",fieldName);
					strcat(fieldName,"(table determined:");
					strcat(fieldName,traverser->isTableDetermined(traverser, field)?"true":"false");
					strcat(fieldName,")");
					if(!fieldInfoList->contains(fieldInfoList, fieldName))
						fieldInfoList->add(fieldInfoList, fieldName);
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

static void _printErrorInfo(FILE *file, const char * format, ...){
	va_list argp;
	char* arg;

	va_start(argp, format);
	arg = va_arg(argp, char*);

	if(file==NULL){
		fprintf(stderr, format, arg);
	}
	else
		fprintf(file, format, arg);
	va_end(argp);
}

static void _printInfo(FILE *file, const char * format, ...){
	va_list argp;
	char* arg;

	va_start(argp, format);
	arg = va_arg(argp, char*);

	if(file==NULL){
		fprintf(stdout, format, arg);
	}
	else
		fprintf(file, format, arg);
	va_end(argp);
}

static void _printListInfo(List *list, FILE *infoResult ){
	Iterator infoIter = list->getIterator(list);
	while(fieldInfoList->hasNext(list,&infoIter))
	{
		char *str = (char *)list->next(&infoIter);
		_printInfo(infoResult,"%s\n", str);
		gsp_free(str);
	}
	_printInfo(infoResult,"\n");
}

int parserData(char *sqlText)
{
	int rc, argIndex, index;
	gsp_sqlparser *parser;
	List *nodeList, *argList;
	Iterator iter;
	struct gsp_visitor *visitor;
	SqlTraverser* traverser;
	FILE *infoResult= NULL;

	//char *sqlText = "select m.revision, m.esttim, t.data from mailstore m join testtable t on t.id = m.id";

	gsp_dbvendor vendor = dbvoracle;

	rc = gsp_parser_create(vendor,&parser);
	if (rc){
		_printErrorInfo(infoResult, "create parser error");
		if(infoResult!=NULL)
			fclose(infoResult);
		return 1;
	}

	_printInfo(infoResult, "Origin SQL:\n");
	_printInfo(infoResult, sqlText);
	_printInfo(infoResult, "\n\n");
	rc= gsp_check_syntax(parser, sqlText);
	if (rc != 0){
		_printErrorInfo(infoResult, "parser error:%s\n",gsp_errmsg(parser));
		gsp_parser_free(parser);
		if(infoResult!=NULL)
			fclose(infoResult);
		return 1;
	}

	tableInfoList = createStringList(TRUE);
	fieldInfoList = createStringList(TRUE);
	
	traverser = createSqlTraverser();
	visitor = _createVisitor(traverser);

	for(index=0;index<parser->nStatement;index++ ){
		nodeList = traverser->traverseSQL(traverser, &parser->pStatement[index]);
		iter = nodeList->getIterator(nodeList);
		while(nodeList->hasNext(nodeList,&iter))
		{
			gsp_node *node = (gsp_node*)nodeList->next(&iter);
			if(node->nodeType == t_gsp_table )
				visitor->handle_node[t_gsp_table](node, visitor);
			else if(node->nodeType == t_gsp_fromTable )
				visitor->handle_node[t_gsp_fromTable](node, visitor);
		}
	}

	if(!tableInfoList->isEmpty(tableInfoList)){
		tableInfoList->sort(tableInfoList);
		_printInfo(infoResult, "\nTables:\n");
		_printListInfo(tableInfoList, infoResult);
	}

	if(!fieldInfoList->isEmpty(fieldInfoList)){
		fieldInfoList->sort(fieldInfoList);
		_printInfo(infoResult,"Columns:\n");
		_printListInfo(fieldInfoList, infoResult);
	}

	tableInfoList->dispose(tableInfoList);
	fieldInfoList->dispose(fieldInfoList);
	
	traverser->dispose(traverser);
	_disposeVisitor(visitor);
	if(infoResult!=NULL)
		fclose(infoResult);
	gsp_parser_free(parser);
	return 0;
}
