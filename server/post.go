package server

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/repository/mysql"
	"github.com/Thewalkers2012/BlogBackend/util/snowflake"
	"go.uber.org/zap"
)

func CreatePost(req *models.Post) (err error) {
	// 1. 生成 post ID
	req.ID = snowflake.GetID()
	// 2. 保存到数据库
	// 3. 返回
	return mysql.CreatePost(req)
}

func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们呢接口想用的数据
	data = new(models.ApiPostDetail)
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}

	// 根据作者 id 查询作者信息
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}

	// 根据社区 id 查询社区信息
	community, err := mysql.GetCommunityDetailsByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailsByID(post.CommunityID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}

	data.AuthorName = user.Username
	data.CommunityDetails = community
	data.Post = post
	return
}

// GetPostList 获取帖子列表
func GetPostList(offset, limit int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(offset, limit)
	if err != nil {
		return nil, err
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			return nil, err
		}

		// 根据社区 id 查询社区信息
		community, err := mysql.GetCommunityDetailsByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailsByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			return nil, err
		}

		postDetail := &models.ApiPostDetail{
			AuthorName:       user.Username,
			Post:             post,
			CommunityDetails: community,
		}

		data = append(data, postDetail)
	}

	return
}
