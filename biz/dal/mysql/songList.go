package mysql

import (
	"music-backEnd/biz/model"
)

func AddSongList(sl model.SongList) (bool, error) {
	res := DB.Create(&sl)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func DeleteSongList(slid int) (bool, error) {
	var sl *model.SongList
	res := DB.Where("id = ?", slid).Delete(&sl)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func AllSongList() ([]*model.SongList, error) {
	sls := make([]*model.SongList, 0)
	res := DB.Find(&sls)
	if res.Error != nil {
		return make([]*model.SongList, 0), res.Error
	}

	return sls, nil
}

func SongListLikeTitle(like string) ([]*model.SongList, error) {
	sls := make([]*model.SongList, 0)
	res := DB.Where("title like ?", like).Find(&sls)
	if res.Error != nil {
		return make([]*model.SongList, 0), res.Error
	}

	return sls, nil
}

func SongListLikeStyle(style string) ([]*model.SongList, error) {
	sls := make([]*model.SongList, 0)
	res := DB.Where("style like ?", style).Find(&sls)
	if res.Error != nil {
		return make([]*model.SongList, 0), res.Error
	}

	return sls, nil
}

func SongListUpdate(sl model.SongList) (bool, error) {
	res := DB.Where("id = ?", sl.Id).Updates(&sl)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func GetSongListCoverPath(slid int) (string, error) {
	var songList *model.SongList
	res := DB.Where("id = ?", slid).First(&songList)
	return songList.Pic, res.Error
}

func SongListUpdateCover(slid int, savePath string) error {
	res := DB.Where("id = ?", slid).Updates(model.SongList{Pic: savePath})
	return res.Error
}
