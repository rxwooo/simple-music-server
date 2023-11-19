package mysql

import (
	"music-backEnd/biz/model"
)

func AddComment(com model.Comment) (bool, error) {
	var user *model.Consumer
	res := DB.Where("id = ?", com.UserId).First(&user)
	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	err := DB.Create(&com).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteComment(cid int) (bool, error) {
	res := DB.Where("id = ?", cid).Delete(&model.Comment{})
	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func SongCommentDetail(sid int) ([]*model.Comment, error) {
	coms := make([]*model.Comment, 0)
	if err := DB.Where("song_id = ?", sid).Find(&coms).Error; err != nil {
		return make([]*model.Comment, 0), err
	}

	return coms, nil
}

func SongListCommentDetail(slid int) ([]*model.Comment, error) {
	coms := make([]*model.Comment, 0)
	if err := DB.Where("song_list_id = ?", slid).Find(&coms).Error; err != nil {
		return make([]*model.Comment, 0), err
	}

	return coms, nil
}

func UpdateComment(com model.Comment) (bool, error) {
	res := DB.Where("id = ?", com.Id).Updates(model.Comment{Up: com.Up})
	if res.Error != nil {
		return false, res.Error
	}

	if res.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
