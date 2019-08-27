package mutations

import (
	"github.com/graphql-go/graphql"

	"go.mongodb.org/mongo-driver/mongo"
)

type Mutations struct {
	Database *mongo.Database
}

func New(database *mongo.Database) (mutations *Mutations, err error) {
	mutations = &Mutations{Database: database}
	return mutations, nil
}

func (mutation *Mutations) GetRootFields() graphql.Fields {
	return graphql.Fields{
		"createuser":           mutation.CreateUser(),
		"createtask":           mutation.CreateTask(),
		"deletetask":           mutation.DeleteTask(),
		"updatestatustask":     mutation.UpdateStatusTask(),
		"createtaskitem":       mutation.CreateTaskItem(),
		"updatetaskitem":       mutation.UpdateTaskItem(),
		"updatestatuscomplete": mutation.UpdateStatusComplete(),
		"deletetaskitem":       mutation.DeleteTaskItem(),
	}
}
