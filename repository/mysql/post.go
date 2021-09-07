package mysql

import (
	"strings"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/jmoiron/sqlx"
)

// CreatePost 创建帖子
func CreatePost(req *models.Post) (err error) {
	query := `insert into post(post_id, title, content, author_id, community_id)
						values(?, ?, ?, ?, ?)`
	_, err = db.Exec(query, req.ID, req.Title, req.Content, req.AuthorID, req.CommunityID)
	return err
}

// GetPostById 根据帖子 id 查询帖子详情数据
func GetPostByID(pid int64) (data *models.Post, err error) {
	data = new(models.Post)
	query := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ?`
	err = db.Get(data, query, pid)
	return
}

// GetPostList 返回 Post 列表
func GetPostList(offset, limit int64) (data []*models.Post, err error) {
	query := `select post_id, title, content, author_id, community_id, create_time from post order by create_time desc limit ?, ?`

	data = make([]*models.Post, 0, 2)
	err = db.Select(&data, query, (offset-1)*limit, limit)

	return
}

// 根据给定的 id 列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id in (?) order by FIND_IN_SET(post_id, ?)`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)

	return
}
