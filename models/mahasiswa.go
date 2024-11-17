package models

import (
	"time"
)

type (
	Mahasiswa struct {
		ID           uint   `gorm:"primary_key" json:"id"`
		Nama         string `gorm:"type:varchar(255); not null" json:"nama"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
		JadwalKuliah []JadwalKuliah
		Nilai        []Nilai
	}
)

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}
