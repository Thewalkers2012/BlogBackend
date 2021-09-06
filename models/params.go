/*
	定义请求需要的参数
*/
package models

// 注册需要的参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// 登录需要的参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 投票所需要的数据
type ParamVoteData struct {
	// UserID 从请求中获取当前的用户 id
	PostID    string `json:"post_id" binding:"required"`                       // 帖子 id
	Direction int    `json:"direction,string" binding:"required,oneof=0 1 -1"` // 赞成票 (1) 还是反对票 (-1) 取消投票 (0)
}
