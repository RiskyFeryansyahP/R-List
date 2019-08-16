package queries

import (
	"context"
	"fmt"
	"log"

	"github.com/confus1on/R-List/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/graphql-go/graphql"
)

func (queries *Queries) SelectTask() *graphql.Field {
	return &graphql.Field{
		Name: "SelectTask",
		Type: types.TaskType,
		Args: graphql.FieldConfigArgument{
			"taskid": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := queries.Database.Collection("task")
			var Task types.Task

			id, err := primitive.ObjectIDFromHex(params.Args["taskid"].(string))
			if err != nil {
				log.Fatalf("Error ObjectID %v \n", err.Error())
			}

			fmt.Println(id)

			filter := bson.D{
				primitive.E{
					Key:   "_id",
					Value: id,
				},
			}

			err = collection.FindOne(context.Background(), filter).Decode(&Task)
			if err != nil {
				log.Fatalf("Error can't find task data \n %s", err.Error())
			}

			return Task, nil
		},
	}
}
