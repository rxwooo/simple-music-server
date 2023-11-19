package mysql

import (
	"music-backEnd/biz/model"
)

func AddCollection(col model.Collect) (bool, error) {
	var consumer *model.Consumer
	res := DB.Where("id = ?", col.UserId).First(&consumer)
	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	if col.Type == 0 {
		var tp *model.Song
		if err := DB.Where("id = ?", col.SongId).First(&tp).Error; err != nil {
			return false, err
		}
	} else {
		var tp *model.SongList
		if err := DB.Where("id = ?", col.SongListId).First(&tp).Error; err != nil {
			return false, err
		}
	}

	err := DB.Create(&col).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func DeleteCollection(uid int, sid int) (bool, error) {
	res := DB.Where("user_id = ? and song_id = ?", uid, sid).Delete(&model.Collect{})
	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func CollectiongStatus(uid int, sid int) (bool, error) {
	res := DB.Where("user_id = ? and song_id = ?", uid, sid).Find(&model.Collect{})
	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func CollectionDetail(uid int) ([]*model.Collect, error) {
	cols := make([]*model.Collect, 0)
	res := DB.Where("user_id = ?", uid).Find(&cols)
	if res.Error != nil {
		return make([]*model.Collect, 0), res.Error
	}

	return cols, nil
}
