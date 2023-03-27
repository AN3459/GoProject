package domain

//import "gorm.io/gorm"

type User struct {
    ID   int64  `gorm:"primaryKey"`
    Name string `gorm:"not null"`
    Age  int    `gorm:"not null"`
}
