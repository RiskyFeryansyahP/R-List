package mutations

import (
	"context"
	"errors"
	"log"

	"github.com/confus1on/R-List/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	"github.com/graphql-go/graphql"
)

func (mutation *Mutations) CreateUser() *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{Type: types.UserInputType},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := mutation.Database.Collection("user")

			users := params.Args["input"].(map[string]interface{})

			filter := bson.D{
				primitive.E{
					Key:   "username",
					Value: users["username"],
				},
			}

			result, _ := collection.CountDocuments(context.Background(), filter)
			if result > 0 {
				return nil, errors.New("Couldn't use this username")
			}

			// hashing password
			bytes, err := bcrypt.GenerateFromPassword([]byte(users["password"].(string)), 14)
			if err != nil {
				log.Fatalf("Error Can't Hashing Password %s", err.Error())
			}
			users["password"] = string(bytes) // hash password with bcrypt

			_, err = collection.InsertOne(context.Background(), users)
			if err != nil {
				log.Fatalf("Error Insert Into Collection %s \n", err.Error())
			}

			return users, nil
		},
	}
}
