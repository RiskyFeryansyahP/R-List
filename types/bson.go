package types

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var BSON = graphql.NewScalar(graphql.ScalarConfig{
	Name:        "BSON",
	Description: "BSON Scalar to represent a BSON ObjectID",
	// Serialize serializes `bson.ObjectID` to `string`.
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case primitive.ObjectID:
			return value.Hex()
		case *primitive.ObjectID:
			v := *value
			return v.Hex()
		default:
			return nil
		}
	},
	// ParseValue parses GraphQL variables from `string` to `bson.ObjectID`
	ParseValue: func(value interface{}) interface{} {
		switch value := value.(type) {
		case string:
			v, _ := primitive.ObjectIDFromHex(value)
			return v
		case *string:
			v, _ := primitive.ObjectIDFromHex(*value)
			return v
		default:
			return nil
		}
	},
	// ParseLiteral parses GraphQL AST to `bson.ObjectID`
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			vAST, _ := primitive.ObjectIDFromHex(valueAST.Value)
			return vAST
		}
		return nil
	},
})
