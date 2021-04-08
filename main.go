package main

import (
	"fmt"
	"os"

	"github.com/kazuki-komori/memo-share/database"
	"github.com/kazuki-komori/memo-share/usecase"
	"github.com/kazuki-komori/memo-share/web"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		fmt.Print("failed to create db=%w", err)
		os.Exit(1)
	}
	userRepo := database.NewUserRepository(*db)
	userUC := usecase.NewUserUsecase(userRepo)
	memoRepo := database.NewMemoRepository(*db)
	memoUC := usecase.NewMemoUsecase(memoRepo)
	web.NewServer(userUC, memoUC)
}
