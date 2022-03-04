package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type TopController struct {
	beego.Controller
}

// 根据频道获取排行榜
// @router /channel/top [*]
func (this *TopController) ChannelTop() {
	// 获取频道ID
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelTop(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4002, "未能获取到资源")
		this.ServeJSON()
	}
}

// 根据频道获取排行榜
// @router /type/top [*]
func (this *VideoController) TypeTop() {
	typeId, _ := this.GetInt("typeId")
	if typeId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定类型")
		this.ServeJSON()
		return
	}
	nums, videos, err := models.GetTypeTop(typeId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, nums)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4002, "未能获取到资源")
		this.ServeJSON()
	}
}
