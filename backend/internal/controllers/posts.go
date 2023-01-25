package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-sqlite3"
	"gorum/internal/auth"
	"gorum/internal/database"
	"gorum/internal/models"
	"net/http"
)

type PostController struct {
	db *database.Database
}

// GetPosts godoc
// @Description Get all posts
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {array} models.SimplePost
// @Router /posts [get]
func (pc *PostController) GetPosts(c *gin.Context) {
	posts, err := models.GetSimplePosts(pc.db)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, posts)
}

// CreatePost godoc
// @Description Create a new post
// @Tags Post
// @Accept json
// @Produce json
// @Param Payload body models.PostRequest true "Request Body"
// @Success 200 {object} models.Post
// @Router /posts [post]
func (pc *PostController) CreatePost(c *gin.Context) {
	var request models.PostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	claims, ok := c.Get("claims")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	postData, err := models.CreatePost(pc.db, &request, claims.(*auth.UserClaims).Name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, postData)
}

// GetPost godoc
// @Description Get full details of a specific post
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{post_id} [get]
func (pc *PostController) GetPost(c *gin.Context) {
	id := c.Param("post_id")
	postData, err := models.GetPost(pc.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, postData)
}

// UpdatePost godoc
// @Description Update a specific post
// @Tags Post
// @Accept json
// @Produce json
// @Param Payload body models.PostRequest true "Request Body"
// @Param post_id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{post_id} [put]
func (pc *PostController) UpdatePost(c *gin.Context) {
	var request models.PostRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id := c.Param("post_id")
	check, err := models.CheckUser(pc.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	claims, ok := c.Get("claims")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if check.Name != claims.(*auth.UserClaims).Name {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	postData, err := models.UpdatePost(pc.db, &request, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, postData)
}

// DeletePost godoc
// @Description Delete a specific post
// @Tags Post
// @Accept json
// @Produce json
// @Param post_id path int true "Post ID"
// @Success 200 {object} models.Post
// @Router /posts/{post_id} [delete]
func (pc *PostController) DeletePost(c *gin.Context) {
	id := c.Param("post_id")
	check, err := models.CheckUser(pc.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	claims, ok := c.Get("claims")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if check.Name != claims.(*auth.UserClaims).Name {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	postData, err := models.DeletePost(pc.db, id)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, postData)
}

// CreateComment godoc
// @Description Create comment on a specific post
// @Tags Post
// @Accept json
// @Produce json
// @Param Payload body models.CommentRequest true "Request Body"
// @Param post_id path int true "Post ID"
// @Success 200 {object} models.Comment
// @Router /posts/{post_id}/comments [post]
func (pc *PostController) CreateComment(c *gin.Context) {
	var request models.CommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id := c.Param("post_id")
	claims, ok := c.Get("claims")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	commentData, err := models.CreateComment(pc.db, &request, id, claims.(*auth.UserClaims).Name)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintForeignKey {
				c.AbortWithStatus(http.StatusNotFound)
				return
			}
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, commentData)
}
