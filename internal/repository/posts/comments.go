package posts

import (
	"context"
	"github.com/maulanafb/forum_golang/internal/model/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by,deleted_at) VALUES (?, ?, ?, ?, ?, ?, ?,?)`
	_, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy, model.DeletedAt)
	if err != nil {
		return err
	}
	return nil
}
