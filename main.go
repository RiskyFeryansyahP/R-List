package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/confus1on/R-List/database"
	"github.com/confus1on/R-List/mutations"
	"github.com/confus1on/R-List/queries"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	fmt.Println("R-List Building Here...!")
	db, err := database.GetMongo()
	if err != nil {
		log.Fatalf("Error Can't connect into Database %s", err)
	}

	queries, _ := queries.New(db)
	mutations, _ := mutations.New(db)

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/", httpHandler)
	log.Print("Ready Listening....\n")

	http.ListenAndServe(":8888", nil)

}
