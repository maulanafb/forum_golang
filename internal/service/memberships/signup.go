package memberships

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/maulanafb/forum_golang/internal/model/memberships"
	"golang.org/x/crypto/bcrypt"
)

// SignUp handles the user registration process
func (s *service) SignUp(ctx context.Context, req memberships.SignUpRequest) error {

	user, err := s.membershipRepo.GetUser(ctx, req.Email, req.Username)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			user = nil
		} else {
			return fmt.Errorf("failed to check existing user: %w", err)
		}
	}

	if user != nil {
		return fmt.Errorf("user already exists with email %s or username %s", req.Email, req.Username)
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
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

	// Save the user in the database
	err = s.membershipRepo.CreateUser(ctx, model)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
