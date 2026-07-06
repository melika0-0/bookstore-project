package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/melika0-0/bookstore-project/api/dto"
	"github.com/melika0-0/bookstore-project/data/models"
	"github.com/melika0-0/bookstore-project/data/cache/db"
	
)

type BookHandler struct{
	
}

func NewBookHandler() *BookHandler {
	return &BookHandler{}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var req dto.CreateBookReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validationError": err.Error(),
		})
		return
	}

	book := models.Book{
		Title:  req.Title,
		Author: req.Author,
		Price:  req.Price,
	}

	if err := db.GetDB().Create(&book).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "could not create book",
		})
		return
	}

	c.JSON(http.StatusCreated, dto.BookRes{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
		Price:  book.Price,
	})
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []models.Book

	if err := db.GetDB().Find(&books).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "could not fetch books",
		})
		return
	}

	response := make([]dto.BookRes, 0, len(books))
	for _, book := range books {
		response = append(response, dto.BookRes{
			ID:     book.ID,
			Title:  book.Title,
			Author: book.Author,
			Price:  book.Price,
		})
	}

	c.JSON(http.StatusOK, response)
}

