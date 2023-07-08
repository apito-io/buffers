package extensions

import (
	"github.com/graphql-go/graphql"
	"github.com/labstack/echo/v4"
)

type ThirdPartyGraphQLSchemas struct {
	Queries   graphql.Fields
	Mutations graphql.Fields
}

type ThirdPartyRESTApi struct {
	Path       string
	Method     string
	Controller func(c echo.Context) error
}
