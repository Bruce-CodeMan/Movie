package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego"
)

type VideoController struct {
	beego.Controller
}

// 频道顶部获取广告
// @router /channel/advert [*]
func (this *VideoController) ChannelAdvert() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须制定频道")
		this.ServeJSON()
		return
	}
	num, adverts, err := models.GetChannelAdvert(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", adverts, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "请求数据失败")
		this.ServeJSON()
	}
}

// 频道页-获取正在热播
// @router /channel/hot [*]
func (this *VideoController) ChannelHotList() {
	isHot, _ := this.GetInt("isHot")
	if isHot == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定热门视频")
		this.ServeJSON()
		return
	}
	num, videos, err := models.GetHotList(isHot)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关的内容")
		this.ServeJSON()
	}
}

// 频道页-获取电视剧推荐
// @router /episode/recommend [*]
func (this *VideoController) EpisodesRecommendList() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		return
	}
	num, videos, err := models.GetEpisodesRecommend(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4003, "没有相关的内容")
		this.ServeJSON()
	}
}

// 频道页-获取电影推荐
// @router /movie/recommend [*]
func (this *VideoController) MovieRecommendList() {
	channelId, _ := this.GetInt("channelId")
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		return
	}
	num, videos, err := models.GetMoviesRecommend(channelId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4003, "没有相关的内容")
		this.ServeJSON()
	}
}

// 频道页-根据频道地区获取推荐的视频
// @router /channel/recommend/region [*]
func (this *VideoController) ChannelRecommendRegionList() {
	channelId, _ := this.GetInt("channelId")
	regionId, _ := this.GetInt("regionId")

	// 开始判断
	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		return
	}
	if regionId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定频道地区")
		this.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelRecommendRegionList(channelId, regionId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4003, "没有相关的内容")
		this.ServeJSON()
	}
}

// 频道页-根据频道页获取推荐视频
// @router /channel/recommend/type [*]
func (this *VideoController) GetChannelRecommendTypeList() {
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		return
	}
	if typeId == 0 {
		this.Data["json"] = ReturnError(4002, "必须指定类型")
		this.ServeJSON()
		return
	}
	num, videos, err := models.GetChannelRecommendTypeList(channelId, typeId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4003, "没有相关的资源")
		this.ServeJSON()
	}
}

// 根据传入的参数获取视频列表
// @router /channel/video [*]
func (this *VideoController) ChannelVideo() {
	// 获取频道ID
	channelId, _ := this.GetInt("channelId")
	// 获取频道地区ID
	regionId, _ := this.GetInt("regionId")
	// 获取频道类型ID
	typeId, _ := this.GetInt("typeId")
	// 获取状态
	end := this.GetString("end")
	// 获取排序
	sort := this.GetString("sort")
	// 获取页码信息
	limit, _ := this.GetInt("limit")
	offset, _ := this.GetInt("offset")

	if channelId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定频道")
		this.ServeJSON()
		return
	}

	if limit == 0 {
		limit = 12
	}
	num, videos, err := models.GetChannelVideoList(channelId, regionId, typeId, end, sort, offset, limit)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", videos, num)
		this.ServeJSON()
	} else {
		this.Data["json"] = ReturnError(4004, "没有相关的内容")
		this.ServeJSON()
	}
}

// 获取视频的详情
// @router /video/info [*]
func (this *VideoController) VideoInfo() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定视频")
		this.ServeJSON()
		return
	}
	video, err := models.GetVideoInfo(videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", video, 1)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = ReturnError(4002, err)
		this.ServeJSON()
		return
	}
}

// 获取剧集列表
// @router /video/episodes/list [*]
func (this *VideoController) VideoEpisodesList() {
	videoId, _ := this.GetInt("videoId")
	if videoId == 0 {
		this.Data["json"] = ReturnError(4001, "必须指定剧集")
		this.ServeJSON()
		return
	}
	num, episodes, err := models.GetVideoEpisodesList(videoId)
	if err == nil {
		this.Data["json"] = ReturnSuccess(0, "success", episodes, num)
		this.ServeJSON()
		return
	} else {
		this.Data["json"] = ReturnError(4002, "请求数据失败")
		this.ServeJSON()
		return
	}
}
