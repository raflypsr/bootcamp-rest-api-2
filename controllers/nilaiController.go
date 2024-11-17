package controllers

import (
	"net/http"
	"time"

	"backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type nilaiInput struct {
	Indeks       string `binding:"required" json:"indeks"`
	Skor         int    `binding:"required" json:"skor"`
	MahasiswaID  uint   `binding:"required" json:"mahasiswa_id"`
	MataKuliahID uint   `binding:"required" json:"mata_kuliah_id"`
	UserID       uint   `binding:"required" json:"user_id"`
}

// GetAllNilai godoc
// @Summary Get all nilai.
// @Description Get a list of nilai.
// @Tags nilai
// @Produce json
// @Success 200 {object} []models.Nilai
// @Router /nilai [get]
func GetAllNilai(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var nilai []models.Nilai
	db.Find(&nilai)

	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// CreateNilai godoc
// @Summary Create New nilai.
// @Description Creating a new nilai.
// @Tags nilai
// @Param Body body nilaiInput true "the body to create a new nilai"
// @Produce json
// @Success 201 {object} models.Nilai
// @Router /nilai [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateNilai(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input nilaiInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nilai := models.Nilai{
		Indeks:       input.Indeks,
		Skor:         input.Skor,
		MahasiswaID:  input.MahasiswaID,
		MataKuliahID: input.MataKuliahID,
		UserID:       input.UserID,
	}
	db.Create(&nilai)

	c.JSON(http.StatusCreated, gin.H{"data": nilai})
}

// GetNilaiById godoc
// @Summary Get nilai.
// @Description Get a nilai by id.
// @Tags nilai
// @Produce json
// @Param id path string true "nilai id"
// @Success 200 {object} models.Nilai
// @Router /nilai/{id} [get]
func GetNilaiById(c *gin.Context) {
	var nilai models.Nilai

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&nilai).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// UpdateNilai godoc
// @Summary Update nilai.
// @Description Update nilai by id.
// @Tags nilai
// @Produce json
// @Param id path string true "nilai id"
// @Param Body body nilaiInput true "the body to update an nilai"
// @Success 200 {object} models.Nilai
// @Router /nilai/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateNilai(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var nilai models.Nilai
	if err := db.Where("id = ?", c.Param("id")).First(&nilai).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input nilaiInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Nilai
	updatedInput.Indeks = input.Indeks
	updatedInput.Skor = input.Skor
	updatedInput.MataKuliahID = input.MataKuliahID
	updatedInput.MahasiswaID = input.MahasiswaID
	updatedInput.UserID = input.UserID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&nilai).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// DeleteNilai godoc
// @Summary Delete one nilai.
// @Description Delete a nilai by id.
// @Tags nilai
// @Produce json
// @Param id path string true "nilai id"
// @Success 200 {object} map[string]boolean
// @Router /nilai/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteNilai(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var nilai models.Nilai
	if err := db.Where("id = ?", c.Param("id")).First(&nilai).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&nilai)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
