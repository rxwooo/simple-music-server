package mysql

import (
	"music-backEnd/biz/model"
)

func AddSong(song model.Song) (int, error) {
	res := DB.Create(&song)
	if res.Error != nil {
		return -1, res.Error
	}

	return song.Id, nil
}

func DeleteSong(sid int) (bool, error) {
	var song *model.Song
	res := DB.Where("id = ?", sid).Delete(&song)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func AllSong() ([]*model.Song, error) {
	songs := make([]*model.Song, 0)
	res := DB.Find(&songs)
	if res.Error != nil {
		return make([]*model.Song, 0), res.Error
	}

	return songs, nil
}

func SongOfId(sid int) ([]*model.Song, error) {
	songs := make([]*model.Song, 0)
	res := DB.Where("id = ?", sid).Find(&songs)
	if res.Error != nil {
		return make([]*model.Song, 0), nil
	}

	return songs, nil
}

func SongLikeSingerName(name string) ([]*model.Song, error) {
	songs := make([]*model.Song, 0)
	singers := make([]*model.Singer, 0)
	res := DB.Where("name like ?", name).Find(&singers)
	if res.Error != nil {
		return make([]*model.Song, 0), res.Error
	}
	singerSp := make([]*int, 0)
	for _, s := range singers {
		singerSp = append(singerSp, &s.Id)
	}

	res = DB.Where("singer_id in ? or name like ?", singerSp, name).Find(&songs)
	if res.Error != nil {
		return make([]*model.Song, 0), res.Error
	}

	return songs, nil
}

func SongOfSingerId(singerid int) ([]*model.Song, error) {
	songs := make([]*model.Song, 0)
	res := DB.Where("singer_id = ?", singerid).Find(&songs)
	if res.Error != nil {
		return make([]*model.Song, 0), res.Error
	}

	return songs, nil
}

func SongUpdate(song model.Song) (bool, error) {
	res := DB.Where("id = ?", song.Id).Updates(&song)
	if res.Error != nil {
		return false, nil
	}

	return res.RowsAffected > 0, nil
}

func SongUpdateCover(songId int, savePath string) error {
	res := DB.Where("id = ?", songId).Updates(model.Song{Pic: savePath})
	return res.Error
}

func GetSongCoverPath(songId int) (string, error) {
	var song *model.Song
	res := DB.Where("id = ?", songId).First(&song)
	return song.Pic, res.Error
}

func SongUpdateUrl(songId int, songPath string) error {
	res := DB.Where("id = ?", songId).Updates(model.Song{Url: songPath})
	return res.Error
}
