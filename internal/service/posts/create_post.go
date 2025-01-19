package posts

import (
	"context"
	"github.com/maulanafb/forum_golang/internal/model/posts"
	"github.com/rs/zerolog/log"

	"strconv"
	"strings"
	"time"
)

func (s *service) CreatePost(ctx context.Context, UserID int64, req posts.CreatePostRequest) error {
	postHashtags := strings.Join(req.PostHashtag, ",")
	now := time.Now()
	model := posts.PostModel{
		UserID:      UserID,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
		PostHashtag: postHashtags,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   strconv.FormatInt(UserID, 10),
		UpdatedBy:   strconv.FormatInt(UserID, 10),
	}

	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Log().Err(err).Msg("Error in postRepo.CreatePost")
		log.Log().Err(err).Msg("failed to create post")
		return err
	}
	return nil
}
