package gqlserver

import (
	"github.com/graphql-go/graphql"
)

type SchemaWrapper struct {
	productResolver *Resolver
	Schema          graphql.Schema
}

func NewSchemaWrapper() *SchemaWrapper {
	return &SchemaWrapper{}
}

func (s *SchemaWrapper) WithProductResolver(pr *Resolver) *SchemaWrapper {
	s.productResolver = pr

	return s
}

func (s *SchemaWrapper) Init() error {
	// add gql schema as needed
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:        "ProductGetter",
			Description: "All query related to getting product data",
			Fields: graphql.Fields{
				"ProductDetail": &graphql.Field{
					Type:        ProductType,
					Description: "Get product by ID",
					Args: graphql.FieldConfigArgument{
						"product_id": &graphql.ArgumentConfig{
							Type: graphql.Int,
						},
					},
					Resolve: s.productResolver.GetProduct(),
				},
			},
		}),
		// uncomment this and add objects for mutation
		// Mutation: graphql.NewObject(graphql.ObjectConfig{}),
	})

	if err != nil {
		return err
	}

	s.Schema = schema

	return nil
}
