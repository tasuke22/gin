package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tasuke/gin-todo/routers/api"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("api")
	{
		todosGroup := apiGroup.Group("todos")
		{
			todosGroup.GET("", api.FetchTodos)
			todosGroup.GET(":id", api.FetchTodo)
			todosGroup.POST("", api.CreateTodo)
			todosGroup.PATCH(":id", api.UpdateTodo)
			todosGroup.DELETE(":id", api.DeleteTodo)
		}
	}

	return router
}
