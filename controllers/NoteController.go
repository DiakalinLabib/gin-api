package controllers

import (
	"gin-api/database"
	"gin-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type NoteController struct{}

func (NoteController) Root(c *gin.Context) {
	c.Redirect(http.StatusFound, "/notes")
}

func (NoteController) Index(c *gin.Context) {
	db := database.Instance()

	query := c.DefaultQuery("q", "")
	var notes []models.Note
	db.Find(&notes, query)

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Data successfully!",
	//	"notes":   notes,
	//})

	c.HTML(http.StatusOK, "notes/index", gin.H{
		"title": "Notes Index",
		"notes": notes,
	})
}

func (NoteController) Detail(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Note
	result := db.First(&note, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Data not found!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Data successfully!",
			"note":    note,
		})
	}
}

func (NoteController) Create(c *gin.Context) {
	db := database.Instance()
	content := c.PostForm("content")

	db.Create(&models.Note{Content: content})

	//c.JSON(http.StatusCreated, gin.H{
	//	"message": "Data created!",
	//})

	c.Redirect(http.StatusFound, "/notes")
}

func (NoteController) Delete(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	db.Delete(&models.Note{}, id)

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Data deleted!",
	//})

	c.Redirect(http.StatusFound, "/notes")
}

func (NoteController) Update(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))
	content := c.PostForm("content")

	db.Model(&models.Note{}).
		Where("id  = ?", id).
		Update("content", content)

	c.JSON(http.StatusOK, gin.H{
		"message": "Data updated!",
	})
}

func (NoteController) Done(c *gin.Context) {
	db := database.Instance()
	id, _ := strconv.Atoi(c.Param("id"))

	var note models.Note
	db.Find(&note, id)

	note.IsDone = !note.IsDone

	db.Save(&note)

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Data updated!",
	//})

	c.Redirect(http.StatusFound, "/notes")
}
