package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id       int
	Name     string
	Password string
	Status   int
	AddTime  int64
	Mobile   string
	Avatar   string
}

type UserInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AddTime int64  `json:"addTime"`
	Avatar  string `json:"avatar"`
}

func init() {
	orm.RegisterModel(new(User))
}

// 根据手机号去判断用户是否存在
func IsUserMobile(mobile string) bool {
	var (
		user User
		err  error
	)
	o := orm.NewOrm()
	user = User{Mobile: mobile}
	err = o.Read(&user, "Mobile")
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	} else {
		return true
	}
}

// 保存用户
func SaveUser(mobile string, password string) error {
	var user User
	o := orm.NewOrm()
	user.Name = ""
	user.Mobile = mobile
	user.Password = password
	user.Status = 1
	user.AddTime = time.Now().Unix()
	_, err := o.Insert(&user)
	return err
}

// 登录功能
func IsMobileLogin(mobile string, password string) (int, string) {
	var user User
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("mobile", mobile).Filter("password", password).One(&user)
	if err == orm.ErrNoRows {
		return 0, ""
	} else if err == orm.ErrMissPK {
		return 0, ""
	} else {
		return user.Id, user.Name
	}
}

// 根据用户Id获取用户信息
func GetUserInfo(uid int) (UserInfo, error) {
	o := orm.NewOrm()
	var user UserInfo
	err := o.Raw("SELECT id,name,add_time,avatar FROM user WHERE id=? LIMIT 1", uid).QueryRow(&user)
	return user, err
}
