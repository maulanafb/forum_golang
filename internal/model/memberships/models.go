package memberships

import (
	"database/sql"
	"time"
)

type SignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserModel struct {
	ID        int64        `json:"id"`
	Email     string       `json:"email"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedBy string       `json:"created_by"`
	UpdatedBy string       `json:"updated_by"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
