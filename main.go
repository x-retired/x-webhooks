package main

import (
	"github.com/xiexianbin/webhooks/cmd"
	_ "github.com/xiexianbin/webhooks/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cmd.RegisterLogger()

	cmd.RunCommand()
	beego.Run()
}
