package handler

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"music-backEnd/biz/dal/mysql"
	"music-backEnd/biz/model"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"

	//"io"

	"github.com/bogem/id3v2"
	"github.com/cloudwego/hertz/pkg/app"
)

func serveFrames(imgByte []byte, savePath string) {

	img, _, err := image.Decode(bytes.NewReader(imgByte))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create(savePath)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 100

	err = jpeg.Encode(out, img, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println(err)
	}

}

func getVideoName() string {
	t := time.Unix(int64(time.Now().Unix()), 0)
	timeStr := t.Format("2006_01_02_15_04_05")
	videoName := timeStr
	return videoName
}

func AddSong(ctx context.Context, c *app.RequestContext) {
	singerId := string(c.FormValue("singerId"))
	name := string(c.FormValue("name"))
	introduction := string(c.FormValue("introduction"))
	lyric := string(c.FormValue("lyric"))
	pic := "./static/songCover/default.jpg"

	var song model.Song
	song.SingerId, _ = strconv.Atoi(singerId)
	song.Name = name
	song.Introduction = introduction
	song.Lyric = lyric
	song.Pic = pic

	form, err := c.MultipartForm()
	if err != nil {
		resp := model.GetFatalMessaage("avatar data error")
		c.JSON(200, resp)
		return
	}

	var songPath string
	files := form.File["file"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)
		songName := getVideoName()
		songPath = "./static/song/" + songName + suffix
		savePath := filepath.Join("./static/song", songName+suffix)
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			resp := model.GetErrorMessage("song file data error")
			c.JSON(200, resp)
			return
		}
		song.Url = songPath

		res, err := mysql.AddSong(song)
		if err != nil {
			resp := model.GetErrorMessage("unexcepted error")
			c.JSON(200, resp)
			return
		}

		if res == -1 {
			resp := model.GetErrorMessage("singer id error")
			c.JSON(200, resp)
			return
		}

		tag, err := id3v2.Open(songPath, id3v2.Options{Parse: true})
		if err != nil {
			log.Fatal("Error while opening mp3 file: ", err)
		} else {
			pictures := tag.GetFrames(tag.CommonID("Attached picture"))
			f := pictures[0]
			pic, ok := f.(id3v2.PictureFrame)
			if !ok {
				log.Fatal("Couldn't assert picture frame")
			} else {
				serveFrames(pic.Picture, "./static/songCover/"+strconv.Itoa(res)+".jpg")
				mysql.SongUpdateCover(res, "./static/songCover/"+strconv.Itoa(res)+".jpg")
			}
		}
		defer tag.Close()
	}
	fmt.Println(songPath)
	resp := model.GetSuccessMessage("upload successfully", songPath)
	c.JSON(200, resp)
}

func DeleteSong(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))

	i, _ := strconv.Atoi(id)
	res, err := mysql.DeleteSong(i)
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

func AllSong(ctx context.Context, c *app.RequestContext) {
	res, err := mysql.AllSong()
	if err != nil {
		rsep := model.GetErrorMessage("unexcepted error")
		c.JSON(200, rsep)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}

func SongDetail(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))

	sid, _ := strconv.Atoi(id)
	res, err := mysql.SongOfId(sid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}

func SongOfSingerId(ctx context.Context, c *app.RequestContext) {
	singerId := string(c.FormValue("singerId"))
	sid, _ := strconv.Atoi(singerId)

	res, err := mysql.SongOfSingerId(sid)
	if err != nil {
		resp := model.GetErrorMessage("unexcepted error")
		c.JSON(200, resp)
		return
	}

	resp := model.GetSuccessMessage("get successfully", res)
	c.JSON(200, resp)
}

func SongLikeSingerName(ctx context.Context, c *app.RequestContext) {
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

func SongUpdate(ctx context.Context, c *app.RequestContext) {
	id := string(c.FormValue("id"))
	singerId := string(c.FormValue("singerId"))
	name := string(c.FormValue("name"))
	introduction := string(c.FormValue("introduction"))
	lyric := string(c.FormValue("lyric"))

	var song model.Song
	song.Id, _ = strconv.Atoi(id)
	song.SingerId, _ = strconv.Atoi(singerId)
	song.Name = name
	song.Introduction = introduction
	song.Lyric = lyric

	res, err := mysql.SongUpdate(song)
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

func GetSongFile(ctx context.Context, c *app.RequestContext) {
	//参数解析
	name := c.Param("name")
	path := filepath.Join("./static/song", name)
	file, err := os.Open(path)
	if err != nil {
		c.String(500, "Error opening file")
		return
	}
	defer file.Close()
	rangeRaw := c.Request.Header.Get("Range")
	content, _ := ioutil.ReadAll(file)
	if rangeRaw == "" {
		c.Status(200)
		c.Header("Content-Type", "audio/mpeg")
		c.Header("Content-Length", fmt.Sprintf("%d", len(content)))
		c.File(path)
		return
	}
	//解析Header中的Range信息
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	res := strings.FieldsFunc(string(c.Request.Header.PeekRange()), f)
	startPos, _ := strconv.Atoi(res[1])
	if startPos > len(content) {
		c.String(416, "Requested Range Not Satisfiable")
		return
	}
	//写入Header信息
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Content-Length", fmt.Sprintf("%d", len(content)-startPos))
	c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", startPos, len(content)-1, len(content)))
	c.Status(206)
	//写入分割后的文件
	c.Write(content[startPos:])
}

func UpdateSongCover(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	songId, _ := strconv.Atoi(id)

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
		imgPath = "./static/songCover/" + id + suffix
		savePath := filepath.Join("./static/songCover", id+suffix)

		oriPath, err := mysql.GetSongCoverPath(songId)
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

		err = mysql.SongUpdateCover(songId, savePath)
		if err != nil {
			resp := model.GetErrorMessage("unexcepted error")
			c.JSON(200, resp)
			return
		}

		if oriPath != "./static/songCover/default.jpg" {
			os.RemoveAll(oriPath)
		}
	}
	fmt.Println(imgPath)
	resp := model.GetSuccessMessage("upload successfully", imgPath)
	c.JSON(200, resp)
}

func SongUrlUpdate(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	songId, _ := strconv.Atoi(id)

	form, err := c.MultipartForm()
	if err != nil {
		resp := model.GetFatalMessaage("song data error")
		c.JSON(200, resp)
		return
	}

	var songPath string
	files := form.File["file"]
	for _, file := range files {
		suffix := filepath.Ext(file.Filename)
		songName := getVideoName()
		songPath = "./static/song/" + songName + suffix
		savePath := filepath.Join("./static/song", songName+suffix)
		err = c.SaveUploadedFile(file, savePath)
		if err != nil {
			resp := model.GetErrorMessage("song file data error")
			c.JSON(200, resp)
			return
		}

		err = mysql.SongUpdateUrl(songId, songPath)
		if err != nil {
			resp := model.GetErrorMessage("song file data error")
			c.JSON(200, resp)
			return
		}
	}
	resp := model.GetSuccessMessage("upload successfully", songPath)
	c.JSON(200, resp)
}
