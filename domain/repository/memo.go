package repository

import "github.com/kazuki-komori/memo-share/domain/entity"

type MemoRepository interface {
	AddMemo(memo entity.Content) error
	GetMemoByID(id string) (memo entity.Content, err error)
}
