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
	result := global.DB.Select("user_id,user_name,wx_name,type,current_workload,user_account").Where(user).Offset((user.PageNumber - 1) * user.PageSize).Limit(user.PageSize).Order("updated_at desc").Find(&users)
	if result.Error != nil {
		log.Print(result.Error)
		return pagination, nil, result.Error
	}
	utils.TranUserType(users)
	result = global.DB.Model(model.User{}).Where(user).Count(&totalTmp)
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
	result := global.DB.Select("user_id,user_name,wx_name,type,current_workload,user_account").Where("user_name LIKE ? or wx_name LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%").Offset((search.PageNumber - 1) * search.PageSize).Limit(search.PageSize).Order("updated_at desc").Find(&users)
	if result.Error != nil {
		log.Print(result.Error)
		return pagination, nil, result.Error
	}
	utils.TranUserType(users)
	result = global.DB.Model(model.User{}).Where("user_name LIKE ? or wx_name LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%").Count(&totalTmp)
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

func GetAllUsers() []model.User {
	var users []model.User
	result := global.DB.Select("user_id,user_name").Order("updated_at desc").Find(&users)
	if result.Error != nil {
		log.Print(result.Error)
		panic(model.MyError{Code: 400, Message: result.Error.Error()})
	}
	return users
}

func DelUserByKeys(del model.DelModel) error {
	result := global.DB.Model(model.User{}).Delete("userId", del.Keys)
	if result.Error != nil {
		log.Print(result.Error)
		return result.Error
	} else {
		return nil
	}
}
