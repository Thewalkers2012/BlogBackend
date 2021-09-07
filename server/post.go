package server

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/repository/mysql"
	"github.com/Thewalkers2012/BlogBackend/repository/redis"
	"github.com/Thewalkers2012/BlogBackend/util/snowflake"
	"go.uber.org/zap"
)

func CreatePost(req *models.Post) (err error) {
	// 1. 生成 post ID
	req.ID = snowflake.GetID()
	// 2. 保存到数据库, 在 redis 中进行记录
	err = mysql.CreatePost(req)
	if err != nil {
		return err
	}
	return redis.CreatePost(req.ID)
	// 3. 返回
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

func GetPostList2(req *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// 2. 去 redis 中查询 id 列表
	ids, err := redis.GetPostIDsInOrder(req)
	if err != nil {
		return
	}

	if len(ids) == 0 {
		// zap.L().Warn("redis.GetPostList2(req) return 0 data")
		zap.L().Warn("redis.GetPostList2(req) return 0 data")
		return
	}
	// 3. 根据 id 去 mysql 数据库中查询帖子详细信息
	// 返回的数据要按照 给定的 id 进行返回
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}

	data = make([]*models.ApiPostDetail, 0, len(posts))
	// 提前查询好数据

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者及分区信息查询，并返回
	for idx, post := range posts {
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
			VoteNum:          voteData[idx],
			AuthorName:       user.Username,
			Post:             post,
			CommunityDetails: community,
		}

		data = append(data, postDetail)
	}

	return
}
