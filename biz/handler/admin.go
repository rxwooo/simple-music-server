package handler

import (
	"context"
	"fmt"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"

	"github.com/cloudwego/hertz/pkg/app"
)

func LoginStatus(ctx context.Context, c *app.RequestContext) {
	username := c.PostForm("name")
	password := c.PostForm("password")
	var err error
	var verifyRes bool

	fmt.Println(username)
	fmt.Println(password)
	verifyRes, err = mysql.LoginVerify(username, password)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !verifyRes {
		resp := model.GetErrorMessage("invalid username or password")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("login success", nil)
	c.JSON(200, resp)
}
