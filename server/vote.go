package server

import (
	"strconv"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/repository/redis"
	"go.uber.org/zap"
)

// 投票功能
// 1. 用户投票的数据

// 本项目使用简化版的投票分数
// 投一票就加 432 分，86400/200 -> 需要 200 张赞成票，可以给你的帖子续一天 -> 《 redis实战 》

/*
	投票的几种情况
	direction = 1时，有两种情况：
		1. 之前没有投过票，现在投赞成票 ---------> 更新分数和投票记录
		2. 之前投返对票，现在该投赞成 ---------> 更新分数和投票记录

	direction = 0时，有两种情况：
		1. 之前投过赞成票，现在要取消投票 ---------> 更新分数和投票记录
		2. 之前投过反对票，现在要取消投票 ---------> 更新分数和投票记录

	direction = -1时，有两种情况：
		1. 之前没有投过票，现在投反对票 ---------> 更新分数和投票记录
		2. 之前投赞成票，现在改为反对票 ---------> 更新分数和投票记录

	投票的限制：
	每个帖子自发表之日起一个星期之内允许投票，超过一个星期就不允许投票了
		1. 到期之后将 redis 中保存的赞成票及反对票存储到 mysql 表中
		2. 到期之后删除那个 KeyPostVotedZSetPF
*/

// VoteForPost 为帖子投票的函数
func VoteForPost(userID int64, req *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.String("postID", req.PostID),
		zap.Int8("direction", int8(req.Direction)),
	)
	return redis.VoteForPost(strconv.Itoa(int(userID)), req.PostID, float64(req.Direction))
}
