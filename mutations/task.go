package mutations

import (
	"context"
	"log"

	"github.com/confus1on/R-List/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/graphql-go/graphql"
)

func (mutation *Mutations) CreateTask() *graphql.Field {
	return &graphql.Field{
		Name: "CreateTask",
		Type: types.TaskType,
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{Type: types.TaskInputType},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := mutation.Database.Collection("task")
			task := params.Args["input"]

			_, err := collection.InsertOne(context.Background(), task)
			if err != nil {
				log.Fatalf("Error can't insert task data %s \n", err.Error())
			}

			return task, nil
		},
	}
}

func (mutation *Mutations) CreateTaskItem() *graphql.Field {
	return &graphql.Field{
		Name: "CreateTaskItem",
		Type: types.TaskItemType,
		Args: graphql.FieldConfigArgument{
			"input":  &graphql.ArgumentConfig{Type: types.TaskItemInputType},
			"taskid": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := mutation.Database.Collection("task")

			task := params.Args["input"].(map[string]interface{})
			id, _ := primitive.ObjectIDFromHex(params.Args["taskid"].(string))

			filter := bson.D{
				primitive.E{
					Key:   "_id",
					Value: id,
				},
			}

			push := bson.D{
				primitive.E{
					Key: "$push",
					Value: bson.D{
						primitive.E{
							Key:   "taskitem",
							Value: task,
						},
					},
				},
			}

			_ = collection.FindOneAndUpdate(context.Background(), filter, push)

			return task, nil
		},
	}
}
