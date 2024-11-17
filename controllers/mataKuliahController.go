package controllers

import (
	"net/http"
	"time"

	"backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type mataKuliahInput struct {
	Nama string `binding:"required" json:"nama"`
}

// GetAllMataKuliah godoc
// @Summary Get all mataKuliah.
// @Description Get a list of mataKuliah.
// @Tags mata kuliah
// @Produce json
// @Success 200 {object} []models.MataKuliah
// @Router /mata-kuliah [get]
func GetAllMataKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mataKuliah []models.MataKuliah
	db.Preload("Dosen").Preload("Nilai").Find(&mataKuliah)

	c.JSON(http.StatusOK, gin.H{"data": mataKuliah})
}

// CreateMataKuliah godoc
// @Summary Create New mataKuliah.
// @Description Creating a new mataKuliah.
// @Tags mata kuliah
// @Param Body body mataKuliahInput true "the body to create a new mataKuliah"
// @Produce json
// @Success 201 {object} models.MataKuliah
// @Router /mata-kuliah [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateMataKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input mataKuliahInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mataKuliah := models.MataKuliah{
		Nama: input.Nama,
	}
	db.Create(&mataKuliah)

	c.JSON(http.StatusCreated, gin.H{"data": mataKuliah})
}

// GetMataKuliahById godoc
// @Summary Get mataKuliah.
// @Description Get a mataKuliah by id.
// @Tags mata kuliah
// @Produce json
// @Param id path string true "mataKuliah id"
// @Success 200 {object} models.MataKuliah
// @Router /mata-kuliah/{id} [get]
func GetMataKuliahById(c *gin.Context) {
	var mataKuliah models.MataKuliah

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Preload("Dosen").Preload("Nilai").First(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mataKuliah})
}

// UpdateMataKuliah godoc
// @Summary Update mataKuliah.
// @Description Update mataKuliah by id.
// @Tags mata kuliah
// @Produce json
// @Param id path string true "mataKuliah id"
// @Param Body body mataKuliahInput true "the body to update an mataKuliah"
// @Success 200 {object} models.MataKuliah
// @Router /mata-kuliah/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateMataKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var mataKuliah models.MataKuliah
	if err := db.Where("id = ?", c.Param("id")).First(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input mataKuliahInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.MataKuliah
	updatedInput.Nama = input.Nama
	updatedInput.UpdatedAt = time.Now()

	db.Model(&mataKuliah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": mataKuliah})
}

// DeleteMataKuliah godoc
// @Summary Delete one mataKuliah.
// @Description Delete a mataKuliah by id.
// @Tags mata kuliah
// @Produce json
// @Param id path string true "mataKuliah id"
// @Success 200 {object} map[string]boolean
// @Router /mata-kuliah/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteMataKuliah(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var mataKuliah models.MataKuliah
	if err := db.Where("id = ?", c.Param("id")).First(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&mataKuliah)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
