package repository

import (
	"context"
	"database/sql"

	"github.com/nnnkkk7/DI_golang/model"
)

type UserRepository interface {
	CreateUser(c context.Context, user *model.User)
}
type userRepository struct {
	db sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{*db}
}
func (ur *userRepository) CreateUser(c context.Context, user *model.User) {
	ur.db.Query("insert * INTO sample")
}
