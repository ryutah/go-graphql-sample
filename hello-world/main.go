package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

type Foo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				log.Println(p.Args)
				return "world", nil
			},
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					DefaultValue: "111",
					Type:         graphql.String,
				},
			},
		},
		"foo": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "foo",
				Fields: graphql.Fields{
					"id":   &graphql.Field{Type: graphql.Int},
					"name": &graphql.Field{Type: graphql.String},
				},
			}),
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.Int},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return &Foo{
					ID:   1,
					Name: "hoge",
				}, nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	query := `
	{
		hello(id: "1233"),
		foo {
			id,
			name
		}
	}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.MarshalIndent(r, "", "  ")
	fmt.Printf("%s\n", rJSON)
}
