package types

import (
	"github.com/graphql-go/graphql"
)

type Task_Item struct {
	StepNum  int    `json:"stepnum"`
	Item     string `json:"taskitem"`
	Complete bool   `json:"complete"`
}

type Task struct {
	Task_ID          string      `json:"_id"`
	Task_Username    string      `json:"taskusername" bson:"taskusername"`
	Task_Description string      `json:"taskdescription" bson:"taskdescription"`
	Task_Item        []Task_Item `json:"taskitem" bson:"taskitem"`
	DueDate          string      `json:"duedate" bson:"duedate"`
	Status           string      `json:"status" bson:"status"`
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
		"taskusername":    &graphql.Field{Type: graphql.String},
		"taskdescription": &graphql.Field{Type: graphql.String},
		"taskitem":        &graphql.Field{Type: graphql.NewList(TaskItemType)},
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
		},
	},
)
