package memberships

import (
	"context"
	"errors"

	"github.com/maulanafb/forum_golang/internal/model/memberships"
	"github.com/maulanafb/forum_golang/pkg/jwt"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("Error in membershipRepo.GetUser")
		log.Error().Err(err).Msg("failed to get user")
		return "", err
	}

	if user == nil {
		return "", errors.New("email not exists")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		log.Error().Err(err).Msg("Error in bcrypt.CompareHashAndPassword")
		log.Error().Err(err).Msg("password not match")
		return "", errors.New("password not match")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJwt)
	if err != nil {
		log.Error().Err(err).Msg("Error in jwtService.CreateToken")
		log.Error().Err(err).Msg("failed to create token")
		return "", err
	}

	return token, nil

}
