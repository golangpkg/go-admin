package routers

import (
	"github.com/golangpkg/go-admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	userInfoController := &controllers.UserInfoController{}
	beego.Router("/admin/userInfo/edit", userInfoController, "get:Edit")
	beego.Router("/admin/userInfo/delete", userInfoController, "post:Delete")
	beego.Router("/admin/userInfo/save", userInfoController, "post:Save")
	beego.Router("/admin/userInfo/list", userInfoController, "get:List")
}
