package handler

import (
	"context"
	"fmt"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func AddSinger(ctx context.Context, c *app.RequestContext) {
	name := string(c.FormValue("name"))
	sex := string(c.FormValue("sex"))
	birth := string(c.FormValue("birth"))
	introduction := string(c.FormValue("introduction"))
	location := string(c.FormValue("location"))
	avator := "defaultPath"

	t, err := time.ParseInLocation("2006-01-02", birth, time.Local)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		fmt.Println(err)
		c.JSON(200, resp)
		return
	}

	var sg model.Singer
	sg.Name = name
	sg.Birth = t
	sg.Sex, _ = strconv.Atoi(sex)
	sg.Introduction = introduction
	sg.Pic = avator
	sg.Location = location

	res, err := mysql.AddSinger(sg)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("singer error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("add singer successfully", nil)
	c.JSON(200, resp)
}

func DeleteSinger(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))

	i, _ := strconv.Atoi(id)
	res, err := mysql.DeleteSinger(i)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("singer id error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("delete successfully", nil)
	c.JSON(200, resp)
}

func AllSinger(ctx context.Context, c *app.RequestContext) {
	res, err := mysql.AllSinger()
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}

func SingerLikeName(ctx context.Context, c *app.RequestContext) {
	name := string(c.FormValue("name"))

	res, err := mysql.SongLikeSingerName("%" + name + "%")
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}

func SingerSex(ctx context.Context, c *app.RequestContext) {
	sex := string(c.FormValue("sex"))
	s, _ := strconv.Atoi(sex)
	res, err := mysql.SingerSex(s)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}

func SingerUpdate(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))
	name := string(c.FormValue("name"))
	sex := string(c.FormValue("sex"))
	birth := string(c.FormValue("birth"))
	introduction := string(c.FormValue("introduction"))
	location := string(c.FormValue("location"))

	t, err := time.ParseInLocation("2006-01-02", birth, time.Local)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	var sg model.Singer
	sg.Id, _ = strconv.Atoi(id)
	sg.Name = name
	sg.Birth = t
	sg.Sex, _ = strconv.Atoi(sex)
	sg.Introduction = introduction
	sg.Location = location
	fmt.Println(sg.Sex)
	res, err := mysql.SingerUpdate(sg)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("singer id error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("update successfully", nil)
	c.JSON(200, resp)
}

func SingerUpdateAvatar(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	singerId, _ := strconv.Atoi(id)

	form, err := c.MultipartForm()
	if err != nil {
		resp := model.GetFatalMessaage("avatar data error")
		c.JSON(200, resp)
		return
	}

	var imgPath string
	files := form.File["file"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)
		imgPath = "./static/singerAvatar/" + id + suffix
		savePath := filepath.Join("./static/singerAvatar", id+suffix)
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			resp := model.GetErrorMessage("avatar data error")
			c.JSON(200, resp)
			return
		}
		res, err := mysql.SingerUpdateAvatar(singerId, savePath)
		if err != nil {
			resp := model.GetErrorMessage("unexcepted error")
			c.JSON(200, resp)
			return
		}

		if !res {
			resp := model.GetErrorMessage("singer id error")
			c.JSON(200, resp)
			return
		}
	}
	fmt.Println(imgPath)
	resp := model.GetSuccessMessage("upload successfully", imgPath)
	c.JSON(200, resp)
}
