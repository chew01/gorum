package controllers

import (
	"github.com/gin-gonic/gin"
	"gorum/internal/models"
	"net/http"
)

// @BasePath /

// GetPosts godoc
// @Description Get all posts
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {array} models.SimplePost
// @Router /posts [get]
func GetPosts(c *gin.Context) {
	res, err := models.GetAllPostSummaries()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreatePost godoc
// @Description Create a new post
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {object} models.Post
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

// GetPost godoc
// @Description Get full details of a specific post
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{post_id} [get]
func GetPost(c *gin.Context) {

}

// UpdatePost godoc
// @Description Update a specific post
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{post_id} [put]
func UpdatePost(c *gin.Context) {

}

// DeletePost godoc
// @Description Delete a specific post
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{post_id} [delete]
func DeletePost(c *gin.Context) {

}

// CreateComment godoc
// @Description Create comment on a specific post
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} models.Comment
// @Router /posts/{post_id}/comments [post]
func CreateComment(c *gin.Context) {

}
