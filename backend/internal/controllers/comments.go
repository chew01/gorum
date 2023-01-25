package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorum/internal/auth"
	"gorum/internal/database"
	"gorum/internal/models"
	"net/http"
)

type CommentController struct {
	db *database.Database
}

// UpdateComment godoc
// @Description Update a specific comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param Payload body models.CommentRequest true "Request Body"
// @Param comment_id path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Router /comments/{comment_id} [put]
func (cc *CommentController) UpdateComment(c *gin.Context) {
	var request models.CommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id := c.Param("comment_id")
	check, err := models.CheckUserComment(cc.db, id)
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

	commentData, err := models.UpdateComment(cc.db, &request, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, commentData)
}

// DeleteComment godoc
// @Description Delete a specific comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param comment_id path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Router /comments/{comment_id} [delete]
func (cc *CommentController) DeleteComment(c *gin.Context) {
	id := c.Param("comment_id")
	check, err := models.CheckUserComment(cc.db, id)
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

	commentData, err := models.DeleteComment(cc.db, id)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, commentData)
}
