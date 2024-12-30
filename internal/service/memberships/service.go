package memberships

import (
	"context"

	"github.com/maulanafb/forum_golang/internal/model/memberships"
)

type MembershipRepository interface {
	GetUser(ctx context.Context, email, username string) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, user *memberships.UserModel) error
	// SignUp(ctx context.Context, req *memberships.UserModel) error
}

type service struct {
	membershipRepo MembershipRepository
}

func NewService(membershipRepo MembershipRepository) *service {
	return &service{
		membershipRepo,
	}
}
