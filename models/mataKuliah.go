package models

import (
	"time"
)

type (
	MataKuliah struct {
		ID        uint   `gorm:"primary_key" json:"id"`
		Nama      string `gorm:"type:varchar(255); not null" json:"nama"`
		CreatedAt time.Time
		UpdatedAt time.Time
		Dosen     []Dosen
		Nilai     []Nilai
	}
)

func (MataKuliah) TableName() string {
	return "mata_kuliah"
}
