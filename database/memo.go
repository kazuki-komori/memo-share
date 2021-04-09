package database

import (
	"fmt"

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
	memo.ID = id
	db.First(&memo)
	fmt.Println(memo)
	return memo, err
}

// メモ一覧を取得
func (r *MemoRepository) GetAllMemo() (memos []entity.Content, err error) {
	db := r.SqlHandler.db
	defer db.Close()
	db.Find(&memos)
	return memos, err
}
