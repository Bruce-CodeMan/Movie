package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type CommentController struct {
	beego.Controller
}

type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userInfo"`
}

// 获取评论列表
// @router /comment/list [*]
func (this *CommentController) CommentList() {
	// 首先要获取剧集
	episodesId, _ := this.GetInt("episodesId")
	// 获取页码信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if episodesId == 0 {
		this.Data["json"] = ReturnError(4001, "必须要指定剧集")
		this.ServeJSON()
		return
	}
	if limit == 0 {
		limit = 12
	}
	num, comments, err := models.GetCommentList(episodesId, offset, limit)
	if err == nil {
		var data []CommentInfo
		var commentInfo CommentInfo
		for _, v := range comments {
			commentInfo.Id = v.Id
			commentInfo.Content = v.Content
			commentInfo.AddTime = v.AddTime
			commentInfo.AddTimeTitle = DateFormat(v.AddTime)
			commentInfo.UserId = v.UserId
			commentInfo.Stamp = v.Stamp
			commentInfo.PraiseCount = v.PraiseCount

			// 获取用户信息
			commentInfo.UserInfo, _ = models.GetUserInfo(v.UserId)
			data = append(data, commentInfo)
		}
		this.Data["json"] = ReturnSuccess(0, "success", data, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4002, "没有获取到资源")
		this.ServeJSON()
		return
	}
}

// 保存评论
// @router /comment/save [*]
func (this *CommentController) CommentSave() {
	content := this.GetString("content")
	uid, _ := this.GetInt("uid")
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")
	if content == "" {
		this.Data["json"] = ReturnError(4001, "内容不可以为空")
		this.ServeJSON()
		return
	}
	if uid == 0 {
		this.Data["json"] = ReturnError(4002, "请先登录")
		this.ServeJSON()
		return
	}
	if episodesId == 0 {
		this.Data["json"] = ReturnError(4003, "请先指定剧集")
		this.ServeJSON()
		return
	}
	if videoId == 0 {
		this.Data["json"] = ReturnError(4004, "必须指定视频ID")
		this.ServeJSON()
		return
	}
	err := models.SaveComment(content, uid, episodesId, videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", "", 1)
		this.ServeJSON()

	} else {
		this.Data["json"] = ReturnError(5000, err)
		this.ServeJSON()
	}
}
