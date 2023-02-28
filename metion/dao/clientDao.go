package dao

import (
	"bizd/metion/global"
	"bizd/metion/model"
	"bizd/metion/utils"
	"log"
)

func GetClientsDao(client model.Client) (model.ResponsePagination, []model.Client, error) {
	if client.PageSize == 0 {
		client.PageSize = 10
	}
	var clients []model.Client
	var pagination model.ResponsePagination
	var totalTmp int64
	result := global.DB.Joins("left join t_user on t_user.user_id = t_client.principal_id").
		Select("client_id,client_name,client_abbreviation,data_link,principal_id,status,canvas_account,canvas_pwd,t_user.user_name as principal_name").Where(client).Offset((client.PageNumber - 1) * client.PageSize).Limit(client.PageSize).Order("t_client.updated_at desc").Find(&clients)
	if result.Error != nil {
		log.Print(result.Error)
		return pagination, nil, result.Error
	}
	utils.TranClientStatus(clients)
	result = global.DB.Model(model.Client{}).Select("client_id,client_name,client_abbreviation,data_link,principal_id,status,canvas_account,canvas_pwd").Where(client).Count(&totalTmp)
	pagination = utils.EncapsulationPage(client.PageNumber, client.PageSize, totalTmp)
	return pagination, clients, nil
}

func GetAllClientsDao() ([]model.Client, error) {
	var clients []model.Client
	result := global.DB.Select("client_id,client_abbreviation").Order("updated_at desc").Find(&clients)
	if result.Error != nil {
		log.Print(result.Error)
		return nil, result.Error
	}
	return clients, nil
}

func GetClientsByKeywordDao(search model.Search) (model.ResponsePagination, []model.Client, error) {
	if search.PageSize == 0 {
		search.PageSize = 10
	}
	var clients []model.Client
	var pagination model.ResponsePagination
	var totalTmp int64
	result := global.DB.Joins("left join t_user on t_user.user_id = t_client.principal_id").
		Select("client_id,client_name,client_abbreviation,data_link,principal_id,status,canvas_account,canvas_pwd,t_user.user_name as principal_name").Where("client_name LIKE ? or client_abbreviation LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%").Offset((search.PageNumber - 1) * search.PageSize).Limit(search.PageSize).Order("t_client.updated_at desc").Find(&clients)
	if result.Error != nil {
		log.Print(result.Error)
		return pagination, nil, result.Error
	}
	utils.TranClientStatus(clients)
	result = global.DB.Model(model.Client{}).Select("client_id,client_name,client_abbreviation,data_link,principal_id,status,canvas_account,canvas_pwd").Where("client_name LIKE ? or client_abbreviation LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%").Count(&totalTmp)
	pagination = utils.EncapsulationPage(search.PageNumber, search.PageSize, totalTmp)
	return pagination, clients, nil
}

func DelClientByKeys(del model.DelModel) error {
	result := global.DB.Model(model.Client{}).Delete("clientId", del.Keys)
	if result.Error != nil {
		log.Print(result.Error)
		return result.Error
	} else {
		return nil
	}
}
