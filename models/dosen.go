package models

import (
	"time"
)

type (
	Dosen struct {
		ID           uint   `gorm:"primary_key" json:"id"`
		Nama         string `gorm:"type:varchar(255); not null" json:"nama"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
		MataKuliahID uint
		JadwalKuliah []JadwalKuliah
	}
)

func (Dosen) TableName() string {
	return "dosen"
}
