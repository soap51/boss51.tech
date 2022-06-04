package repositories

import (
	"context"

	"github.com/soap51/boss51.tech/internal/domain"
)

type mongoUserRepository struct {
}

func NewMongoUserRepository() *mongoUserRepository {
	return &mongoUserRepository{}
}

func (u *mongoUserRepository) Find(c context.Context, id string) (*domain.User, error) {
	return nil, nil
}

func (u *mongoUserRepository) Update(c context.Context, user domain.User, id string) error {
	return nil
}

func (u *mongoUserRepository) Create(c context.Context, user domain.User) (*domain.User, error) {
	return nil, nil
}

func (u *mongoUserRepository) Delete(c context.Context, id string) error {
	return nil
}
