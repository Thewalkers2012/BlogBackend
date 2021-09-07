package redis

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/go-redis/redis"
)

func GetPostIDsInOrder(req *models.ParamPostList) ([]string, error) {
	// 从 redis 获取 id
	// 根据用户请求中携带的 order 参数确定要查询的 redis 中的 key
	key := getRedisKey(KeyPostTimeZset)
	if req.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZset)
	}

	// 2. 确定查询的索引起始点
	start := (req.Page - 1) * req.Size
	end := start + req.Size - 1

	// 3. ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}

// GetPostVoteData 根据 ids 查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	// 使用 pipeline 发送多条命令，防止发生 RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}

	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}

	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}
