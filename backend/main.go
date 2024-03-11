package main

import (
	"log"
	"todo_app/infra"
	"todo_app/presentation"
	"todo_app/usecase"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
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
	th := presentation.NewTaskHandler(tu)
	uh := presentation.NewUserHandler(uu)

	e := echo.New()
	presentation.InitRouting(e, th, uh)
	e.Logger.Fatal(e.Start(":8888"))
}
