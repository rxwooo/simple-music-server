package handler

import (
	"context"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func AddCollection(ctx context.Context, c *app.RequestContext) {
	uid := string(c.FormValue("userId"))
	rType := string(c.FormValue("type"))
	songId := string(c.FormValue("songId"))
	songListId := string(c.FormValue("songListId"))

	var m_collect model.Collect
	m_collect.UserId, _ = strconv.Atoi(uid)
	m_collect.Type, _ = strconv.Atoi(rType)
	if m_collect.Type == 0 {
		m_collect.SongId, _ = strconv.Atoi(songId)
	} else {
		m_collect.SongListId, _ = strconv.Atoi(songListId)
	}

	m_collect.CreateTime = time.Now()
	res, err := mysql.AddCollection(m_collect)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("collect failed")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("collect successfully", true)
	c.JSON(200, resp)
}

func DeleteCollection(ctx context.Context, c *app.RequestContext) {
	userId := string(c.FormValue("userId"))
	songId := string(c.FormValue("songId"))

	uid, _ := strconv.Atoi(userId)
	sid, _ := strconv.Atoi(songId)

	res, err := mysql.DeleteCollection(uid, sid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("collectiong not exists")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("cancel collection successfully", false)
	c.JSON(200, resp)
}

func CollectiongStatus(ctx context.Context, c *app.RequestContext) {
	userId := string(c.FormValue("userId"))
	songId := string(c.FormValue("songId"))

	uid, _ := strconv.Atoi(userId)
	sid, _ := strconv.Atoi(songId)

	res, err := mysql.CollectiongStatus(uid, sid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	var resp model.ReturnMessage

	if res {
		resp = model.GetSuccessMessage("collected", true)
	} else {
		resp = model.GetSuccessMessage("not collected", false)
	}

	c.JSON(200, resp)
}

func CollectionDetail(ctx context.Context, c *app.RequestContext) {
	userId := string(c.FormValue("userId"))

	uid, _ := strconv.Atoi(userId)

	cols, err := mysql.CollectionDetail(uid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("collection list", cols)
	c.JSON(200, resp)
}
