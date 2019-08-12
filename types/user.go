package types

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullName"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"username": &graphql.Field{Type: graphql.String},
		"password": &graphql.Field{Type: graphql.String},
		"fullname": &graphql.Field{Type: graphql.String},
	},
})
