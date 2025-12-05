package models

import (
	"time"

	"gorm.io/gorm"
)

// Expense representa un gasto en el sistema
type Expense struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Amount      float64        `gorm:"not null" json:"amount" binding:"required,gt=0"`
	Category    string         `gorm:"size:100;not null" json:"category" binding:"required"`
	Description string         `gorm:"size:500" json:"description"`
	Date        time.Time      `gorm:"not null;type:datetime" json:"date" binding:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// CreateExpenseRequest es el DTO para crear un gasto
type CreateExpenseRequest struct {
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	Category    string  `json:"category" binding:"required"`
	Description string  `json:"description"`
	Date        int64   `json:"date" binding:"required"`
}

// UpdateExpenseRequest es el DTO para actualizar un gasto
type UpdateExpenseRequest struct {
	Amount      *float64 `json:"amount" binding:"omitempty,gt=0"`
	Category    *string  `json:"category" binding:"omitempty"`
	Description *string  `json:"description"`
	Date        *int64   `json:"date"`
}

// CategoryStats representa estadísticas por categoría
type CategoryStats struct {
	Category string  `json:"category"`
	Total    float64 `json:"total"`
	Count    int64   `json:"count"`
}
