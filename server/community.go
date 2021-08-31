package server

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/repository/mysql"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查数据库，查找到所有的 community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetails(id int64) (*models.CommunityDetails, error) {
	return mysql.GetCommunityDetailsByID(id)
}
