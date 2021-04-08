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
func (r *MemoRepository) AddMemo(memo entity.Content) error {
	// dto := new(content)
	// dto.ID = CreateULID()
	// dto.ContentTitle = memo.ContentTitle
	// dto.Description = memo.Description
	// dto.UserID = memo.UserID
	// dto.SubjectID = memo.SubjectID
	db := r.SqlHandler.db
	res := db.Create(&memo)
	if res.Error != nil {
		return fmt.Errorf("failed to create user=%w", res.Error)
	}
	return nil
}

func (r *MemoRepository) GetMemoByID(id string) (memo entity.Content, err error) {
	db := r.SqlHandler.db
	defer db.Close()
	dto := new(content)
	dto.ID = id
	memo.ID = dto.ID
	memo.ContentTitle = dto.ContentTitle
	memo.Description = dto.Description
	db.First(&memo)
	fmt.Println(memo)
	return memo, err
}

// メモのデータオブジェクト
type content struct {
	ID           string    `json:"id"`
	ContentTitle string    `json:"content_title"`
	Description  string    `json:"description"`
	UserID       string    `json:"user_id"`
	SubjectID    string    `json:"subject_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
