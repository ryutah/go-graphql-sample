package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Info  string  `json:"info,omitempty"`
	Price float64 `json:"price"`
}

var products []Product

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.Int},
			"name":  &graphql.Field{Type: graphql.String},
			"info":  &graphql.Field{Type: graphql.String},
			"price": &graphql.Field{Type: graphql.Float},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// Get (read) single product by id
			// http://localhost:8080/product?query=(product(id:1){name,info,price})
			"product": &graphql.Field{
				Type:        productType,
				Description: "Get product by id",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						// Find product
						for _, product := range products {
							if int(product.ID) == id {
								return product, nil
							}
						}
					}
					return nil, nil
				},
			},
			// Get (read) product list
			// http://localhost:8080/product?query=(list{name,info,price})
			"list": &graphql.Field{
				Type:        graphql.NewList(productType),
				Description: "Get product list",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return products, nil
				},
			},
		},
	},
)

var mutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		// Create new product item
		// http://localhost:8080/product?query=mutation+_{create(name:"Inca Kola",info:"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)",price:1.99){id,name,info,price}}
		"create": &graphql.Field{
			Type:        productType,
			Description: "Create new product",
			Args: graphql.FieldConfigArgument{
				"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"info":  &graphql.ArgumentConfig{Type: graphql.String},
				"price": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				rand.Seed(time.Now().UnixNano())
				product := Product{
					ID:    int64(rand.Intn(100000)), // generate random ID
					Name:  p.Args["name"].(string),
					Info:  p.Args["info"].(string),
					Price: p.Args["price"].(float64),
				}
				products = append(products, product)
				return product, nil
			},
		},
		// Update product by id
		// http://localhost:8080/product?query=mutation+_{update(id:1,price:3.95){id,name,info,price}}
		"update": &graphql.Field{
			Type:        productType,
			Description: "Update product by id",
			Args: graphql.FieldConfigArgument{
				"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				"name":  &graphql.ArgumentConfig{Type: graphql.String},
				"info":  &graphql.ArgumentConfig{Type: graphql.String},
				"price": &graphql.ArgumentConfig{Type: graphql.Float},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var (
					id, _          = p.Args["id"].(int)
					name, nameOk   = p.Args["name"].(string)
					info, infoOk   = p.Args["info"].(string)
					price, priceOk = p.Args["price"].(float64)
				)
				product := Product{}
				for i := range products {
					p := &products[i]
					if int64(id) == p.ID {
						if nameOk {
							p.Name = name
						}
						if infoOk {
							p.Info = info
						}
						if priceOk {
							p.Price = price
						}
						product = *p
						break
					}
				}
				return product, nil
			},
		},
		"delete": &graphql.Field{
			Type:        productType,
			Description: "Delete product by id",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, _ := p.Args["id"].(int)
				product := Product{}
				for i, p := range products {
					if int64(id) == p.ID {
						product = p
						products = append(products[:i], products[i+1:]...)
					}
				}
				return product, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	},
)

func executeQuery(req *request, scheme graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:         scheme,
		RequestString:  req.Query,
		VariableValues: req.Variables,
	})
	if len(result.Errors) > 0 {
		log.Printf("errors: %v", result.Errors)
	}
	return result
}

func initProductsData(p *[]Product) {
	var (
		product1 = Product{
			ID: 1, Name: "Chicha Morada", Price: 7.99,
			Info: "Chicha morada is a beverage originated in the Andean regions of Perú but is actually consumed at a national level (wiki)",
		}
		product2 = Product{
			ID: 2, Name: "Chicha de jora", Price: 5.95,
			Info: "Chicha de jora is a corn beer chicha prepared by germinating maize, extracting the malt sugars, boiling the wort, and fermenting it in large vessels (traditionally huge earthenware vats) for several days (wiki)",
		}
		product3 = Product{
			ID: 3, Name: "Pisco", Price: 9.95,
			Info: "Pisco is a colorless or yellowish-to-amber colored brandy produced in winemaking regions of Peru and Chile (wiki)",
		}
	)
	*p = append(*p, product1, product2, product3)
}

type request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

func parseRequest(r *http.Request) (*request, error) {
	req := new(request)
	err := json.NewDecoder(r.Body).Decode(req)
	return req, err
}

func main() {
	initProductsData(&products)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/static/app.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "app.js")
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		req, err := parseRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		result := executeQuery(req, schema)
		resp, _ := json.MarshalIndent(result, "", "  ")
		fmt.Fprintf(w, "%s\n", resp)
	})

	log.Println("Start server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
