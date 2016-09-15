package objects

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/graphql-go/graphql"
)

type Todo struct {
	ID   bson.ObjectId `json:"id"   bson:"_id,omitempty"`
	Text string        `json:"text" bson:"text"`
	Done bool          `json:"done" bson:"done"`
}

var TodoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

