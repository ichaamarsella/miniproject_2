package entities

import "time"

type Customer struct {
	ID        uint      `gorm:"primaryKey"`
	Firstname string    `gorm:"column:first_name"`
	Lastname  string    `gorm:"column:last_name"`
	Email     string    `gorm:"column:email"`
	Avatar    string    `gorm:"column:avatar"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
