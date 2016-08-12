package gql

import (
	"github.com/gocraft/web"
	"fmt"
	"github.com/minivac/lebroc/app/core"
	"github.com/graphql-go/graphql"
	"github.com/minivac/lebroc/app/models"
	"net/http"
)

type GQLContext struct {
	*core.Context
}

var (
	bookType *graphql.Object
	thematicType *graphql.Object
	imageType *graphql.Object
	queryType *graphql.Object
	schema graphql.Schema
)

func Configure(router *web.Router) {

	imageType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Image",
		Description: "An image, with width a height",
		Fields: graphql.Fields{
			"url": &graphql.Field{Type: graphql.String},
			"width": &graphql.Field{Type: graphql.Int},
			"height": &graphql.Field{Type: graphql.Int},
		},
	})

	bookType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Book",
		Description: "A book in the catalog",
		Fields: graphql.Fields{
			"id": &graphql.Field{Type: graphql.String },
			"title_text": &graphql.Field{Type: graphql.String},
			"publisher_name": &graphql.Field{Type: graphql.String },
			"imprint_name": &graphql.Field{Type: graphql.String },
			"languages": &graphql.Field{Type: graphql.NewList(graphql.String) },
			"thematics": &graphql.Field{Type: graphql.NewList(graphql.String) },
			"cover": &graphql.Field{
				Type: imageType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return p.Source.(models.Book).Cover, nil
				},
			},
		},
	})

	thematicType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Thematic",
		Description: "A thematic, with a description and a cover",
		Fields: graphql.Fields{
			"id": &graphql.Field{Type: graphql.String, },
			"name":&graphql.Field{Type: graphql.String, },
			"image":&graphql.Field{Type: imageType},
			"books":&graphql.Field{
				Type: graphql.NewList(bookType),
				Args: graphql.FieldConfigArgument{
					"offset" : &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue:0},
					"limit" : &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue:4},
				},
				Resolve:func(p graphql.ResolveParams) (interface{}, error) {
					limit := p.Args["limit"].(int)
					offset := p.Args["offset"].(int)
					id := p.Source.(models.Thematic).Id
					return core.BookService.Repo.FindBooksWithThematic(id, offset, limit)
				},
			},
		},
	})

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id" : &graphql.ArgumentConfig{
						Type: graphql.String,
						Description: "Book id, in the form b0, b1...",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return core.BookService.Repo.FindBook(p.Args["id"].(string))
				},
			},
			"thematics" : &graphql.Field{
				Type: graphql.NewList(thematicType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return core.BookService.Repo.FindAllThematics()
				},
			},
		},
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	gqlRouter := router.Subrouter(GQLContext{}, "/gql")
	gqlRouter.Get("/", (*GQLContext).handleGraphQuery)
}

func (c *GQLContext) handleGraphQuery(rw web.ResponseWriter, req *web.Request) {
	query := req.URL.Query()["query"][0]
	if (query == "") {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	result := graphql.Do(graphql.Params{
		Schema: schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	fmt.Fprint(rw, models.ToPrettyJsonString(result))
}
