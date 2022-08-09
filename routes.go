package main

import (
	"gin-api/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	noteController := controllers.NoteController{}

	r.GET("/", noteController.Root)

	r.GET("/notes", noteController.Index)

	r.GET("/notes/:id", noteController.Detail)

	r.POST("/notes", noteController.Create)

	r.GET("/notes/:id/delete", noteController.Delete)

	r.POST("/notes/:id", noteController.Update)

	r.GET("/notes/:id/done", noteController.Done)
}
