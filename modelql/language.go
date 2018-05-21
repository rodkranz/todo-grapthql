package modelql

import (
	"github.com/graphql-go/graphql"

	"github.com/rodkranz/todo-grapthql/model"
	"github.com/rodkranz/todo-grapthql/module/convert"
)

var LanguageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Language",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

func LanguageQl() *graphql.Field {
	return &graphql.Field{
		Type:        LanguageType,
		Description: "Get one language by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)
			result, err := model.FindLanguageById(id)
			if err != nil {
				return nil, err
			}

			return convert.ToLanguage(result), nil
		},
	}
}

func LanguageListQl() *graphql.Field {
	return &graphql.Field{
		Type:        graphql.NewList(LanguageType),
		Description: "List of languages",
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var count int = 0
			if pFirst, isOK := params.Args["limit"].(int); isOK {
				count = pFirst
			}

			result, err := model.FindAllLanguages(count)
			return convert.ToLanguages(result), err
		},
	}
}

func CreateLanguage() *graphql.Field {
	return &graphql.Field{
		Type:        LanguageType,
		Description: "Create new language",
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// marshall and cast the argument value
			name, _ := params.Args["text"].(string)

			newLanguage := &model.Language{
				Name: name,
			}

			err := model.CreateLanguage(newLanguage)
			return convert.ToLanguage(newLanguage), err
		},
	}
}

func UpdateLanguage() *graphql.Field {
	return &graphql.Field{
		Type:        TodoType,
		Description: "Update language",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			// marshall and cast the argument value
			id, _ := params.Args["id"].(string)
			name, _ := params.Args["text"].(string)
			editLanguage, err := model.FindLanguageById(id)
			if err != nil {
				return nil, err
			}

			editLanguage.Name = name

			err = model.UpdateLanguage(editLanguage)
			return convert.ToLanguage(editLanguage), err
		},
	}
}

func DeleteLanguage() *graphql.Field {
	return &graphql.Field{
		Type:        TodoType,
		Description: "Delete language by id",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			id, _ := params.Args["id"].(string)

			deleted, err := model.DeleteLanguageById(id)
			return convert.ToLanguage(deleted), err
		},
	}
}
