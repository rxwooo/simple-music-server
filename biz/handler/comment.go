package handler

import (
	"context"
	"fmt"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func AddComment(ctx context.Context, c *app.RequestContext) {
	userId := string(c.FormValue("userId"))
	rType := string(c.FormValue("type"))
	songListId := string(c.FormValue("songListId"))
	songId := string(c.FormValue("songId"))
	content := string(c.FormValue("content"))

	fmt.Println(songId)

	var comment model.Comment
	comment.UserId, _ = strconv.Atoi(userId)
	comment.Type, _ = strconv.Atoi(rType)
	if comment.Type == 0 {
		comment.SongId, _ = strconv.Atoi(songId)
	} else {
		comment.SongListId, _ = strconv.Atoi(songListId)
	}
	comment.Content = content
	comment.CreateTime = time.Now()

	res, err := mysql.AddComment(comment)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("comment faild")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("comment successfully", nil)
	c.JSON(200, resp)
}

func DeleteComment(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))

	tp, _ := strconv.Atoi(id)

	res, err := mysql.DeleteComment(tp)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("no such comment")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("delete successfully", nil)
	c.JSON(200, resp)
}

func SongCommentDetail(ctx context.Context, c *app.RequestContext) {
	songId := string(c.FormValue("songId"))

	sid, _ := strconv.Atoi(songId)

	coms, err := mysql.SongCommentDetail(sid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get comments successfully", coms)
	c.JSON(200, resp)
}

func SongListCommentDetail(ctx context.Context, c *app.RequestContext) {
	songListId := string(c.FormValue("songListId"))

	slid, _ := strconv.Atoi(songListId)

	coms, err := mysql.SongListCommentDetail(slid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get comments successfully", coms)
	c.JSON(200, resp)
}

func LikeComment(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))
	up := string(c.FormValue("up"))

	var comment model.Comment
	comment.Id, _ = strconv.Atoi(id)
	comment.Up, _ = strconv.Atoi(up)

	res, err := mysql.UpdateComment(comment)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("like faild")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("liek successfully", nil)
	c.JSON(200, resp)
}
