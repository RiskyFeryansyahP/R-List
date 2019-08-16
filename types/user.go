package types

import (
	"github.com/graphql-go/graphql"
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	FullName string `json:"fullname" bson:"fullname"`
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"username": &graphql.Field{Type: graphql.String},
		"password": &graphql.Field{Type: graphql.String},
		"fullname": &graphql.Field{Type: graphql.String},
	},
})

var UserInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"username": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"fullname": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
