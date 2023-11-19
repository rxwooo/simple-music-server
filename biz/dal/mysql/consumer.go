package mysql

import (
	"music-backEnd/biz/model"
)

func UserExistence(uname string) (bool, error) {
	tp := make([]*model.Consumer, 0)
	res := DB.Where("username = ?", uname).Find(&tp)
	if res.Error != nil {
		return false, res.Error
	}
	return res.RowsAffected == 0, nil
}

func AddConsumer(user model.Consumer) (bool, error) {
	res := DB.Create(&user)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func ConsumerVerify(uname string, passwd string) (bool, error) {
	users := make([]*model.Consumer, 0)
	res := DB.Where("username = ? and password = ?", uname, passwd).Find(&users)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func LoginStatus(uname string) ([]*model.Consumer, error) {
	users := make([]*model.Consumer, 0)
	res := DB.Where("username = ?", uname).Find(&users)

	if res.Error != nil {
		return make([]*model.Consumer, 0), res.Error
	}

	return users, nil
}

func AllUsers() ([]*model.Consumer, error) {
	users := make([]*model.Consumer, 0)
	if err := DB.Find(&users).Error; err != nil {
		return make([]*model.Consumer, 0), err
	}

	return users, nil
}

func UserDetail(uid int) (model.Consumer, error) {
	var user *model.Consumer
	res := DB.Where("id = ?", uid).First(&user)
	if res.Error != nil {
		return model.Consumer{}, res.Error
	}

	if res.RowsAffected == 0 {
		return model.Consumer{}, nil
	}

	return *user, nil
}

func DeleteUser(uid int) (bool, error) {
	var user *model.Consumer
	res := DB.Where("id = ?", uid).Delete(&user)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func UpdateUser(user model.Consumer) (bool, error) {
	//res := DB.Where("id = ?", user.Id).Updates(user)
	res := DB.Model(&model.Consumer{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"username":     user.Username,
		"sex":          user.Sex,
		"phone_num":    user.PhoneNum,
		"email":        user.Email,
		"birth":        user.Birth,
		"introduction": user.Introduction,
		"location":     user.Location,
	})
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func UserVerifyPassword(uname string, passwd string) (bool, error) {
	var user *model.Consumer
	res := DB.Where("username = ? and password = ?", uname, passwd).First(&user)

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func UserUpdatePasswd(uname string, passwd string) (bool, error) {
	res := DB.Where("username = ?", uname).Updates(model.Consumer{Password: passwd})
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func UserUpdateAvatar(uid int, apath string) (bool, error) {
	res := DB.Where("id = ?", uid).Updates(model.Consumer{Avator: apath})
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func GetUserCover(uid int) (string, error) {
	var user *model.Consumer
	res := DB.Where("id = ?", uid).First(&user)
	return user.Avator, res.Error
}
