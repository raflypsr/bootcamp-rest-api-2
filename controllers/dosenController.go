package controllers

import (
	"net/http"
	"time"

	"backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type dosenInput struct {
	Nama         string `binding:"required" json:"nama"`
	MataKuliahID uint
}

// GetAllDosen godoc
// @Summary Get all Dosen.
// @Description Get a list of dosen.
// @Tags dosen
// @Produce json
// @Success 200 {object} []models.Dosen
// @Router /dosen [get]
func GetAllDosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var dosen []models.Dosen
	db.Preload("JadwalKuliah").Find(&dosen)

	c.JSON(http.StatusOK, gin.H{"data": dosen})
}

// CreateDosen godoc
// @Summary Create New dosen.
// @Description Creating a new dosen.
// @Tags dosen
// @Param Body body dosenInput true "the body to create a new dosen"
// @Produce json
// @Success 201 {object} models.Dosen
// @Router /dosen [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateDosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input dosenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dosen := models.Dosen{
		Nama:         input.Nama,
		MataKuliahID: input.MataKuliahID,
	}
	db.Create(&dosen)

	c.JSON(http.StatusCreated, gin.H{"data": dosen})
}

// GetDosenById godoc
// @Summary Get dosen.
// @Description Get a dosen by id.
// @Tags dosen
// @Produce json
// @Param id path string true "dosen id"
// @Success 200 {object} models.Dosen
// @Router /dosen/{id} [get]
func GetDosenById(c *gin.Context) {
	var dosen models.Dosen

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).Preload("JadwalKuliah").First(&dosen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dosen})
}

// UpdateDosen godoc
// @Summary Update dosen.
// @Description Update dosen by id.
// @Tags dosen
// @Produce json
// @Param id path string true "dosen id"
// @Param Body body dosenInput true "the body to update an dosen"
// @Success 200 {object} models.Dosen
// @Router /dosen/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateDosen(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var dosen models.Dosen
	if err := db.Where("id = ?", c.Param("id")).First(&dosen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input dosenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Dosen
	updatedInput.Nama = input.Nama
	updatedInput.MataKuliahID = input.MataKuliahID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&dosen).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": dosen})
}

// DeleteDosen godoc
// @Summary Delete one dosen.
// @Description Delete a dosen by id.
// @Tags dosen
// @Produce json
// @Param id path string true "dosen id"
// @Success 200 {object} map[string]boolean
// @Router /dosen/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteDosen(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var dosen models.Dosen
	if err := db.Where("id = ?", c.Param("id")).First(&dosen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&dosen)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
