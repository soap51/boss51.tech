package driven

import (
	"context"

	"github.com/soap51/boss51.tech/internal/domain"
)

type IUserRepository interface {
	Find(c context.Context, id string) (*domain.User, error)
	Update(c context.Context, user domain.User, id string) error
	Create(c context.Context, user domain.User) (*domain.User, error)
	Delete(c context.Context, id string) error
}
