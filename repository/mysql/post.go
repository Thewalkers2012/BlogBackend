package mysql

import "github.com/Thewalkers2012/BlogBackend/models"

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
