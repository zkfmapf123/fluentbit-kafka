package models

type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null;size:100"`
	Age  int    `gorm:"not null"`
}
