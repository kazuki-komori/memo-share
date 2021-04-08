package repository

import "github.com/kazuki-komori/memo-share/domain/entity"

type MemoRepository interface {
	AddMemo(memo entity.Memo) error
}
