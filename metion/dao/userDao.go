package dao

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"log"
)

func GetUsersDao(user model.User) (model.ResponsePagination, []model.User, error) {
	if user.PageSize == 0 {
		user.PageSize = 10
	}
	var users []model.User
	var pagination model.ResponsePagination
	var totalTmp int64
	result := global.DB.Select("user_id,user_name,wx_name,type,current_workload").Where(user).Offset((user.PageNumber - 1) * user.PageSize).Limit(user.PageSize).Find(&users)
	if result.Error != nil {
		log.Print(result.Error)
		return pagination, nil, result.Error
	}
	result = global.DB.Model(model.User{}).Select("user_id,user_name,wx_name,type,current_workload").Where(user).Count(&totalTmp)
	pagination = utils.EncapsulationPage(user.PageNumber, user.PageSize, totalTmp)
	return pagination, users, nil
}

func GetUsersByKeywordDao(search model.Search) (model.ResponsePagination, []model.User, error) {
	if search.PageSize == 0 {
		search.PageSize = 10
	}
	var users []model.User
	var pagination model.ResponsePagination
	var totalTmp int64
	result := global.DB.Select("user_id,user_name,wx_name,type,current_workload").Where("user_name LIKE ? or wx_name LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%").Offset((search.PageNumber - 1) * search.PageSize).Limit(search.PageSize).Find(&users)
	if result.Error != nil {
		log.Print(result.Error)
		return pagination, nil, result.Error
	}
	result = global.DB.Select("user_id,user_name,wx_name,type,current_workload").Where("user_name LIKE ? or wx_name LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%").Count(&totalTmp)
	pagination = utils.EncapsulationPage(search.PageNumber, search.PageSize, totalTmp)
	return pagination, users, nil
}

func GetUserDao(userAccount string, password string) (model.User, error) {
	var user model.User
	result := global.DB.First(&user, "user_account = ? and user_pwd = ?", userAccount, password)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
