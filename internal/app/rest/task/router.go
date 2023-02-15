package task

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-service/internal/app/rest"
	"task-service/internal/pkg/task"
)

type Router struct {
	ginContext *gin.Engine
	repo       task.Repo
}

func NewRouter(repo task.Repo) *Router {
	return &Router{ginContext: gin.Default(), repo: repo}
}

func (r *Router) SetUpRouter(e *gin.Engine) {
	e.GET("/tasks", r.getTasks)
	e.GET("/task/:id", r.getTaskByID)
	e.POST("/task", r.createTask)
	e.DELETE("/task/:id", r.deleteTask)
	e.PUT("/task/:id", r.updateTask)
}

func (r *Router) createTask(c *gin.Context) {
	var task task.Task
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}

	task, err := r.repo.Create(c, task.Name, task.UserID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, task)
}

func (r *Router) getTasks(c *gin.Context) {

	tasks, err := r.repo.GetTasks(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (r *Router) getTaskByID(c *gin.Context) {
	ID := c.Param("id")
	task, err := r.repo.GetByID(c, ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (r *Router) deleteTask(c *gin.Context) {
	ID := c.Param("id")
	err := r.repo.Delete(c, ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, "Task was deleted")
}

func (r *Router) updateTask(c *gin.Context) {
	ID := c.Param("id")
	var task task.Task
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}
	Task, err := r.repo.Update(c, ID, task.Name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, Task)
}
