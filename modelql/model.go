package modelql

import (
	"github.com/graphql-go/graphql"
	"fmt"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Description: "List methods",
	Fields: graphql.Fields{
		/*
		   curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,text,done}}'
		*/
		"todo": TodoQl(),

		/*
		   curl -g 'http://localhost:8080/graphql?query={todoList{id,text,done}}'
		*/
		"todoList": TodoListQl(),

		/*
		   curl -g 'http://localhost:8080/graphql?query={langList{id,name}}'
		*/
		"languageList": LanguageListQl(),

		/*
		   curl -g 'http://localhost:8080/graphql?query={todo(id:"b"){id,nome}}'
		*/
		"language": LanguageQl(),
	},
})


// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Description: "Change methods",
	Fields: graphql.Fields{
		/*
		   curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:"My+new+todo"){id,text,done}}'
		*/
		"createTodo": CreateTodo(),

		/*
		   curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:"57bf797df0c011702af7739c",text:"My+new+todo",done:true){id,text,done}}'
		*/
		"updateTodo": UpdateTodo(),

		/*
		   curl -g 'http://localhost:8080/graphql?query=mutation+_{deleteTodo(id:"57bf797df0c011702af7739c"){id,text,done}}'
		*/
		"deleteTodo": DeleteTodo(),
	},
})

// define schema, with our rootQuery and rootMutation
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}


