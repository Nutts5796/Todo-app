package main

import (
	"github.com/Nutts5796/todo-app/db"
	"github.com/Nutts5796/todo-app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDatabase()

	r := routes.SetupRouter()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Todo App!")
	})

	r.Run(":4040")
}
