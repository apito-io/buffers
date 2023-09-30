package interfaces

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/ast"
)

type GqlServer interface {
	SolvePublicQuery(ctx context.Context, model string, _args map[string]interface{}, selectionSet *ast.SelectionSet, router echo.Context) ([]byte, error)
	SolvePublicQueryCount(ctx context.Context, model string, _args map[string]interface{}, router echo.Context) ([]byte, error)
}
