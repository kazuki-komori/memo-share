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
func (h *UserHandler) PostUser(c echo.Context) (err error) {
	user := new(entity.User)
	err = c.Bind(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid params")
	}
	err = h.userUC.AddUser(*user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to register User")
	}
	return c.JSON(http.StatusCreated, user)
}

// GET /user/:id
func (h *UserHandler) GetUserByID(c echo.Context) error {
	userID := c.Param("id")
	user, err := h.userUC.GetUserByID(userID)
	if err != nil {
		return fmt.Errorf("failed to get user by id")
	}
	return c.JSON(http.StatusOK, user)
}
