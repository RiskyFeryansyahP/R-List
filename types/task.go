package types

import (
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task_Item struct {
	StepNum  int    `json:"stepnum" bson:"stepnum"`
	Item     string `json:"item" bson:"item"`
	Complete bool   `json:"complete" bson:"complete"`
}

type Task struct {
	Task_ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Task_Username    string             `json:"taskusername" bson:"taskusername"`
	Task_Description string             `json:"taskdescription" bson:"taskdescription"`
	Task_Item        []Task_Item        `json:"taskitem" bson:"taskitem"`
	DueDate          string             `json:"duedate" bson:"duedate"`
	Status           string             `json:"status" bson:"status"`
}

var TaskItemType = graphql.NewObject(graphql.ObjectConfig{
	Name: "TaskItem",
	Fields: graphql.Fields{
		"stepnum":  &graphql.Field{Type: graphql.Int},
		"item":     &graphql.Field{Type: graphql.String},
		"complete": &graphql.Field{Type: graphql.Boolean},
	},
})

var TaskType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Task",
	Fields: graphql.Fields{
		"_id":             &graphql.Field{Type: BSON},
		"taskusername":    &graphql.Field{Type: graphql.String},
		"taskdescription": &graphql.Field{Type: graphql.String},
		"taskitem":        &graphql.Field{Type: graphql.NewList(TaskItemType)},
		"duedate":         &graphql.Field{Type: graphql.String},
		"status":          &graphql.Field{Type: graphql.String},
	},
})

var TaskItemInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "TaskItemInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"stepnum": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
			"item": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"complete": &graphql.InputObjectFieldConfig{
				Type: graphql.Boolean,
			},
		},
	},
)

var TaskInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "TaskInputType",
		Fields: graphql.InputObjectConfigFieldMap{
			"taskusername": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"taskdescription": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"duedate": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"status": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
