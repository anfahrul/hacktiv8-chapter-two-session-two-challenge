package routers

import (
	"hacktiv8-chapter-two-session-two-challenge/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:book_id", controllers.GetBookByID)
	router.PUT("/books/:book_id", controllers.UpdateBook)
	router.DELETE("/books/:book_id", controllers.DeleteBook)

	return router
}
