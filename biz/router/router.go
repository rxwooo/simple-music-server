package router

import (
	"music-backEnd/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func Register(r *server.Hertz) {
	r.Static("/static", "./")
	r.GET("/static/song/:name", handler.GetSongFile)
	r.POST("admin/login/status", handler.LoginStatus)

	r.POST("collection/add", handler.AddCollection)
	r.DELETE("collection/delete", handler.DeleteCollection)
	r.POST("collection/status", handler.CollectiongStatus)
	r.GET("collection/detail", handler.CollectionDetail)

	r.POST("/comment/add", handler.AddComment)
	r.GET("/comment/delete", handler.DeleteComment)
	r.GET("/comment/song/detail", handler.SongCommentDetail)
	r.GET("/comment/songList/detail", handler.SongListCommentDetail)
	r.POST("/comment/like", handler.LikeComment)

	r.POST("/user/add", handler.AddConsumer)
	r.POST("/user/login/status", handler.UserLoginStatus)
	r.GET("/user", handler.AllUsers)
	r.GET("/user/detail", handler.UserDetail)
	r.GET("/user/delete", handler.DeleteUser)
	r.GET("/user/update", handler.UpdateUser)
	r.POST("/user/updatePassword", handler.UserUpdatePasswd)
	r.POST("/user/avatar/update", handler.UserUpdateAvatar)

	r.POST("/listSong/add", handler.AddListSong)
	r.GET("/listSong/delete", handler.DeleteListSong)
	r.GET("/listSong/detail", handler.ListSongDetail)
	r.POST("/listSong/update", handler.ListSongUpdate)

	r.POST("/rankList/add", handler.AddRankList)
	r.GET("/rankList", handler.RankListOfSongList)
	r.GET("/rankList/user", handler.RankListOfUserSongList)

	r.POST("/singer/add", handler.AddSinger)
	r.GET("/singer/delete", handler.DeleteSinger)
	r.GET("/singer", handler.AllSinger)
	r.GET("/singer/name/detail", handler.SingerLikeName)
	r.GET("/singer/sex/detail", handler.SingerSex)
	r.POST("/singer/update", handler.SingerUpdate)
	r.POST("/singer/avatar/update", handler.SingerUpdateAvatar)

	r.POST("/song/add", handler.AddSong)
	r.GET("/song/delete", handler.DeleteSong)
	r.GET("/song", handler.AllSong)
	r.GET("/song/detail", handler.SongDetail)
	r.GET("/song/singer/detail", handler.SongOfSingerId)
	r.GET("/song/singerName/detail", handler.SingerLikeName)
	r.POST("/song/update", handler.SongUpdate)
	r.POST("/song/img/update", handler.UpdateSongCover)
	r.POST("song/url/update", handler.SongUrlUpdate)

	r.POST("/songList/add", handler.AddSongList)
	r.GET("/songList/delete", handler.DeleteSongList)
	r.GET("/songList", handler.AllSongList)
	r.GET("/songList/likeTitle/detail", handler.SongListLikeTitle)
	r.GET("/songList/style/detail", handler.SongListLikeStyle)
	r.POST("/songList/update", handler.SongListUpdate)
	r.POST("/songList/img/update", handler.SongListUpdateCover)

}
