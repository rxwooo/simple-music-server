package mysql

import (
	"fmt"
	"music-backEnd/biz/model"
)

func AddSinger(sg model.Singer) (bool, error) {
	res := DB.Create(&sg)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func DeleteSinger(sid int) (bool, error) {
	var sg *model.Singer
	res := DB.Where("id = ?", sid).Delete(&sg)

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func AllSinger() ([]*model.Singer, error) {
	sgs := make([]*model.Singer, 0)
	res := DB.Find(&sgs)

	if res.Error != nil {
		return make([]*model.Singer, 0), res.Error
	}
	return sgs, nil
}

func SingerLikeName(name string) ([]*model.Singer, error) {
	sgs := make([]*model.Singer, 0)
	res := DB.Where("name like ?", name).Find(&sgs)

	if res.Error != nil {
		return make([]*model.Singer, 0), res.Error
	}
	return sgs, nil
}

func SingerSex(sex int) ([]*model.Singer, error) {
	sgs := make([]*model.Singer, 0)
	res := DB.Where("sex = ?", sex).Find(&sgs)

	if res.Error != nil {
		return make([]*model.Singer, 0), res.Error
	}
	return sgs, nil
}

func SingerUpdate(sg model.Singer) (bool, error) {
	fmt.Println(sg)
	res := DB.Model(&model.Singer{}).Where("id = ?", sg.Id).Updates(map[string]interface{}{
		"name":         sg.Name,
		"sex":          sg.Sex,
		"birth":        sg.Birth,
		"location":     sg.Location,
		"introduction": sg.Introduction,
	})
	//res = DB.Where("id = ?", sg.Id).Updates(model.Singer{Sex: 0})
	if res.Error != nil {
		//println(res.Error)
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func SingerUpdateAvatar(singerId int, avaPath string) (bool, error) {
	//var singer model.Singer
	res := DB.Where("id = ?", singerId).Updates(model.Singer{Pic: avaPath})
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}
