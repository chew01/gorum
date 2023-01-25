package controllers

import "gorum/internal/database"

type MainController struct {
	PostController
	CommentController
	UserController
}

func NewMainController(db *database.Database) *MainController {
	return &MainController{
		PostController:    PostController{db},
		CommentController: CommentController{db},
		UserController:    UserController{db},
	}
}
