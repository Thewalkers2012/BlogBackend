package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

var (
	ErrorVoteTimeExpire = errors.New("投票时间已过")
)

/*
	投票的几种情况
	direction = 1时，有两种情况：
		1. 之前没有投过票，现在投赞成票 ---------> 更新分数和投票记录 差值的绝对值：1 +432
		2. 之前投返对票，现在该投赞成 ---------> 更新分数和投票记录 差值的绝对值：2 +432 * 2

	direction = 0时，有两种情况：
		1. 之前投过赞成票，现在要取消投票 ---------> 更新分数和投票记录 差值的绝对值：1 -432
		2. 之前投过反对票，现在要取消投票 ---------> 更新分数和投票记录 差值的绝对值：1 + 432

	direction = -1时，有两种情况：
		1. 之前没有投过票，现在投反对票 ---------> 更新分数和投票记录 差值的绝对值：1 -432
		2. 之前投赞成票，现在改为反对票 ---------> 更新分数和投票记录 差值的绝对值：2 -432 * 2

	投票的限制：
	每个帖子自发表之日起一个星期之内允许投票，超过一个星期就不允许投票了
		1. 到期之后将 redis 中保存的赞成票及反对票存储到 mysql 表中
		2. 到期之后删除那个 KeyPostVotedZSetPF
*/

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票的限制
	// 去帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZset), postID).Val()
	if time.Now().Unix()-int64(postTime) > oneWeekInSeconds {
		return ErrorVoteTimeExpire // 时间超过一个星期就不能再进行投票了
	}
	// 2. 更新分数
	// 先查当前用户给当前帖子的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedPrefix+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	// 原子性地增加分数
	_, err := client.ZIncrBy(getRedisKey(KeyPostScoreZset), op*diff*scorePerVote, postID).Result()
	if err != nil {
		return err
	}
	// 3. 记录用户为该帖子投过票的数据
	if value == 0 {
		// 移除
		_, err = client.ZRem(getRedisKey(KeyPostVotedPrefix+postID), userID).Result()
	} else {
		_, err = client.ZAdd(getRedisKey(KeyPostVotedPrefix+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		}).Result()
	}

	return err
}

func CreatePost(postID int64) error {
	_, err := client.ZAdd(getRedisKey(KeyPostTimeZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: nil,
	}).Result()
	zap.L().Error("redis.CreatePost() failed, ", zap.Error(err))
	return err
}
