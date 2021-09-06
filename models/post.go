package models

import "time"

// 内存对齐
type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Content     string    `json:"content" db:"content" binding:"required"`
	Title       string    `json:"title" db:"title" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情接口结构体
type ApiPostDetail struct {
	AuthorName        string `json:"author_name"`
	*Post             `json:"post"`
	*CommunityDetails `json:"community"`
}
