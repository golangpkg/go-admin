package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
	"time"
)

type UserInfo struct {
	Id         int64  `orm:"auto"`
	UserName   string `orm:"size(255)"`    //登录名
	Password   string `orm:"size(255)"`    //密码
	Name       string `orm:"size(255)"`    //用户名
	BirthDate  string `orm:"size(255)"`    //生日
	Gender     int8       //性别
	Email      string `orm:"size(255)"`    //Email
	Phone      string `orm:"size(255)"`    //电话
	Status     int8     //状态
	CreateTime time.Time //创建时间
	UpdateTime time.Time  //更新时间
}

func init() {
	orm.RegisterModel(new(UserInfo))
}

//创建&更新
func SaveUserInfoById(m *UserInfo) (err error) {
	o := orm.NewOrm()
	var num int64
	if m.Id == 0 {
		m.CreateTime = time.Now()
		m.UpdateTime = time.Now()
		if num, err = o.Insert(m); err == nil {
			logs.Info("Number of records insert in database:", num)
		}
	} else {
		var tmp *UserInfo
		tmp, err = GetUserInfoById(m.Id)

		if err == nil {

			//修改几个参数的名称。
			tmp.UserName = m.UserName
			tmp.Name = m.Name
			tmp.BirthDate = m.BirthDate
			tmp.Gender = m.Gender
			tmp.Email = m.Email
			tmp.Phone = m.Phone
			tmp.Status = m.Status
			tmp.UpdateTime = time.Now()

			if num, err = o.Update(tmp); err == nil {
				logs.Info("Number of records updated in database:", num)
			}
		}
	}
	return
}

//删除
func DeleteUserInfo(id int64) (err error) {
	o := orm.NewOrm()
	v := UserInfo{Id: id}
	if err = o.Read(&v, "Id"); err == nil {
		if num, err := o.Delete(&UserInfo{Id: id}); err == nil {
			logs.Info("Number of records deleted in database:", num)
		}
	}
	return
}

//按id查询
func GetUserInfoById(id int64) (v *UserInfo, err error) {
	o := orm.NewOrm()
	v = &UserInfo{Id: id}
	if err = o.Read(v, "Id"); err == nil {
		return v, nil
	}
	return nil, err
}

//查询数据
func QueryAllUserInfo() (dataList []interface{}, err error) {
	var list []UserInfo
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserInfo))
	//查询
	//查询数据
	if _, err = qs.All(&list); err == nil {
		for _, v := range list {
			dataList = append(dataList, v)
		}
		return dataList, nil
	}
	return nil, err
}
