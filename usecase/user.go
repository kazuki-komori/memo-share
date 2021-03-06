package usecase

import (
	"github.com/kazuki-komori/memo-share/domain/entity"
	"github.com/kazuki-komori/memo-share/domain/repository"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepo}
}

// ユーザーを追加
func (u *UserUsecase) AddUser(user entity.User) (err error) {
	err = u.userRepository.AddUser(user)
	return
}

// ユーザーを ID から取得
func (u *UserUsecase) GetUserByID(userID string) (user entity.User, err error) {
	user, err = u.userRepository.GetUserByID(userID)
	return
}
