package controllers

import (
	"net/http"
	"regexp"
	"time"

	"backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookInput struct {
	Title       string `binding:"required" json:"title"`
	Description string `binding:"required" json:"description"`
	ImageUrl    string `binding:"required,url" json:"image_url"`
	ReleaseYear int    `binding:"required,min=1980,max=2021" json:"release_year"`
	Price       string `binding:"required" json:"price"`
	TotalPage   int    `binding:"required" json:"total_page"`
}

// GetAllBook godoc
// @Summary Get all books.
// @Description Get a list of books.
// @Tags book
// @Produce json
// @Success 200 {object} []models.Books
// @Router /books [get]
func GetAllBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Books
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook godoc
// @Summary Create New book.
// @Description Creating a new book.
// @Tags book
// @Param Body body bookInput true "the body to create a new book"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Books
// @Router /books [post]
func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input bookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var validationErrors []string

	re := regexp.MustCompile(`^https?://.+`)
	if !re.MatchString(input.ImageUrl) {
		validationErrors = append(validationErrors, "Image URL is not valid")
	}

	if input.ReleaseYear < 1980 || input.ReleaseYear > 2021 {
		validationErrors = append(validationErrors, "Release year must be between 1980 and 2021")
	}

	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	var thickness string
	if input.TotalPage <= 100 {
		thickness = "tipis"
	} else if input.TotalPage >= 101 && input.TotalPage <= 200 {
		thickness = "sedang"
	} else if input.TotalPage >= 201 {
		thickness = "tebal"
	}

	// Create book
	book := models.Books{
		Title:       input.Title,
		ReleaseYear: input.ReleaseYear,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   thickness,
	}
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GetBookById godoc
// @Summary Get book.
// @Description Get a book by id.
// @Tags book
// @Produce json
// @Param id path string true "book id"
// @Success 200 {object} models.Books
// @Router /books/{id} [get]
func GetBookById(c *gin.Context) {
	var book models.Books

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary Update book.
// @Description Update book by id.
// @Tags book
// @Produce json
// @Param id path string true "book id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param Body body bookInput true "the body to update an book"
// @Success 200 {object} models.Books
// @Router /books/{id} [patch]
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var book models.Books
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input bookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var validationErrors []string

	re := regexp.MustCompile(`^https?://.+`)
	if !re.MatchString(input.ImageUrl) {
		validationErrors = append(validationErrors, "Image URL is not valid")
	}

	if input.ReleaseYear < 1980 || input.ReleaseYear > 2021 {
		validationErrors = append(validationErrors, "Release year must be between 1980 and 2021")
	}

	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	var thickness string
	if input.TotalPage <= 100 {
		thickness = "tipis"
	} else if input.TotalPage >= 101 && input.TotalPage <= 200 {
		thickness = "sedang"
	} else if input.TotalPage >= 201 {
		thickness = "tebal"
	}

	var updatedInput models.Books
	updatedInput.Title = input.Title
	updatedInput.ReleaseYear = input.ReleaseYear
	updatedInput.Description = input.Description
	updatedInput.ImageUrl = input.ImageUrl
	updatedInput.TotalPage = input.TotalPage
	updatedInput.Price = input.Price
	updatedInput.Thickness = thickness
	updatedInput.UpdatedAt = time.Now()

	db.Model(&book).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary Delete one book.
// @Description Delete a book by id.
// @Tags book
// @Produce json
// @Param id path string true "book id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Books
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
