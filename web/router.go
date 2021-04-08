package web

import (
	"net/http"
	"os"

	"github.com/kazuki-komori/memo-share/usecase"
	"github.com/kazuki-komori/memo-share/web/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewServer(userUC *usecase.UserUsecase) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := handler.NewUserHandler(userUC)

	v1 := e.Group("/api/v1")

	v1.GET("/health", health)

	v1.GET("/memo", handler.GetMemo)

	v1.POST("/user/register", userHandler.PostUser)
	v1.GET("/user/:id", userHandler.GetUserByID)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

// Active check
type Health struct {
	Status string `json:"status"`
}

func health(c echo.Context) error {
	res := &Health{
		Status: "OK",
	}

	return c.JSON(http.StatusOK, res)
}
