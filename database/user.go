package database

import (
	"fmt"
	"time"

	"github.com/kazuki-komori/memo-share/domain/entity"
	"github.com/kazuki-komori/memo-share/domain/repository"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	userRepository := UserRepository{sqlHandler}
	return &userRepository
}

// ユーザーを新規登録
func (r *UserRepository) AddUser(userEntity entity.User) error {
	db := r.SqlHandler.db
	defer db.Close()
	dto := new(user)
	dto.ID = userEntity.ID
	dto.UserName = userEntity.UserName
	dto.FullName = userEntity.FullName
	dto.Mail = userEntity.Mail
	fmt.Println("hoge", dto)
	res := db.Create(&dto)
	if res.Error != nil {
		return fmt.Errorf("failed to create user=%w", res.Error)
	}
	return nil
}

// ユーザーを ID で取得
func (r *UserRepository) GetUserByID(userID string) (user entity.User, err error) {
	db := r.SqlHandler.db
	defer db.Close()
	user.ID = userID
	db.First(&user)
	return user, nil
}

type user struct {
	ID        string    `json:"id"`
	UserName  string    `json:"user_name"`
	FullName  string    `json:"full_name"`
	Mail      string    `json:"mail"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
