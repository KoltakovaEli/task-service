package usecase

import "task-service/internal/usecase/task"

type UseCase interface {
	//User() user.UseCase
	Task() task.UseCase
}
