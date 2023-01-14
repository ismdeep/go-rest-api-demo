package handler

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/ismdeep/go-rest-api-demo/app/server/store"
	"github.com/ismdeep/go-rest-api-demo/internal/request"
	"github.com/ismdeep/go-rest-api-demo/pkg/log"
)

type authHandler struct {
}

// Auth handler
var Auth *authHandler

// SignUp sign up
func (receiver *authHandler) SignUp(ctx context.Context, req request.SignUp) error {
	if err := store.User.Create(req.Username, req.Password); err != nil {
		log.WithContext(ctx).Error("failed to create user", zap.Error(err))
		return errors.New("failed to create user")
	}

	return nil
}
