package handler

import (
	"context"
	"fmt"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func AddSongList(ctx context.Context, c *app.RequestContext) {
	title := string(c.FormValue("title"))
	introduction := string(c.FormValue("introductioin"))
	style := string(c.FormValue("style"))
	pic := "./static/songListCover/default.jpg"

	var songList model.SongList
	songList.Title = title
	songList.Introduction = introduction
	songList.Style = style
	songList.Pic = pic

	res, err := mysql.AddSongList(songList)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("imformation error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("add songlist successfully", nil)
	c.JSON(200, resp)
}

func DeleteSongList(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))

	slid, _ := strconv.Atoi(id)
	res, err := mysql.DeleteSongList(slid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("id error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("delete successfully", nil)
	c.JSON(200, resp)
}

func AllSongList(ctx context.Context, c *app.RequestContext) {
	sls, err := mysql.AllSongList()
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", sls)
	c.JSON(200, resp)
}

func SongListLikeTitle(ctx context.Context, c *app.RequestContext) {
	title := string(c.FormValue("title"))

	sls, err := mysql.SongListLikeTitle("%" + title + "%")
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", sls)
	c.JSON(200, resp)
}

func SongListLikeStyle(ctx context.Context, c *app.RequestContext) {
	style := string(c.FormValue("style"))

	sls, err := mysql.SongListLikeStyle("%" + style + "%")
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", sls)
	c.JSON(200, resp)
}

func SongListUpdate(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))
	title := string(c.FormValue("title"))
	introduction := string(c.FormValue("introduction"))
	style := string(c.FormValue("style"))

	var sl model.SongList
	sl.Title = title
	sl.Introduction = introduction
	sl.Style = style
	sl.Id, _ = strconv.Atoi(id)

	res, err := mysql.SongListUpdate(sl)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("id error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("update successfully", nil)
	c.JSON(200, resp)
}

func SongListUpdateCover(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	songListId, _ := strconv.Atoi(id)

	form, err := c.MultipartForm()
	if err != nil {
		resp := model.GetFatalMessaage("image data error")
		c.JSON(200, resp)
		return
	}

	var imgPath string
	files := form.File["file"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)
		imgPath = "./static/songListCover/" + id + suffix
		savePath := filepath.Join("./static/songListCover", id+suffix)

		oriPath, err := mysql.GetSongListCoverPath(songListId)
		if err != nil {
			resp := model.GetErrorMessage("song id error")
			c.JSON(200, resp)
			return
		}

		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			resp := model.GetErrorMessage("image data error")
			c.JSON(200, resp)
			return
		}

		err = mysql.SongListUpdateCover(songListId, savePath)
		if err != nil {
			resp := model.GetErrorMessage("unexcepted error")
			c.JSON(200, resp)
			return
		}

		if oriPath != "./static/songListCover/default.jpg" {
			os.RemoveAll(oriPath)
		}
	}
	fmt.Println(imgPath)
	resp := model.GetSuccessMessage("upload successfully", imgPath)
	c.JSON(200, resp)
}
