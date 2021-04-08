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
func (u *MemoUsecase) AddMemo(memo entity.Memo) (err error) {
	err = u.memoRepository.AddMemo(memo)
	return
}
