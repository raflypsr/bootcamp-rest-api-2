package models

import (
	"time"
)

type (
	Nilai struct {
		ID           uint   `gorm:"primary_key" json:"id"`
		Indeks       string `gorm:"type:varchar(255); not null" json:"indeks"`
		Skor         int    `gorm:"type:int; not null" json:"skor"`
		CreatedAt    time.Time
		UpdatedAt    time.Time
		MahasiswaID  uint
		MataKuliahID uint
		UserID       uint
	}
)

func (Nilai) TableName() string {
	return "nilai"
}
