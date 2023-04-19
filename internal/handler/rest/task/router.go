package task

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"task-service/internal/entity"
	"task-service/internal/handler/rest"
	usctask "task-service/internal/usecase/task"
)

type Router struct {
	ginContext *gin.Engine
	service    usctask.UseCase
}

func NewRouter(service usctask.UseCase) *Router {
	return &Router{ginContext: gin.Default(), service: service}
}

func (r *Router) SetUpRouter(e *gin.Engine) {
	e.GET("/tasks", r.getTasks)
	e.GET("/task/:id", r.getTaskByID)
	e.POST("/task", r.createTask)
	e.DELETE("/task/:id", r.deleteTask)
	e.PUT("/task/:id", r.updateTask)
}

func (r *Router) createTask(c *gin.Context) {
	var task *entity.Task
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}

	if task.Name == "" || task.UserID == "" || !IsValidUUID(task.UserID) {
		c.IndentedJSON(http.StatusBadRequest, "invalid name or user id")
		return
	}

	task, err := r.service.Create(c, task.Name, task.UserID)
	if err != nil {
		if errors.Is(err, usctask.ErrUserNotFound) {
			c.IndentedJSON(http.StatusNotFound, err.Error())
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, task)
}

func (r *Router) getTasks(c *gin.Context) {

	tasks, err := r.service.GetTasks(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func (r *Router) getTaskByID(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" || !IsValidUUID(ID) {
		c.IndentedJSON(http.StatusBadRequest, "invalid id")
		return
	}
	task, err := r.service.GetByID(c, ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func (r *Router) deleteTask(c *gin.Context) {
	ID := c.Param("id")
	err := r.service.Delete(c, ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, "Task was deleted")
}

func (r *Router) updateTask(c *gin.Context) {
	ID := c.Param("id")
	var task entity.Task
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}
	Task, err := r.service.Update(c, ID, task.Name)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, Task)
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
