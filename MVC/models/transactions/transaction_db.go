package transactions

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	TotalPrice string         `json:"totalprice"`
	Status     string         `json:"Status"`
}
