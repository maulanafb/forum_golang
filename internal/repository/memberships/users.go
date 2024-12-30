package memberships

import (
	"context"
	"database/sql"

	"github.com/maulanafb/forum_golang/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error) {
	query := `SELECT id, email, username, password, created_at, updated_at, created_by, updated_by, deleted_at FROM users WHERE email = ? OR username = ?`

	row := r.db.QueryRowContext(ctx, query, email, username)
	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Username, &response.Password, &response.CreatedAt, &response.UpdatedAt,
		&response.CreatedBy, &response.UpdatedBy, &response.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, user *memberships.UserModel) error {
	query := `INSERT INTO users (email, password, username, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Password, user.Username, user.CreatedAt, user.UpdatedAt, user.CreatedBy, user.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
