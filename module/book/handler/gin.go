package handler

import (
	"errors"
	"golang-hexagonal-arch/module/book/dto"
	"golang-hexagonal-arch/module/book/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ginHandler struct {
	engine  *gin.Engine
	service service.BookServiceInterface
}

func NewGinHandler(gin *gin.Engine, service service.BookServiceInterface) *ginHandler {

	return &ginHandler{gin, service}
}

func (h *ginHandler) GetListBooksHandler() {

	h.engine.GET("/books", func(c *gin.Context) {
		//paramsTitle := c.Query("title")

		books, err := h.service.GetListBooks()

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Error",
			})

		}

		if len(books) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Data not found",
				"status":  "NOT_FOUND",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "Data found",
			"data":   books,
		})
	})
}

func (h *ginHandler) RegisterRoutes() {

	h.GetListBooksHandler()
	h.GetDetailBookByIdHandler()
}

func (h *ginHandler) GetDetailBookByIdHandler() {

	h.engine.GET("/books/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		book, _ := h.service.GetBookDetailId(id)

		c.JSON(http.StatusOK, gin.H{
			"status": "Data found",
			"data":   book,
		})
	})

}

func (h *ginHandler) DeleteBook(c *gin.Context) {

	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"title": "Delete  book by id " + id,
		"id":    id,
	})
}

func (h *ginHandler) RegisterNewBookHandler(c *gin.Context) {

	var newBook dto.BookDto

	err := c.ShouldBindJSON(&newBook)

	if err != nil {

		errorMessages := []string{}

		var ve validator.ValidationErrors

		if errors.As(err, &ve) {
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := e.Field() + " " + e.ActualTag()
				errorMessages = append(errorMessages, errorMessage)

			}

			// handle validator error
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"errors":  errorMessages,
		})

	} else {

		c.JSON(http.StatusOK, gin.H{
			"status":    "success created the boook",
			"title":     newBook.Title,
			"price":     newBook.Price,
			"sub_title": newBook.SubTitle,
		})
	}

}
