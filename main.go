package main

import (
	_ "github.com/golangpkg/go-admin/routers"
	"github.com/astaxie/beego"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/context"
)

func main() {

	//注册sqlite3
	orm.RegisterDataBase("default", "sqlite3", "go-admin.db")
	//同步 ORM 对象和数据库
	//这时, 在你重启应用的时候, beego 便会自动帮你创建数据库表。
	orm.Debug = true

	orm.RunSyncdb("default", false, true)

	//增加拦截器。
	var filterAdmin = func(ctx *context.Context) {
		url := ctx.Input.URL()
		logs.Info("##### filter url : %s", url)
		//TODO 如果判断用户未登录。

	}
	beego.InsertFilter("/admin/*", beego.BeforeExec, filterAdmin)

	beego.Run()
}
