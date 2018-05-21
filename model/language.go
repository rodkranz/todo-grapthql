package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Language struct {
	ID   bson.ObjectId `json:"id"   bson:"_id,omitempty"`
	Name string        `json:"text" bson:"name"`
}

func (l *Language) HasID() bool {
	return l.ID.Valid()
}

func (l *Language) GetID() string {
	return l.ID.Hex()
}

func FindAllLanguages(count int) ([]*Language, error) {
	var results []*Language
	c := Sess.DB(DB_NAME).C(TABLE_LANGUAGE)

	sq := c.Find(bson.M{})
	if count > 0 {
		sq.Limit(count)
	}

	err := sq.All(&results)
	return results, err
}

func FindLanguageById(id string) (*Language, error) {
	result := new(Language)
	if err := IsValidID(id); err != nil {
		return result, err
	}

	c := Sess.DB(DB_NAME).C(TABLE_LANGUAGE)
	err := c.FindId(bson.ObjectIdHex(id)).One(result)

	return result, err
}

func FindLanguageByName(name string) (*Language, error) {
	result := new(Language)

	c := Sess.DB(DB_NAME).C(TABLE_LANGUAGE)
	err := c.Find(bson.M{"name": name}).One(result)

	return result, err
}

func CreateLanguage(newLanguage *Language) error {
	c := Sess.DB(DB_NAME).C(TABLE_LANGUAGE)

	if !newLanguage.HasID() {
		newLanguage.ID = bson.NewObjectId()
	}

	return c.Insert(&newLanguage)
}

func UpdateLanguage(editLanguage *Language) error {
	c := Sess.DB(DB_NAME).C(TABLE_LANGUAGE)

	colQuery := bson.M{"_id": editLanguage.ID}
	change := bson.M{"$set": bson.M{"name": editLanguage.Name}}

	return c.Update(colQuery, change)
}

func DeleteLanguageById(id string) (*Language, error) {
	todo, has := FindLanguageById(id)
	if has != nil {
		return todo, has
	}

	c := Sess.DB(DB_NAME).C(TABLE_LANGUAGE)
	return todo, c.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}
