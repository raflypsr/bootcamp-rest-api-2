package controllers

import (
	"net/http"
	"time"

	"backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type mahasiswaInput struct {
	Nama string `binding:"required" json:"nama"`
}

// GetAllMahasiswa godoc
// @Summary Get all Mahasiswa.
// @Description Get a list of mahasiswa.
// @Tags mahasiswa
// @Produce json
// @Success 200 {object} []models.Mahasiswa
// @Router /mahasiswa [get]
func GetAllMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa []models.Mahasiswa
	db.Preload("JadwalKuliah").Preload("Nilai").Find(&mahasiswa)

	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

// CreateMahasiswa godoc
// @Summary Create New mahasiswa.
// @Description Creating a new mahasiswa.
// @Tags mahasiswa
// @Param Body body mahasiswaInput true "the body to create a new mahasiswa"
// @Produce json
// @Success 201 {object} models.Mahasiswa
// @Router /mahasiswa [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input mahasiswaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mahasiswa := models.Mahasiswa{
		Nama: input.Nama,
	}
	db.Create(&mahasiswa)

	c.JSON(http.StatusCreated, gin.H{"data": mahasiswa})
}

// GetMahasiswaById godoc
// @Summary Get mahasiswa.
// @Description Get a mahasiswa by id.
// @Tags mahasiswa
// @Produce json
// @Param id path string true "mahasiswa id"
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa/{id} [get]
func GetMahasiswaById(c *gin.Context) {
	var mahasiswa models.Mahasiswa

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Preload("JadwalKuliah").Preload("Nilai").First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

// UpdateMahasiswa godoc
// @Summary Update mahasiswa.
// @Description Update mahasiswa by id.
// @Tags mahasiswa
// @Produce json
// @Param id path string true "mahasiswa id"
// @Param Body body mahasiswaInput true "the body to update an mahasiswa"
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateMahasiswa(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var mahasiswa models.Mahasiswa
	if err := db.Where("id = ?", c.Param("id")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input mahasiswaInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Mahasiswa
	updatedInput.Nama = input.Nama
	updatedInput.UpdatedAt = time.Now()

	db.Model(&mahasiswa).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

// DeleteMahasiswa godoc
// @Summary Delete one mahasiswa.
// @Description Delete a mahasiswa by id.
// @Tags mahasiswa
// @Produce json
// @Param id path string true "mahasiswa id"
// @Success 200 {object} map[string]boolean
// @Router /mahasiswa/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteMahasiswa(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var mahasiswa models.Mahasiswa
	if err := db.Where("id = ?", c.Param("id")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&mahasiswa)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
