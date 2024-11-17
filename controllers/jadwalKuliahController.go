package controllers

import (
	"net/http"
	"time"

	"backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type jadwalKuliahInput struct {
	Hari        string    `binding:"required" json:"hari"`
	JamMulai    time.Time `binding:"required" json:"jam_mulai"`
	JamSelesai  time.Time `binding:"required" json:"jam_selesai"`
	DosenID     uint      `binding:"required" json:"dosen_id"`
	MahasiswaID uint      `binding:"required" json:"mahasiswa_id"`
}

// GetAllJadwalKuliah godoc
// @Summary Get all jadwalKuliah.
// @Description Get a list of jadwalKuliah.
// @Tags jadwal kuliah
// @Produce json
// @Success 200 {object} []models.JadwalKuliah
// @Router /jadwal-kuliah [get]
func GetAllJadwalKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var jadwalKuliah []models.JadwalKuliah
	db.Find(&jadwalKuliah)

	c.JSON(http.StatusOK, gin.H{"data": jadwalKuliah})
}

// CreateJadwalKuliah godoc
// @Summary Create New jadwalKuliah.
// @Description Creating a new jadwalKuliah.
// @Tags jadwal kuliah
// @Param Body body jadwalKuliahInput true "the body to create a new jadwalKuliah"
// @Produce json
// @Success 201 {object} models.JadwalKuliah
// @Router /jadwal-kuliah [post]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func CreateJadwalKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// Validate input
	var input jadwalKuliahInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jadwalKuliah := models.JadwalKuliah{
		Hari:        input.Hari,
		JamMulai:    input.JamMulai,
		JamSelesai:  input.JamSelesai,
		DosenID:     input.DosenID,
		MahasiswaID: input.MahasiswaID,
	}
	db.Create(&jadwalKuliah)

	c.JSON(http.StatusCreated, gin.H{"data": jadwalKuliah})
}

// GetJadwalKuliahById godoc
// @Summary Get jadwalKuliah.
// @Description Get a jadwalKuliah by id.
// @Tags jadwal kuliah
// @Produce json
// @Param id path string true "jadwalKuliah id"
// @Success 200 {object} models.JadwalKuliah
// @Router /jadwal-kuliah/{id} [get]
func GetJadwalKuliahById(c *gin.Context) {
	var jadwalKuliah models.JadwalKuliah

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": jadwalKuliah})
}

// UpdateJadwalKuliah godoc
// @Summary Update jadwalKuliah.
// @Description Update jadwalKuliah by id.
// @Tags jadwal kuliah
// @Produce json
// @Param id path string true "jadwalKuliah id"
// @Param Body body jadwalKuliahInput true "the body to update an jadwalKuliah"
// @Success 200 {object} models.JadwalKuliah
// @Router /jadwal-kuliah/{id} [patch]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func UpdateJadwalKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var jadwalKuliah models.JadwalKuliah
	if err := db.Where("id = ?", c.Param("id")).First(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input jadwalKuliahInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.JadwalKuliah
	updatedInput.Hari = input.Hari
	updatedInput.JamMulai = input.JamMulai
	updatedInput.JamSelesai = input.JamSelesai
	updatedInput.MahasiswaID = input.MahasiswaID
	updatedInput.DosenID = input.DosenID
	updatedInput.UpdatedAt = time.Now()

	db.Model(&jadwalKuliah).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": jadwalKuliah})
}

// DeleteJadwalKuliah godoc
// @Summary Delete one jadwalKuliah.
// @Description Delete a jadwalKuliah by id.
// @Tags jadwal kuliah
// @Produce json
// @Param id path string true "jadwalKuliah id"
// @Success 200 {object} map[string]boolean
// @Router /jadwal-kuliah/{id} [delete]
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
func DeleteJadwalKuliah(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var jadwalKuliah models.JadwalKuliah
	if err := db.Where("id = ?", c.Param("id")).First(&jadwalKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&jadwalKuliah)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
