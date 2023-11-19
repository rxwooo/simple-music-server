package handler

import (
	"context"
	"fmt"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func AddConsumer(ctx context.Context, c *app.RequestContext) {
	username := string(c.FormValue("username"))
	password := string(c.FormValue("password"))
	sex := string(c.FormValue("sex"))
	phoneNum := string(c.FormValue("phone_num"))
	email := string(c.FormValue("email"))
	birth := string(c.FormValue("birth"))
	introduction := string(c.FormValue("introduction"))
	avator := "./static/consumerAvatar/default.jpg"
	location := string(c.FormValue("location"))

	existence, err := mysql.UserExistence(username)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		fmt.Println(err)
		c.JSON(200, resp)
		return
	}

	if !existence {
		resp := model.GetWarningMessage("username is used")
		c.JSON(200, resp)
		return
	}

	t, err := time.ParseInLocation("2006-01-02", birth, time.Local)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		fmt.Println(err)
		c.JSON(200, resp)
		return
	}

	var user model.Consumer
	user.Username = username
	user.Password = password
	user.Birth = t
	user.Sex, _ = strconv.Atoi(sex)
	user.Introduction = introduction
	user.Avator = avator
	user.Location = location
	if phoneNum != "" {
		user.PhoneNum = phoneNum
	}
	if email != "" {
		user.Email = email
	}
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()

	res, err := mysql.AddConsumer(user)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}
	resp := model.GetSuccessMessage("register successfully", nil)
	c.JSON(200, resp)
}

func UserLoginStatus(ctx context.Context, c *app.RequestContext) {
	username := string(c.FormValue("username"))
	password := string(c.FormValue("password"))

	res, err := mysql.ConsumerVerify(username, password)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("username or password error")
		c.JSON(200, resp)
		return
	}

	users, err := mysql.LoginStatus(username)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("login successfully", users)
	c.JSON(200, resp)
}

func AllUsers(ctx context.Context, c *app.RequestContext) {
	users, err := mysql.AllUsers()
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get users", users)
	c.JSON(200, resp)
}

func UserDetail(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))

	uid, _ := strconv.Atoi(id)
	res, err := mysql.UserDetail(uid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get user", res)
	c.JSON(200, resp)
}

func DeleteUser(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))

	uid, _ := strconv.Atoi(id)
	res, err := mysql.DeleteUser(uid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("delete faild")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("delete successfully", nil)
	c.JSON(200, resp)
}

func UpdateUser(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))
	username := string(c.FormValue("username"))
	sex := string(c.FormValue("sex"))
	phoneNum := string(c.FormValue("phone_num"))
	email := string(c.FormValue("email"))
	birth := string(c.FormValue("birth"))
	introduction := string(c.FormValue("introduction"))
	location := string(c.FormValue("location"))

	t, err := time.ParseInLocation("2006-01-02", birth, time.Local)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	var user model.Consumer
	user.Id, _ = strconv.Atoi(id)
	user.Username = username
	user.Birth = t
	user.Sex, _ = strconv.Atoi(sex)
	user.Introduction = introduction
	user.Location = location
	if phoneNum != "" {
		user.PhoneNum = phoneNum
	}
	if email != "" {
		user.Email = email
	}
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()

	res, err := mysql.UpdateUser(user)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		fmt.Println(err)
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("update faild")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("update successfully", nil)
	c.JSON(200, resp)
}

func UserUpdatePasswd(ctx context.Context, c *app.RequestContext) {
	username := string(c.FormValue("username"))
	oldPasswd := string(c.FormValue("old_password"))
	password := string(c.FormValue("password"))

	res, err := mysql.UserVerifyPassword(username, oldPasswd)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("password error")
		c.JSON(200, resp)
		return
	}

	res, err = mysql.UserUpdatePasswd(username, password)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	if !res {
		resp := model.GetErrorMessage("password error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("password update successfully", nil)
	c.JSON(200, resp)
}

func UserUpdateAvatar(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	userId, _ := strconv.Atoi(id)

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
		imgPath = "./static/consumerAvatar/" + id + suffix
		savePath := filepath.Join("./static/consumerAvatar", id+suffix)

		oriPath, err := mysql.GetUserCover(userId)
		if err != nil {
			resp := model.GetErrorMessage("user id error")
			c.JSON(200, resp)
			return
		}

		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			resp := model.GetErrorMessage("avatar data error")
			c.JSON(200, resp)
			return
		}
		res, err := mysql.UserUpdateAvatar(userId, savePath)
		if err != nil {
			resp := model.GetErrorMessage("unexcepted error")
			c.JSON(200, resp)
			return
		}

		if !res {
			resp := model.GetErrorMessage("user id error")
			c.JSON(200, resp)
			return
		}

		if oriPath != "./static/consumerAvatar/default.jpg" {
			os.RemoveAll(savePath)
		}
	}
	fmt.Println(imgPath)
	resp := model.GetSuccessMessage("upload successfully", imgPath)
	c.JSON(200, resp)
}

