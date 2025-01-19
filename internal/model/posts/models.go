package posts

import "time"

type (
	CreatePostRequest struct {
		PostTitle   string   `json:"postTitle" binding:"required"`
		PostContent string   `json:"postContent" binding:"required"`
		PostHashtag []string `json:"postHashtag" binding:"required"`
	}
)

type (
	PostModel struct {
		ID          int64     `db:"id"`
		UserID      int64     `db:"user_id"`
		PostTitle   string    `db:"post_title"`
		PostContent string    `db:"post_content"`
		PostHashtag string    `db:"post_hashtag"`
		CreatedAt   time.Time `db:"created_at"`
		UpdatedAt   time.Time `db:"updated_at"`
		CreatedBy   string    `db:"created_by"`
		UpdatedBy   string    `db:"updated_by"`
		DeletedAt   time.Time `db:"deleted_at"`
	}
)
