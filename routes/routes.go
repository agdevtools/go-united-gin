package routes
import (
"net/http"
"github.com/gin-gonic/gin"
controllers "go-pg-gin/controllers"
)

func Routes(router *gin.Engine) {
    router.GET("/", welcome)
    router.GET("/api/team", controllers.GetAllPlayers)
    router.POST("/api/team", controllers.CreatePlayer)
    router.GET("/api/team/:Id", controllers.GetSinglePlayer)
//     router.PUT("/api/team/{id}", controllers.UpdatePlayer)
//     router.DELETE("/api/deleteplayer/{id}", controllers.DeletePlayer)
//     router.GET("/todos", controllers.GetAllTodos)
//     router.POST("/todo", controllers.CreateTodo)
//     router.GET("/todo/:todoId", controllers.GetSingleTodo)
//     router.PUT("/todo/:todoId", controllers.EditTodo)
//     router.DELETE("/todo/:todoId", controllers.DeleteTodo)
}

func welcome(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"status":  200,
"message": "Welcome To API",
})
return
}
