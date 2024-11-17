package models

import (
	"time"
)

type (
	JadwalKuliah struct {
		ID          uint      `gorm:"primary_key" json:"id"`
		Hari        string    `gorm:"type:varchar(255); not null" json:"hari"`
		JamMulai    time.Time `json:"jam_mulai"`
		JamSelesai  time.Time `json:"jam_selesai"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DosenID     uint
		MahasiswaID uint
	}
)

func (JadwalKuliah) TableName() string {
	return "jadwal_kuliah"
}
