package handler

import (
	"fmt"
	"net/http"

	"github.com/kazuki-komori/memo-share/domain/entity"
	"github.com/kazuki-komori/memo-share/usecase"
	"github.com/labstack/echo"
)

type UserHandler struct {
	userUC *usecase.UserUsecase
}

func NewUserHandler(userUC *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUC: userUC}
}

// POST /user/register ユーザーの登録
func (h *UserHandler) PostUser(c echo.Context) error {
	user := new(entity.User)
	err := c.Bind(user)
	h.userUC.AddUser(*user)
	fmt.Println(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid params")
	}
	return c.JSON(http.StatusCreated, "{}")
}
