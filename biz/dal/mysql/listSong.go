package mysql

import (
	"music-backEnd/biz/model"
)

func AddListSong(ls model.ListSong) (bool, error) {
	tp := make([]*model.ListSong, 0)
	res := DB.Where("song_id = ? and song_list_id = ?", ls.SongId, ls.SongListId).Find(&tp)
	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected > 0 {
		return false, nil
	}

	res = DB.Create(&ls)
	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func DeleteListSong(ls model.ListSong) (bool, error) {
	var tpls *model.ListSong
	res := DB.Where("song_id = ? and song_list_id = ?", ls.SongId, ls.SongListId).Delete(&tpls)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func ListSongDetail(slid int) ([]*model.ListSong, error) {
	sls := make([]*model.ListSong, 0)
	res := DB.Where("song_list_id = ?", slid).Find(&sls)
	if res.Error != nil {
		return make([]*model.ListSong, 0), res.Error
	}

	return sls, nil
}

func ListSongUpdate(sl model.ListSong) (bool, error) {
	res := DB.Where("id = ?", sl.Id).Updates(sl)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}
