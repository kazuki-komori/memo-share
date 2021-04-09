package usecase

import (
	"github.com/kazuki-komori/memo-share/domain/entity"
	"github.com/kazuki-komori/memo-share/domain/repository"
)

type MemoUsecase struct {
	memoRepository repository.MemoRepository
}

func NewMemoUsecase(memoRepo repository.MemoRepository) *MemoUsecase {
	return &MemoUsecase{memoRepository: memoRepo}
}

// メモの登録
func (u *MemoUsecase) AddMemo(memo entity.Content) (err error) {
	err = u.memoRepository.AddMemo(memo)
	return
}

// メモを ID で取得
func (u *MemoUsecase) GetMemoByID(id string) (memo entity.Content, err error) {
	memo, err = u.memoRepository.GetMemoByID(id)
	return
}

// メモ一覧を取得
func (u *MemoUsecase) GetAllMemo() (memos []entity.Content, err error) {
	memos, err = u.memoRepository.GetAllMemo()
	return
}
