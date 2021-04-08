package database

import (
	"fmt"
	"time"

	"github.com/kazuki-komori/memo-share/domain/entity"
	"github.com/kazuki-komori/memo-share/domain/repository"
)

type MemoRepository struct {
	SqlHandler
}

func NewMemoRepository(sqlHandler SqlHandler) repository.MemoRepository {
	memoRepository := MemoRepository{sqlHandler}
	return &memoRepository
}

// メモを登録
func (r *MemoRepository) AddMemo(memo entity.Memo) error {
	dto := new(content)
	dto.Name = memo.Name
	dto.Description = memo.Description
	dto.UserID = memo.UserID
	dto.SubjectID = memo.SubjectID
	db := r.SqlHandler.db
	res := db.Create(&dto)
	if res.Error != nil {
		return fmt.Errorf("failed to create user=%w", res.Error)
	}
	return nil
}

// メモのデータオブジェクト
type content struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
	SubjectID   string    `json:"subject_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
