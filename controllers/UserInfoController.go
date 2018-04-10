package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego"
	"github.com/freewebsys/go-admin/models"
)

type UserInfoController struct {
	beego.Controller
}

//修改
func (c *UserInfoController) Edit() {
	//获得id
	id, _ := c.GetInt64("Id", 0)
	userInfo, err := models.GetUserInfoById(id)
	if err == nil {
		c.Data["UserInfo"] = userInfo
	} else {
		tmpUserInfo := &models.UserInfo{}
		tmpUserInfo.Status = -1
		tmpUserInfo.Gender = -1
		c.Data["UserInfo"] = tmpUserInfo
	}
	c.TplName = "userInfo/edit.html"
}

//删除
func (c *UserInfoController) Delete() {
	//获得id
	id, _ := c.GetInt64("Id", 0)
	if err := models.DeleteUserInfo(id); err == nil {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

//保存
func (c *UserInfoController) Save() {
	//自动解析绑定到对象中
	userInfo := models.UserInfo{}
	if err := c.ParseForm(&userInfo); err == nil {
		if err := models.SaveUserInfoById(&userInfo); err == nil {
			c.Data["json"] = ""
		} else {
			c.Data["json"] = "error"
		}
	} else {
		c.Data["json"] = "error"
	}
	c.ServeJSON()
}

//返回全部数据
func (c *UserInfoController) List() {

	dataList, err := models.QueryAllUserInfo()
	if err == nil {
		c.Data["List"] = dataList
	}
	logs.Info("dataList :", dataList)
	c.TplName = "userInfo/list.html"

}
