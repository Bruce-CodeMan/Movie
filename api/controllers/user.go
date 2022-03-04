package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
	"regexp"
	"strconv"
	"strings"
)

type UserController struct {
	beego.Controller
}

// 用户注册功能
// @router /register/save [post]
func (this *UserController) RegisterSave() {
	var (
		mobile   string
		password string
	)
	mobile = this.GetString("mobile")
	password = this.GetString("password")

	// 判断手机号是否为空
	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不可以为空")
		this.ServeJSON()
		return
	}

	// 判断手机号是否符合规范
	IsMobile, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}$`, mobile)
	if !IsMobile {
		this.Data["json"] = ReturnError(4002, "手机号不符合规范")
		this.ServeJSON()
		return
	}

	// 判断密码是否为空
	if password == "" {
		this.Data["json"] = ReturnError(4003, "密码不可以为空")
		this.ServeJSON()
		return
	}

	// 判断手机号是否已经注册
	status := models.IsUserMobile(mobile)
	if status {
		this.Data["json"] = ReturnError(4005, "手机号已经注册")
		this.ServeJSON()
		return
	} else {
		err := models.SaveUser(mobile, MD5V(password))
		if err == nil {
			this.Data["json"] = ReturnSuccess(0, "手机号注册成功", nil, 0)
			this.ServeJSON()
		} else {
			this.Data["json"] = ReturnError(5000, err)
			this.ServeJSON()
		}
	}
}

// 用户登录
// @router /login/do [*]
func (this *UserController) LoginDo() {
	var (
		mobile   string
		password string
	)

	mobile = this.GetString("mobile")
	password = this.GetString("password")

	// 判断手机号是否为空
	if mobile == "" {
		this.Data["json"] = ReturnError(4001, "手机号不可以为空")
		this.ServeJSON()
		return
	}
	// 判断手机号是否符合规范
	IsMobile, _ := regexp.MatchString(`^1(3|4|5|7|8)[0-9]\d{8}`, mobile)
	if !IsMobile {
		this.Data["json"] = ReturnError(4002, "手机号格式不正确")
		this.ServeJSON()
		return
	}
	// 判断密码是否为空
	if password == "" {
		this.Data["json"] = ReturnError(4003, "密码不可以为空")
		this.ServeJSON()
		return
	}
	// 判断手机号是否存在
	uid, username := models.IsMobileLogin(mobile, MD5V(password))
	if uid != 0 {
		this.Data["json"] = ReturnSuccess(0, "登录成功", map[string]interface{}{"uid": uid, "username": username}, 1)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = ReturnError(4004, "手机号或者密码错误")
		this.ServeJSON()
	}
}

// 批量发送消息
// @router /send/message [*]
func (this *UserController) SendMessageDo() {
	uids := this.GetString("uids")
	content := this.GetString("content")

	if uids == "" {
		this.Data["json"] = ReturnError(4001, "请填写接收人")
		this.ServeJSON()
		return
	}
	if content == "" {
		this.Data["json"] = ReturnError(4002, "请填写内容")
		this.ServeJSON()
		return
	}
	messageId, err := models.SendMessageDo(content)
	if err == nil {
		// 分割
		uidConfig := strings.Split(uids, ",")
		for _, value := range uidConfig {
			userId, _ := strconv.Atoi(value)
			err := models.SendMessageUser(userId, messageId)
			if err == nil {
				this.Data["json"] = ReturnSuccess(0, "success", "", 0)
				this.ServeJSON()
				return
			} else {
				this.Data["json"] = ReturnError(4004, "发送消息失败")
				this.ServeJSON()
				return
			}
		}
	} else {
		this.Data["json"] = ReturnError(4003, "发送消息失败，请联系管理员")
		this.ServeJSON()
		return
	}
}