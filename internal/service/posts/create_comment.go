package posts

import (
	"context"
	"github.com/maulanafb/forum_golang/internal/model/posts"
	"github.com/rs/zerolog/log"
	"strconv"
	"time"
)

func (s *service) CreateComment(ctx context.Context, postID, userId int64, request posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		PostID:         postID,
		UserID:         userId,
		CommentContent: request.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userId, 10),
		UpdatedBy:      strconv.FormatInt(userId, 10),
		DeletedAt:      now,
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Error in postRepo.CreateComment")
		return err
	}
	return nil
}
