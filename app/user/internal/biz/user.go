package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID    int64
	Name  string
	Email string
	Age   int32
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	ListAll(context.Context) ([]*User, error)
	Delete(context.Context, int64) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Create(ctx context.Context, u *User) (*User, error) {
	return uc.repo.Save(ctx, u)
}
