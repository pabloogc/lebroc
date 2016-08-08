package gql

import (
	"github.com/gocraft/web"
	"fmt"
	"github.com/minivac/lebroc/app/core"
	"github.com/graphql-go/graphql"
	"github.com/minivac/lebroc/app/models"
)

type GQLContext struct {
	*core.Context
}

var BookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "book",
	Description: "A book in the catalog",
	Fields: graphql.Fields{
		"id": &graphql.Field{Type: graphql.String, },
		"title_text": &graphql.Field{Type: graphql.String, },
		"publisher_name": &graphql.Field{Type: graphql.String, },
		"imprint_name": &graphql.Field{Type: graphql.String, },
		"languages": &graphql.Field{Type: graphql.NewList(graphql.String)},
	},
})

var QueryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"book": &graphql.Field{
			Type: BookType,
			Args: graphql.FieldConfigArgument{
				"id" : &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				book, err := core.BookService.Repo.FindBook(p.Args["id"].(string))
				return book, err
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: QueryType,
})

func (c *GQLContext) rootGQL(rw web.ResponseWriter, req *web.Request) {
	q := req.URL.Query()["query"][0]
	result := graphql.Do(graphql.Params{
		Schema: Schema,
		RequestString: q,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	fmt.Fprint(rw, models.ToJsonString(result))
}

func Configure(router *web.Router) {
	gqlRouter := router.Subrouter(GQLContext{}, "/gql")
	gqlRouter.Get("/", (*GQLContext).rootGQL)
}
