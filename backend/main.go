package main

import (
	"log"
	"todo_app/infra"
	"todo_app/presentation"
	"todo_app/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=app password=password dbname=app_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	tr := infra.NewTaskRepository(db)
	ur := infra.NewUserRepository(db)

	tu := usecase.NewTaskUsecase(tr)
	uu := usecase.NewUserUsecase(ur)
	lu := usecase.NewLoginUsecase(ur)

	th := presentation.NewTaskHandler(tu)
	uh := presentation.NewUserHandler(uu)
	lh := presentation.NewLoginHandler(lu)

	e := echo.New()
	presentation.InitRouting(e, th, uh, lh)
	e.Logger.Fatal(e.Start(":8888"))
}
