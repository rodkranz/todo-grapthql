package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Todo struct {
	ID       bson.ObjectId   `json:"id"       bson:"_id,omitempty"`
	Text     string          `json:"text"     bson:"text"`
	Done     bool            `json:"done"     bson:"done"`
	LangId   string                          `bson:"langId"`
	Language *Language       `json:"language" bson:"language"`
}

func (t *Todo) HasID() bool {
	return t.ID.Valid()
}

func (t *Todo) GetID() string {
	return t.ID.Hex()
}

func FindAllTodo(w map[string]interface{}, i int) ([]*Todo, error) {
	var results []*Todo
	c := Sess.DB(DB_NAME).C(TABLE_TODO)

	where := bson.M{}
	for k, v := range w {
		where[k] = v
	}

	sq := c.Find(where)
	if i > 0 {
		sq.Limit(i)
	}

	err := sq.All(&results)

	FindTodoChildren(results)
	return results, err
}

func FindTodoById(id string) (*Todo, error) {
	result := new(Todo)
	if err := IsValidID(id); err != nil {
		return result, err
	}

	c := Sess.DB(DB_NAME).C(TABLE_TODO)
	err := c.FindId(bson.ObjectIdHex(id)).One(result)

	FindTodoChild(result)
	return result, err
}

func CreateTodo(newTodo *Todo) error {
	c := Sess.DB(DB_NAME).C(TABLE_TODO)

	if !newTodo.HasID() {
		newTodo.ID = bson.NewObjectId()
	}

	return c.Insert(&newTodo)
}

func UpdateTodo(editTodo *Todo) error {
	c := Sess.DB(DB_NAME).C(TABLE_TODO)

	colQuery := bson.M{"_id": editTodo.ID}
	change := bson.M{"$set": bson.M{"text": editTodo.Text, "done": editTodo.Done}}

	return c.Update(colQuery, change)
}

func DeleteTodoById(id string) (*Todo, error) {
	todo, has := FindTodoById(id);
	if has != nil {
		return todo, has
	}

	c := Sess.DB(DB_NAME).C(TABLE_TODO)
	return todo, c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

func DeleteTodo(deleteTodo *Todo) (*Todo, error) {
	return DeleteTodoById(deleteTodo.GetID())
}

func FindTodoChildren(todos []*Todo) {
	for _, t := range todos {
		FindTodoChild(t)
	}
}
func FindTodoChild(todo *Todo) {
	if len(todo.LangId) > 0 {
		todo.Language, _ = FindLanguageById(todo.LangId)
	}
}