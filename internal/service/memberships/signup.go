package memberships

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/maulanafb/forum_golang/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}

	if user != nil {
		return errors.New("user already exists")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	now := time.Now()
	model := &memberships.UserModel{
		Email:     req.Email,
		Password:  string(pass),
		Username:  req.Username,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.Email,
		UpdatedBy: req.Email,
		DeletedAt: sql.NullTime{},
	}

	err = s.membershipRepo.CreateUser(ctx, model)
	if err != nil {
		return err
	}

	return nil
}
