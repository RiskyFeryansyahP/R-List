package types

import "github.com/graphql-go/graphql"

type Message struct {
	Message string `json:"message" bson:"message"`
}

var MessageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "MessageType",
	Fields: graphql.Fields{
		"message": &graphql.Field{Type: graphql.String},
	},
})
