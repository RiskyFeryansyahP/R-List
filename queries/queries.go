package queries

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/mongo"
)

type Queries struct {
	Database *mongo.Database
}

func New(database *mongo.Database) (queries *Queries, err error) {
	queries = &Queries{Database: database}
	fmt.Println(queries)
	return queries, nil
}

func (query *Queries) GetRootFields() graphql.Fields {
	return graphql.Fields{
		"user":  query.FindUserQuery(),
		"users": query.GetAllUsers(),
		"task":  query.SelectTask(),
		"tasks": query.GetTasks(),
	}
}
