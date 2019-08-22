package queries

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/confus1on/R-List/types"
	"golang.org/x/crypto/bcrypt"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (query *Queries) FindUserQuery() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			username, ok := params.Args["username"].(string)
			if ok {
				collection := query.Database.Collection("user")
				var User types.User

				filter := bson.D{primitive.E{Key: "username", Value: username}}

				err := collection.FindOne(context.Background(), filter).Decode(&User)
				if err != nil {
					log.Fatalf("Error FindOne Collection %s", err.Error())
					return nil, nil
				}
				return User, nil
			}
			return nil, nil
		},
	}
}

func (queries *Queries) UserAuthentication() *graphql.Field {
	return &graphql.Field{
		Name: "FindUser",
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{Type: graphql.String},
			"password": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := queries.Database.Collection("user")
			username := params.Args["username"].(string)
			password := params.Args["password"].(string)

			var User types.User

			filter := bson.D{
				primitive.E{
					Key:   "username",
					Value: username,
				},
			}

			err := collection.FindOne(context.Background(), filter).Decode(&User)
			if err != nil {
				fmt.Printf("Error %s \n", err.Error())
				return nil, errors.New("Can't find username")
			}

			err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(password))
			if err == nil {
				return User, nil
			}

			return nil, errors.New("Password is not match")
		},
	}
}

func (query *Queries) GetAllUsers() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := query.Database.Collection("user")
			var users []types.User

			result, err := collection.Find(context.Background(), bson.M{})
			if err != nil {
				log.Fatalf("Error can't Find data users %s", err.Error())
				return nil, err
			}
			defer result.Close(context.Background())

			for result.Next(context.Background()) {
				var user types.User
				err = result.Decode(&user)
				if err != nil {
					log.Fatalf("Error Decode %s", err.Error())
					return nil, nil
				}

				users = append(users, user)
			}

			return users, nil
		},
	}
}
