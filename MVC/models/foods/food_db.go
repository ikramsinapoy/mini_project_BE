package foods

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	Id          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Name        string         `json:"name"`
	Price       string         `json:"price"`
	Description string         `json:"description"`
	Category    string         `json:"category"`
	Restaurant  string         `json:"restaurant"`
	Dob         string         `json:"dob"`
}
