package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func setup(t *testing.T, req *http.Request, rec *httptest.ResponseRecorder) echo.Context {
	// Setup Echo
	e := echo.New()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := e.NewContext(req, rec)
	return c
}

// メモ一覧を取得するテスト
func TestGetMemo(t *testing.T) {
	// expectJSON := `{"contents":[{"id":1,"contents_title":"メモのタイトル","subject_id":1}]}`
	// req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
	// rec := httptest.NewRecorder()
	// c := setup(t, req, rec)
	// c.SetPath("/memo")
	// if assert.NoError(t, handler.GetMemo(c)) {
	// 	assert.Equal(t, strings.TrimRight(rec.Body.String(), "\n"), expectJSON)
	// }
}
