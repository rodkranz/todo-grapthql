package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/urfave/cli"

	"github.com/rodkranz/todo-grapthql/cmd"
)

const APP_VER = "v0.0.1"

func init() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "TGG"
	app.Usage = "Api Todo with Go"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdApi,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
