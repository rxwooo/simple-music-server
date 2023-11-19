package handler

import (
	"context"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func AddListSong(ctx context.Context, c *app.RequestContext) {
	songId := string(c.FormValue("songId"))
	songListId := string(c.FormValue("songListId"))

	var listSong model.ListSong
	listSong.SongId, _ = strconv.Atoi(songId)
	listSong.SongListId, _ = strconv.Atoi(songListId)

	res, err := mysql.AddListSong(listSong)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("songid or songlistid error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("add song to list successfully", nil)
	c.JSON(200, resp)
}

func DeleteListSong(ctx context.Context, c *app.RequestContext) {
	songId := string(c.FormValue("songId"))
	songListId := string(c.FormValue("songListId"))

	sid, _ := strconv.Atoi(songId)
	var ls model.ListSong
	ls.SongId = sid
	ls.SongListId, _ = strconv.Atoi(songListId)
	res, err := mysql.DeleteListSong(ls)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("songId or songListId error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("delete successfully", nil)
	c.JSON(200, resp)
}

func ListSongDetail(ctx context.Context, c *app.RequestContext) {
	songListId := string(c.FormValue("songListId"))

	slid, _ := strconv.Atoi(songListId)
	res, err := mysql.ListSongDetail(slid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}

func ListSongUpdate(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))
	songId := string(c.FormValue("song_id"))
	songListId := string(c.FormValue("song_list_id"))

	var ls model.ListSong
	ls.Id, _ = strconv.Atoi(id)
	ls.SongId, _ = strconv.Atoi(songId)
	ls.SongListId, _ = strconv.Atoi(songListId)

	res, err := mysql.ListSongUpdate(ls)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("song_id or song_list_id error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("update successfully", nil)
	c.JSON(200, resp)
}
