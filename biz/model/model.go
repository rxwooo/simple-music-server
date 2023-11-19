package model

import (
	"time"
)

type Admin struct {
	Id       int    `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name     string `gorm:"column:name;type:varchar(45);NOT NULL" json:"name"`
	Password string `gorm:"column:password;type:varchar(45);NOT NULL" json:"password"`
}

func (m *Admin) TableName() string {
	return "admin"
}

type Collect struct {
	Id         int       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserId     int       `gorm:"column:user_id;type:int(10) unsigned;NOT NULL" json:"userId"`
	Type       int       `gorm:"column:type;type:int(10);NOT NULL" json:"type"`
	SongId     int       `gorm:"column:song_id;type:int(10) unsigned" json:"songId"`
	SongListId int       `gorm:"column:song_list_id;type:int(10) unsigned;default:NULL" json:"songListId"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"createTime"`
}

func (m *Collect) TableName() string {
	return "collect"
}

type Comment struct {
	Id         int       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserId     int       `gorm:"column:user_id;type:int(10) unsigned;NOT NULL" json:"userId"`
	SongId     int       `gorm:"column:song_id;type:int(10) unsigned;default:NULL" json:"songId"`
	SongListId int       `gorm:"column:song_list_id;type:int(10) unsigned;default:NULL" json:"songListId"`
	Content    string    `gorm:"column:content;type:varchar(255)" json:"content"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`
	Type       int       `gorm:"column:type;type:int(10);NOT NULL" json:"type"`
	Up         int       `gorm:"column:up;type:int(10) unsigned;default:0;NOT NULL" json:"up"`
}

func (m *Comment) TableName() string {
	return "comment"
}

type Consumer struct {
	Id           int       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Username     string    `gorm:"column:username;type:varchar(255);NOT NULL" json:"username"`
	Password     string    `gorm:"column:password;type:varchar(100);NOT NULL" json:"password"`
	Sex          int       `gorm:"column:sex;type:int(10)" json:"sex"`
	PhoneNum     string    `gorm:"column:phone_num;type:char(15);default: NULL" json:"phone_num"`
	Email        string    `gorm:"column:email;type:char(30);default: NULL" json:"email"`
	Birth        time.Time `gorm:"column:birth;type:datetime" json:"birth"`
	Introduction string    `gorm:"column:introduction;type:varchar(255)" json:"introduction"`
	Location     string    `gorm:"column:location;type:varchar(45)" json:"location"`
	Avator       string    `gorm:"column:avator;type:varchar(255)" json:"avator"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;NOT NULL" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;NOT NULL" json:"update_time"`
}

func (m *Consumer) TableName() string {
	return "consumer"
}

type ListSong struct {
	Id         int `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	SongId     int `gorm:"column:song_id;type:int(10) unsigned;NOT NULL" json:"songId"`
	SongListId int `gorm:"column:song_list_id;type:int(10) unsigned;NOT NULL" json:"songListId"`
}

func (m *ListSong) TableName() string {
	return "list_song"
}

type RankList struct {
	Id         int `gorm:"column:id;type:bigint(20) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	SongListId int `gorm:"column:songListId;type:bigint(20) unsigned;NOT NULL" json:"songListId"`
	ConsumerId int `gorm:"column:consumerId;type:bigint(20) unsigned;NOT NULL" json:"consumerId"`
	Score      int `gorm:"column:score;type:int(10) unsigned;default:0;NOT NULL" json:"score"`
}

func (m *RankList) TableName() string {
	return "rank_list"
}

type Singer struct {
	Id           int       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Name         string    `gorm:"column:name;type:varchar(45);NOT NULL" json:"name"`
	Sex          int       `gorm:"column:sex;type:int(10);default:0" json:"sex"`
	Pic          string    `gorm:"column:pic;type:varchar(255)" json:"pic"`
	Birth        time.Time `gorm:"column:birth;type:datetime" json:"birth"`
	Location     string    `gorm:"column:location;type:varchar(45)" json:"location"`
	Introduction string    `gorm:"column:introduction;type:varchar(255)" json:"introduction"`
}

func (m *Singer) TableName() string {
	return "singer"
}

type Song struct {
	Id           int       `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	SingerId     int       `gorm:"column:singer_id;type:int(10) unsigned;NOT NULL" json:"singer_id"`
	Name         string    `gorm:"column:name;type:varchar(45);NOT NULL" json:"name"`
	Introduction string    `gorm:"column:introduction;type:varchar(255)" json:"introduction"`
	CreateTime   time.Time `gorm:"column:create_time;type:datetime;comment:发行时间;NOT NULL" json:"createTime"`
	UpdateTime   time.Time `gorm:"column:update_time;type:datetime;NOT NULL" json:"update_time"`
	Pic          string    `gorm:"column:pic;type:varchar(255)" json:"pic"`
	Lyric        string    `gorm:"column:lyric;type:text" json:"lyric"`
	Url          string    `gorm:"column:url;type:varchar(255);NOT NULL" json:"url"`
}

func (m *Song) TableName() string {
	return "song"
}

type SongList struct {
	Id           int    `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Title        string `gorm:"column:title;type:varchar(255);NOT NULL" json:"title"`
	Pic          string `gorm:"column:pic;type:varchar(255)" json:"pic"`
	Introduction string `gorm:"column:introduction;type:text" json:"introduction"`
	Style        string `gorm:"column:style;type:varchar(10);default:无" json:"style"`
}

func (m *SongList) TableName() string {
	return "song_list"
}
