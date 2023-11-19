package mysql

import "music-backEnd/biz/model"

func LoginVerify(userName string, passWord string) (bool, error) {
	ads := make([]*model.Admin, 0)
	res := DB.Where("name = ?", userName).Where("password = ?", passWord).Find(&ads)
	if err := res.Error; err != nil {
		return false, err
	}
	if res.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}
