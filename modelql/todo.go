package modelql

import (
	"github.com/graphql-go/graphql"
	"github.com/rodlps22/todo-grapthql/model"
	"github.com/rodlps22/todo-grapthql/module/convert"
)

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
		"language": &graphql.Field{
			Type: LanguageType,
		},
	},
})

func TodoQl() *graphql.Field {
	return &graphql.Field{
		Type: TodoType,
		Description: "get todo by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)

			result, err := model.FindTodoById(id)
			if err != nil {
				return nil, err
			}

			return convert.ToTodo(result), nil
		},
	}
}

func TodoListQl() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(TodoType),
		Description: "List of todos",
		Args: graphql.FieldConfigArgument{
			"isDone": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"lang": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			where := make(map[string]interface{})

			if pLang, isOK := params.Args["lang"].(string); isOK {
				lang, err := model.FindLanguageByName(pLang);
				if err != nil {
					return nil, err
				}

				where["langId"] = lang.GetID()
			}

			if pDone, isOK := params.Args["isDone"].(bool); isOK {
				where["done"] = pDone;
			}

			var count int = 0
			if pFirst, isOK := params.Args["limit"].(int); isOK {
				count = pFirst;
			}

			results, err := model.FindAllTodo(where, count)
			if err != nil {
				return nil, err
			}
			return convert.ToTodos(results), nil
		},
	}
}

func CreateTodo() *graphql.Field {
	return &graphql.Field{
		Type: TodoType, // the return type for this field
		Description: "Create new todo",
		Args: graphql.FieldConfigArgument{
			"text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"done": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// marshall and cast the argument value
			text, _ := params.Args["text"].(string)
			done, hd := params.Args["done"].(bool)

			if !hd {
				done = false
			}

			newTodo := &model.Todo{
				Text: text,
				Done: done,
			}

			err := model.CreateTodo(newTodo)
			if err != nil {
				return nil, err
			}

			return convert.ToTodo(newTodo), nil
		},
	}
}

func UpdateTodo() *graphql.Field {
	return &graphql.Field{
		Type: TodoType,
		Description: "Update todo by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"text": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"done": &graphql.ArgumentConfig{
				Type: graphql.Boolean,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// marshall and cast the argument value
			id, _ := params.Args["id"].(string)
			text, _ := params.Args["text"].(string)
			done, hd := params.Args["done"].(bool)

			if !hd {
				done = false
			}

			editTodo, err := model.FindTodoById(id)
			if err != nil {
				return nil, err
			}

			editTodo.Text = text
			editTodo.Done = done

			err = model.UpdateTodo(editTodo)
			if err != nil {
				return nil, err
			}

			return convert.ToTodo(editTodo), nil
		},
	}
}

func DeleteTodo() *graphql.Field {
	return &graphql.Field{
		Type: TodoType,
		Description: "Delete todo by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)

			deleted, err := model.DeleteTodoById(id)
			if err != nil {
				return nil, err
			}

			return convert.ToTodo(deleted), nil
		},
	}
}
