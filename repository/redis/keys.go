package redis

// redis key
// redis key尽量使用命名空间的方式，方便查询和拆分

const (
	KeyPrefix          = "blog:"
	KeyPostTimeZset    = "post:time"   // zset：帖子以发帖时间为分数
	KeyPostScoreZset   = "post:score"  // zset：帖子及投票分数
	KeyPostVotedPrefix = "post:voted:" // zset：记录用户及投票的类型；参数是 post id
)

// 给 redis key 加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}

// "post:voted:4162959018299392"
