package model

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/iris-contrib/errors"
)

const (
	DB_NAME        = "graphql"
	TABLE_LANGUAGE = "language"
	TABLE_TODO     = "todo"
)

var Sess *mgo.Session

func init() {
	var err error
	Sess, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	//defer Sess.Close()
	Sess.SetMode(mgo.Monotonic, true)
	Sess.Ping()
	fmt.Println("MongoDB connected");
}

func IsValidID(id string) (error) {
	if !bson.IsObjectIdHex(id) {
		return errors.New("ID is not a valid value.")
	}
	return nil
}
