package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-service/internal/app/rest"
	"task-service/internal/pkg/user"
)

type Router struct {
	ginContext *gin.Engine
	repo       user.Repo
}

func NewRouter(repo user.Repo) *Router {
	return &Router{ginContext: gin.Default(), repo: repo}
}

func (r *Router) SetUpRouter(e *gin.Engine) {
	e.GET("/users", r.getUsers)
	e.GET("/user/:id", r.getUserByID)
	e.POST("/user", r.createUser)
	e.DELETE("/user/:id", r.deleteUser)
	e.PUT("/user/:id", r.updateUser)
}

func (r *Router) getUsers(c *gin.Context) {

	users, err := r.repo.GetUsers(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (r *Router) getUserByID(c *gin.Context) {
	ID := c.Param("id")
	user, err := r.repo.GetByID(c, ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (r *Router) createUser(c *gin.Context) {
	var user user.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}

	user, err := r.repo.Create(c, user.Login, user.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, user)
}

func (r *Router) deleteUser(c *gin.Context) {
	ID := c.Param("id")
	err := r.repo.Delete(c, ID)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, "Task was deleted")
}

func (r *Router) updateUser(c *gin.Context) {
	ID := c.Param("id")
	var user user.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, rest.ErrorModel{Error: err.Error()})
		return
	}
	User, err := r.repo.Update(c, ID, user.Login, user.Email)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, rest.ErrorModel{Error: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, User)
}
