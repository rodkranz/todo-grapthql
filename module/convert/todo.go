package convert

import (
	"github.com/rodlps22/todo-grapthql/model"
)

type Todo struct {
	ID   string     `json:"id"`
	Text string     `json:"text"`
	Done bool       `json:"done"`
	Lang *Language  `json:"language"`
}

func ToTodo(td *model.Todo) *Todo {
	todo := &Todo{
		ID   : td.ID.Hex(),
		Text : td.Text,
		Done : td.Done,
	}

	if td.Language != nil {
		todo.Lang = ToLanguage(td.Language);
	}

	return todo
}

func ToTodos(tds []*model.Todo) []*Todo {
	ts := make([]*Todo, len(tds))
	for i, t := range tds {
		ts[i] = ToTodo(t)
	}
	return ts
}

