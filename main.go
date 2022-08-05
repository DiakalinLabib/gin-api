package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	r := gin.Default()

	type Note struct {
		ID      int    `json:"id"`
		Content string `json:"content"`
	}
	notes := []Note{
		{ID: 1, Content: "Kerjakan Golang!"},
		{ID: 2, Content: "Buat menggunakan Gin!"},
	}

	r.GET("/notes", func(c *gin.Context) {
		query := c.DefaultQuery("q", "")

		newNotes := notes

		if query != "" {
			newNotes = []Note{}

			for _, value := range notes {
				if strings.Contains(value.Content, query) {
					newNotes = append(newNotes, value)
				}
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Data successfully!",
			"notes":   newNotes,
		})
	})

	r.GET("/notes/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var note Note

		for _, value := range notes {
			if value.ID == id {
				note = value
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Data successfully!",
			"note":    note,
		})
	})

	r.POST("/notes", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.PostForm("id"))
		content := c.PostForm("content")

		notes = append(notes, Note{ID: id, Content: content})

		c.JSON(http.StatusCreated, gin.H{
			"message": "Data created!",
		})
	})

	r.Run()
}
