package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"task-service/internal/app/rest"
	"task-service/internal/app/rest/task"
	"task-service/internal/app/rest/user"
	taskpkg "task-service/internal/pkg/task"
	userpkg "task-service/internal/pkg/user"
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

	taskRouter := task.NewRouter(taskRepo)
	userRouter := user.NewRouter(userRepo)

	router := rest.NewRouter(taskRouter, userRouter)
	router.SetUpRouter()
	router.Run()
}
