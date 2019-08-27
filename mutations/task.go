package mutations

import (
	"context"
	"errors"
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
				log.Printf("Error can't insert task data %s \n", err.Error())
				return nil, errors.New(err.Error())
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

func (mutation *Mutations) UpdateTaskItem() *graphql.Field {
	return &graphql.Field{
		Name: "UpdateTaskItem",
		Type: types.TaskItemType,
		Args: graphql.FieldConfigArgument{
			"taskid": &graphql.ArgumentConfig{Type: graphql.String},
			"input":  &graphql.ArgumentConfig{Type: types.TaskItemInputType},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := mutation.Database.Collection("task")
			var TaskItem types.Task_Item

			id, _ := primitive.ObjectIDFromHex(params.Args["taskid"].(string))
			task := params.Args["input"].(map[string]interface{})

			// cara pertama
			filter := bson.D{
				primitive.E{
					Key: "$and",
					Value: []bson.M{
						bson.M{"_id": id},
						bson.M{"taskitem.stepnum": task["stepnum"]},
					},
				},
			}

			update := bson.D{
				primitive.E{
					Key: "$set",
					Value: bson.D{
						primitive.E{
							Key:   "taskitem.$",
							Value: task,
						},
					},
				},
			}

			// cara kedua
			/*
				filter := bson.D{
					primitive.E{
						Key: "$and",
						Value: []bson.D{
							primitive.D{
								primitive.E{
									Key:   "_id",
									Value: id,
								},
								primitive.E{
									Key:   "taskitem.stepnum",
									Value: task["stepnum"],
								},
							},
						},
					},
				}
			*/

			err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&TaskItem)
			if err != nil {
				log.Printf("Error can't update data %s \n", err.Error())
				return nil, errors.New(err.Error())
			}

			return task, nil
		},
	}
}

func (mutation *Mutations) UpdateStatusComplete() *graphql.Field {
	return &graphql.Field{
		Name: "UpdateStatusComplete",
		Type: types.MessageType,
		Args: graphql.FieldConfigArgument{
			"taskid":  &graphql.ArgumentConfig{Type: graphql.String},
			"stepnum": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := mutation.Database.Collection("task")
			taskid, _ := primitive.ObjectIDFromHex(params.Args["taskid"].(string))
			stepnum := params.Args["stepnum"]

			var Task types.Task
			var message types.Message

			filter := bson.D{
				primitive.E{
					Key: "$and",
					Value: []bson.D{
						primitive.D{
							primitive.E{
								Key:   "_id",
								Value: taskid,
							},
							primitive.E{
								Key:   "taskitem.stepnum",
								Value: stepnum,
							},
						},
					},
				},
			}

			update := bson.D{
				primitive.E{
					Key: "$set",
					Value: bson.D{
						primitive.E{
							Key:   "taskitem.$.complete",
							Value: true,
						},
					},
				},
			}

			err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&Task)
			if err != nil {
				log.Printf("error can't find task %s \n", err.Error())
				return nil, errors.New(err.Error())
			}
			message.Message = "Success updated status task item"

			return message, nil
		},
	}
}

func (mutation *Mutations) DeleteTaskItem() *graphql.Field {
	return &graphql.Field{
		Name: "DeleteTaskItem",
		Type: types.MessageType,
		Args: graphql.FieldConfigArgument{
			"stepnum": &graphql.ArgumentConfig{Type: graphql.Int},
			"taskid":  &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := mutation.Database.Collection("task")
			taskid, _ := primitive.ObjectIDFromHex(params.Args["taskid"].(string))
			stepnum := params.Args["stepnum"]

			var Task types.Task
			var message types.Message

			filter := bson.D{
				primitive.E{
					Key: "$and",
					Value: []bson.D{
						primitive.D{
							primitive.E{
								Key:   "_id",
								Value: taskid,
							},
							primitive.E{
								Key:   "taskitem.stepnum",
								Value: stepnum,
							},
						},
					},
				},
			}

			delete := bson.D{
				primitive.E{
					Key: "$pull",
					Value: bson.D{
						primitive.E{
							Key: "taskitem",
							Value: bson.D{
								primitive.E{
									Key:   "stepnum",
									Value: stepnum,
								},
							},
						},
					},
				},
			}

			err := collection.FindOneAndUpdate(context.Background(), filter, delete).Decode(&Task)
			if err != nil {
				log.Printf("Error %s \n", err.Error())
				return nil, errors.New(err.Error())
			}

			message.Message = "Success delete Task item!"

			return message, nil
		},
	}
}

func (mutation *Mutations) DeleteTask() *graphql.Field {
	return &graphql.Field{
		Name: "DeleteTask",
		Type: types.MessageType,
		Args: graphql.FieldConfigArgument{
			"taskid": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := mutation.Database.Collection("task")
			taskid, _ := primitive.ObjectIDFromHex(params.Args["taskid"].(string))

			var Task types.Task
			var message types.Message

			filter := bson.D{
				primitive.E{
					Key:   "_id",
					Value: taskid,
				},
			}

			err := collection.FindOneAndDelete(context.Background(), filter).Decode(&Task)

			if err != nil {
				log.Printf("Error %s \n", err.Error())
				return nil, errors.New(err.Error())
			}

			message.Message = "Success delete Task!"

			return message, nil
		},
	}
}

func (mutation *Mutations) UpdateStatusTask() *graphql.Field {
	return &graphql.Field{
		Name: "UpdateStatusTask",
		Type: types.MessageType,
		Args: graphql.FieldConfigArgument{
			"taskid": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			collection := mutation.Database.Collection("task")
			taskid, _ := primitive.ObjectIDFromHex(params.Args["taskid"].(string))

			var task types.Task
			var message types.Message

			filter := bson.D{
				primitive.E{
					Key:   "_id",
					Value: taskid,
				},
			}

			update := bson.D{
				primitive.E{
					Key:   "status",
					Value: "done",
				},
			}

			err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&task)
			if err != nil {
				log.Printf("Error : %s \n", err.Error())
				return nil, errors.New(err.Error())
			}

			message.Message = "Task `Status` has been updated!"

			return message, nil
		},
	}
}
