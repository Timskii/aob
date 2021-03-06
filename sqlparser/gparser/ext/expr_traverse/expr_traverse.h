#ifndef _EXPR_TRAVERSE_H
#define _EXPR_TRAVERSE_H

#ifdef __cplusplus
extern "C" {
#endif

#include "../../core/gsp_base.h"
#include "../../core/gsp_list.h"
#include "../../core/gsp_sourcetoken.h"
#include "../../core/gsp_node.h"
#include "../../core/gsp_sqlparser.h"
#include "../collection/includes/lifo_stack.h"
#include "../collection/includes/memento.h"
/*
 ** Return code from the expression tree walking primitives and their
 ** callbacks.
*/

typedef enum gsp_walking_result{
	gsp_walking_continue,	/* Continue down into children */
	gsp_walking_prune,		/* Omit children but continue walking siblings */
	gsp_walking_abort,		/* Abandon the tree walk */
}gsp_walking_result;

typedef struct gsp_walker {
  void *context;
  gsp_walking_result (*exprVisit)(struct gsp_walker*, gsp_expr *pExpr, int isLeafNode);     /* Callback for expressions */
}ExprTraverser;

gsp_walking_result preOrderTraverse (ExprTraverser *pTraverser, gsp_expr *pExpr);
gsp_walking_result inOrderTraverse  (ExprTraverser *pTraverser, gsp_expr *pExpr);
gsp_walking_result postOrderTraverse(ExprTraverser *pTraverser, gsp_expr *pExpr);

#ifdef __cplusplus
}
#endif

#endif /* _EXPR_TRAVERSE_H */
