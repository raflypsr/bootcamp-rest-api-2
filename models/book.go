package models

import (
	"time"
)

type (
	Books struct {
		ID          uint   `gorm:"primary_key" json:"id"`
		Title       string `gorm:"type:varchar(255); not null" json:"title"`
		Description string `gorm:"type:text; not null" json:"description"`
		ImageUrl    string `gorm:"type:text; not null" json:"image_url"`
		ReleaseYear int    `gorm:"type:int; not null" json:"release_year"`
		Price       string `gorm:"type:varchar(255); not null" json:"price"`
		TotalPage   int    `gorm:"type:int; not null" json:"total_page"`
		Thickness   string `gorm:"type:varchar(255); not null" json:"thickness"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
