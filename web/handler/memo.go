package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

// GET /memo メモ一覧を取得
func GetMemo(c echo.Context) error {
	return c.JSON(http.StatusOK, toMemoJSON())
}

func toMemoJSON() *memosJSON {
	mj := &memoJSON{
		ID:            1,
		ContentsTitle: "メモのタイトル",
		SubjectID:     1,
	}
	mjs := memosJSON{}
	mjs.Memos = append(mjs.Memos, mj)
	return &mjs
}

func PostMemo(c echo.Context) error {
	return c.JSON(http.StatusCreated, "{}")
}

type memoJSON struct {
	ID            int    `json:"id"`
	ContentsTitle string `json:"contents_title"`
	SubjectID     int    `json:"subject_id"`
}

type memosJSON struct {
	Memos []*memoJSON `json:"contents"`
}
