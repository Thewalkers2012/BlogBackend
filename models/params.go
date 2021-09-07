/*
	定义请求需要的参数
*/
package models

// ParamSignUp 注册需要的参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录需要的参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票所需要的数据
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户 id
	PostID    string `json:"post_id" binding:"required"`              // 帖子 id
	Direction int8   `json:"direction,string" binding:"oneof=0 1 -1"` // 赞成票 (1) 还是反对票 (-1) 取消投票 (0)
}

// 分页所需要的参数
type ParamPostList struct {
	Page  int64  `form:"page"`
	Size  int64  `form:"size"`
	Order string `form:"order"`
}

const (
	OrderTime  = "time"
	OrderScore = "score"
)
