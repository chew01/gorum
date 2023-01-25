package controllers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mattn/go-sqlite3"
	"gorum/internal/auth"
	"gorum/internal/database"
	"gorum/internal/models"
	"net/http"
)

type UserController struct {
	db *database.Database
}

// GetUser godoc
// @Description Get current user
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /auth [get]
func (uc *UserController) GetUser(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, models.User{Name: claims.(*auth.UserClaims).Name})
}

// LoginUser godoc
// @Description Log in user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Payload body models.User true "Request Body"
// @Success 200 {object} models.UserLoginResponse
// @Router /auth [post]
func (uc *UserController) LoginUser(c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userData, err := models.SelectUser(uc.db, request.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": userData.Name,
	})

	tokenString, err := token.SignedString([]byte("JWT_SECRET")) // TODO: Change
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, models.UserLoginResponse{Token: tokenString})
}

// CreateUser godoc
// @Description Create user
// @Tags Auth
// @Accept json
// @Produce json
// @Param Payload body models.User true "Request Body"
// @Success 200 {object} models.User
// @Router /auth/create [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var request models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userData, err := models.InsertUser(uc.db, &request)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				c.AbortWithStatus(http.StatusConflict)
				return
			}
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, userData)
}
