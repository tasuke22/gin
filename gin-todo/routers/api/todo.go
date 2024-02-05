package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tasuke/gin-todo/models"
	"net/http"
)

func FetchTodos(c *gin.Context) {
	todos, err := models.GetTodos()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todos)
}

type GetTodoRequest struct {
	ID int `uri:"id" binding:"required" json:"id"`
}

func FetchTodo(c *gin.Context) {
	var req GetTodoRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	todo, err := models.GetTodo(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todo)
}

type CreateTodoRequest struct {
	Name string `form:"name" binding:"required"`
	Done *bool  `form:"done" binding:"required"`
}

func CreateTodo(c *gin.Context) {
	var req CreateTodoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	todo, err := models.CreateTodo(&models.Todo{
		Name: req.Name,
		Done: *req.Done,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, todo)
}

type UpdateTodoRequest struct {
	ID   int     `uri:"id" binding:"required"`
	Name *string `form:"name"`
	Done *bool   `form:"done"`
}

func UpdateTodo(c *gin.Context) {
	var req UpdateTodoRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := models.UpdateTodo(&models.Todo{
		ID:   req.ID,
		Name: *req.Name,
		Done: *req.Done,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, todo)
}

type DeleteTodoRequest struct {
	ID int `uri:"id" json:"id" binding:"required"`
}

func DeleteTodo(c *gin.Context) {
	var req DeleteTodoRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)

	err := models.DeleteTodo(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.Status(http.StatusNoContent)
}
