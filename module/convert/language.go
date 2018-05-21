package convert

import (
	"github.com/rodkranz/todo-grapthql/model"
)

type Language struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ToLanguage(td *model.Language) *Language {
	return &Language{
		ID:   td.ID.Hex(),
		Name: td.Name,
	}
}

func ToLanguages(tds []*model.Language) []*Language {
	ts := make([]*Language, len(tds))

	for i, t := range tds {
		ts[i] = ToLanguage(t)
	}

	return ts
}
