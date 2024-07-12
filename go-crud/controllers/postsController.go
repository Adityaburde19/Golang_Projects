package controllers

import (
	"github.com/gin-gonic/gin"
	"go-curd/initializers"
	"go-curd/models"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a Post
	post := models.POST{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//get the post
	var posts []models.POST
	initializers.DB.Find(&posts)
	//response with them
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")
	//get the post
	var post models.POST
	initializers.DB.First(&post, id)
	//response with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//GET THE ID OFF THE url
	id := c.Param("id")
	//Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//find the post were updating
	var post models.POST
	initializers.DB.First(&post, id)

	//update it
	initializers.DB.Model(&post).Updates(models.POST{
		Title: body.Title,
		Body:  body.Body,
	})
	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostsDelete(c *gin.Context) {
	//get the id off the url
	id := c.Param("id")
	//delete the posts
	initializers.DB.Delete(models.POST{}, id)
	//Respond
	c.Status(200)
}
