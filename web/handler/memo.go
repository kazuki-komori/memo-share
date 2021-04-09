package handler

import (
	"fmt"
	"net/http"

	"github.com/kazuki-komori/memo-share/domain/entity"
	"github.com/kazuki-komori/memo-share/usecase"
	"github.com/labstack/echo"
)

type MemoHandler struct {
	memoUC *usecase.MemoUsecase
}

func NewMemoHandler(memoUC *usecase.MemoUsecase) *MemoHandler {
	return &MemoHandler{memoUC: memoUC}
}

// GET /memo メモ一覧を取得
func (h *MemoHandler) GetAllMemo(c echo.Context) (err error) {
	var memos []entity.Content
	memos, err = h.memoUC.GetAllMemo()
	return c.JSON(http.StatusCreated, memos)
}

// POST /memo/create メモを登録
func (h *MemoHandler) PostMemo(c echo.Context) (err error) {
	memo := new(entity.Content)
	c.Bind(memo)
	err = h.memoUC.AddMemo(*memo)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal server error")
	}
	return c.JSON(http.StatusCreated, memo)
}

// GET /memo/:id id 指定でメモを取得
func (h *MemoHandler) GetMemoByID(c echo.Context) (err error) {
	memoID := c.Param("id")
	memo, err := h.memoUC.GetMemoByID(memoID)
	if err != nil {
		return fmt.Errorf("failed to get memo=%w", err)
	}
	return c.JSON(http.StatusOK, memo)
}
