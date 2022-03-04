package models

import (
	"github.com/astaxie/beego/orm"
)

type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	EpisodesUpdateTime int64
	Comment            int
	UserId             int
	IsRecommend        int
}

type VideoData struct {
	Id            int
	Title         string
	AddTime       string
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
}

type Episodes struct {
	Id      int
	Title   string
	AddTime int64
	Num     int
	PlayUrl string
	Comment int
}

func init() {
	orm.RegisterModel(new(Video))
}

// 获取正在热播
func GetHotList(isHot int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time, img,img1,episodes_count,is_end FROM video WHERE status=1 AND is_hot=? ORDER BY hot_sort DESC LIMIT 9", isHot).QueryRows(&videos)
	return num, videos, err
}

// 获取电视剧推荐
func GetEpisodesRecommend(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time, img,img1,episodes_count,is_end FROM video WHERE status=1 AND channel_id=? AND is_recommend=1 ORDER BY episode_sort DESC LIMIT 9", channelId).QueryRows(&videos)
	return num, videos, err
}

// 获取电视剧推荐
func GetMoviesRecommend(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id,title,sub_title,add_time, img,img1,episodes_count,is_end FROM video WHERE status=1 AND channel_id=? AND is_recommend=1 ORDER BY movie_sort DESC LIMIT 9", channelId).QueryRows(&videos)
	return num, videos, err
}

// 获取推荐
func GetChannelRecommendRegionList(channelId int, regionId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id, title, sub_title, add_time, img, img1, episodes_count, is_end FROM "+
		"video WHERE status=1 AND is_recommend=1 AND region_id=? AND channel_id=? ORDER BY "+
		"episodes_update_time DESC LIMIT 9", regionId, channelId).QueryRows(&videos)
	return num, videos, err
}

// 类型推荐
func GetChannelRecommendTypeList(channelId int, typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id, title, sub_title, add_time, img, img1, episodes_count, is_end FROM "+
		"video WHERE status=1 AND type_id=? AND is_recommend=1 AND channel_id=? ORDER BY episodes_update_time DESC LIMIT 9", typeId, channelId).QueryRows(&videos)
	return num, videos, err
}

// 根据传入的参数来获取所有的数据
func GetChannelVideoList(channelId int, regionId int, typeId int, end string, sort string, offset int, limit int) (int64, []orm.Params, error) {
	o := orm.NewOrm()
	var videos []orm.Params
	qs := o.QueryTable("video")
	qs = qs.Filter("channel_id", channelId)
	qs = qs.Filter("status", 1)
	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}
	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}
	if end == "n" {
		qs = qs.Filter("is_end", 0)
	} else {
		qs = qs.Filter("is_end", 1)
	}
	if sort == "episodesUpdateTime" {
		qs = qs.OrderBy("-episodes_update_time")
	} else if sort == "comment" {
		qs = qs.OrderBy("-comment")
	} else if sort == "addTime" {
		qs = qs.OrderBy("-add_time")
	} else {
		qs = qs.OrderBy("-add_time")
	}
	nums, _ := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	qs = qs.Limit(limit, offset)
	_, err := qs.Values(&videos, "id", "title", "sub_title", "add_time", "img", "img1", "episodes_count", "is_end")
	return nums, videos, err
}

// 获取视频的详情
func GetVideoInfo(videoId int) (Video, error) {
	o := orm.NewOrm()
	var video Video
	err := o.Raw("SELECT * FROM video WHERE id=? LIMIT 1", videoId).QueryRow(&video)
	return video, err
}

// 获取剧集的列表
func GetVideoEpisodesList(videoId int) (int64, []Episodes, error) {
	o := orm.NewOrm()
	var episodes []Episodes
	num, err := o.Raw("SELECT id, title, add_time, num, play_url, comment FROM video_episodes WHERE video_id=? AND status=1 ORDER BY num ASC", videoId).QueryRows(&episodes)
	return num, episodes, err
}

// 根据频道获取排行榜
func GetChannelTop(channelId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id, title, sub_title,img,img1, add_time,episodes_count,is_end FROM video WHERE status=1 AND channel_id=? ORDER BY comment DESC LIMIT 10", channelId).QueryRows(&videos)
	return num, videos, err
}

// 根据类型获取排行榜
func GetTypeTop(typeId int) (int64, []VideoData, error) {
	o := orm.NewOrm()
	var videos []VideoData
	num, err := o.Raw("SELECT id, title, sub_title,img,img1, add_time,episodes_count,is_end FROM video WHERE status=1 AND type_id=? ORDER BY comment DESC LIMIT 10", typeId).QueryRows(&videos)
	return num, videos, err
}
