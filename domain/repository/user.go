package repository

import (
	"github.com/kazuki-komori/memo-share/domain/entity"
)

type UserRepository interface {
	AddUser(user entity.User) error
}
