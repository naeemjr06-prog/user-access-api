package application

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/naeemjr06-prog/user-access-api/internal/user/domain"
)

type RegisterUserInput struct {
	Email    string
	Password string
}

type RegisterUserHandler struct {
	repo domain.Repository
	hash PasswordHasher
}

func NewRegisterUserHandler(repo domain.Repository, hash PasswordHasher) *RegisterUserHandler {
	return &RegisterUserHandler{repo: repo, hash: hash}
}

func (h *RegisterUserHandler) Handle(ctx context.Context, in RegisterUserInput) (*domain.User, error) {
	if !strings.Contains(in.Email, "@") {
		return nil, errors.New("invalid email")
	}
	if len(in.Password) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}

	existing, _ := h.repo.FindByEmail(ctx, in.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := h.hash.Hash(in.Password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:     in.Email,
		Password:  hashed,
		Role:      domain.RoleUser,
		Active:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}
