package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"task-service/internal/handler/rest"
	"task-service/internal/handler/rest/task"
	"task-service/internal/handler/rest/user"
	taskpkg "task-service/internal/repository/task"
	userpkg "task-service/internal/repository/user"
	task2 "task-service/internal/usecase/task"
)

const (
	host     = "localhost"
	port     = 15432
	dbUser   = "koltakova_e"
	password = "task-service"
	dbname   = "koltakova_e"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, dbUser, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	taskRepo := taskpkg.NewPostgresRepository(db)
	userRepo := userpkg.NewPostgresRepository(db)

	taskService := task2.NewService(taskRepo, userRepo)

	taskRouter := task.NewRouter(taskService)
	userRouter := user.NewRouter(userRepo)

	router := rest.NewRouter(taskRouter, userRouter)
	router.SetUpRouter()
	router.Run()
}
