package mysql

import (
	"music-backEnd/biz/model"
)

func AddRankList(rl model.RankList) (bool, error) {
	res := DB.Create(&rl)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected > 0, nil
}

func RankListOfSongList(sid int) (int, error) {
	rls := make([]*model.RankList, 0)
	res := DB.Where("song_list_id = ?", sid).Find(&rls)
	if res.Error != nil {
		return 0, res.Error
	}

	cnt := len(rls)
	sum := 0
	for i := 0; i < cnt; i++ {
		sum += rls[i].Score
	}

	return sum / cnt, nil
}

func RankListOfUserSongList(rl model.RankList) (int, error) {
	rls := make([]*model.RankList, 0)
	res := DB.Where("song_list_id = ? and consumer_id = ?", rl.SongListId, rl.ConsumerId).Find(&rls)
	if res.Error != nil {
		return 0, res.Error
	}

	if res.RowsAffected == 0 {
		return -1, nil
	}

	return rls[0].Score, nil
}