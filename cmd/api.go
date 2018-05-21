package cmd

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/urfave/cli"

	"github.com/rodkranz/todo-grapthql/modelql"
)

var (
	HttpAttr = "HttpAttr"
	HttpPort = "8080"
)

var CmdApi = cli.Command{
	Name:        "api",
	Usage:       "Start api graphQL",
	Description: `Start Application api todo with Golang`,
	Action:      runApi,
	Flags: []cli.Flag{
		stringFlag("addr, a", "0.0.0.0", "Address ip of server"),
		stringFlag("port, p", "8080", "Temporary port number to prevent conflict"),
	},
}

func runApi(ctx *cli.Context) error {
	if ctx.IsSet("port") {
		HttpPort = ctx.String("port")
	}
	if ctx.IsSet("addr") {
		HttpAttr = ctx.String("addr")
	}

	h := handler.New(&handler.Config{
		Schema: &modelql.Schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)
	http.Handle("/", http.FileServer(http.Dir("static")))

	fmt.Fprintf(ctx.App.Writer, "List todos:          curl -g 'http://localhost:%v/graphql?query={todoList{id,text,done}}'\n", HttpPort)
	fmt.Fprint(ctx.App.Writer, "List Todo w. filder: curl -g 'http://localhost:8080/graphql?query=mutation+_{todoList(filterDone:false){id,text,done}}}'\n")
	fmt.Fprint(ctx.App.Writer, "Create new Todo:     curl -g 'http://localhost:8080/graphql?query=mutation+_{createTodo(text:\"My+new+todo\"){id,text,done}}'\n")
	fmt.Fprint(ctx.App.Writer, "Update Todo object:  curl -g 'http://localhost:8080/graphql?query=mutation+_{updateTodo(id:\"57bf797df0c011702af7739c\",text:\"My+new+todo\",done:true){id,text,done}}'\n")
	fmt.Fprint(ctx.App.Writer, "Delete Todo object:  curl -g 'http://localhost:8080/graphql?query=mutation+_{deleteTodo(id:\"57bf797df0c011702af7739c\"){id,text,done}}'\n")

	fmt.Printf("Server is running: http://%v:%v\n", HttpAttr, HttpPort)
	err := http.ListenAndServe(":"+HttpPort, nil)
	if err != nil {
		fmt.Fprintf(ctx.App.Writer, "Fail to start server: %v\n", err)
	}

	return nil
}
