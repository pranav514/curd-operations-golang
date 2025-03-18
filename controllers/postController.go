package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pranav514/curd-operations-golang/initializers"
	"github.com/pranav514/curd-operations-golang/models"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"post": post,
	})

}

func GetPost(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusAccepted, gin.H{
		"posts": posts,
	})
}

func GetSpecific(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(http.StatusAccepted, gin.H{
		"post": post,
	})

}

func Update(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	var post models.Post
	initializers.DB.First(&post, id)
	posts := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	initializers.DB.Model(&post).Updates(posts)
	c.JSON(http.StatusAccepted, gin.H{
		"post": post,
	})

}

func Delete(c *gin.Context){
	id := c.Param("id");
	var post models.Post
	initializers.DB.Delete(&post , id)
	c.JSON(http.StatusAccepted , gin.H{
		"post" : post,
	})

}
