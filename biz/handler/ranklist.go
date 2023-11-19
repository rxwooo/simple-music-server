package handler

import (
	"context"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

func AddRankList(ctx context.Context, c *app.RequestContext) {
	songListId := string(c.FormValue("songListId"))
	consumerid := string(c.FormValue("consumerId"))
	score := string(c.FormValue("score"))

	var rl model.RankList
	rl.SongListId, _ = strconv.Atoi(songListId)
	rl.ConsumerId, _ = strconv.Atoi(consumerid)
	rl.Score, _ = strconv.Atoi(score)

	res, err := mysql.AddRankList(rl)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("songlistid or consumerid error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("comment successfully", nil)
	c.JSON(200, resp)
}

func RankListOfSongList(ctx context.Context, c *app.RequestContext) {
	songListId := string(c.FormValue("songListId"))

	slid, _ := strconv.Atoi(songListId)
	res, err := mysql.RankListOfSongList(slid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}

func RankListOfUserSongList(ctx context.Context, c *app.RequestContext) {
	consumerId := string(c.FormValue("consumerId"))
	songListid := string(c.FormValue("songListI"))

	cid, _ := strconv.Atoi(consumerId)
	slid, _ := strconv.Atoi(songListid)

	var rl model.RankList
	rl.ConsumerId = cid
	rl.SongListId = slid

	res, err := mysql.RankListOfUserSongList(rl)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}
