package domain

type User struct {
    ID   int64  `gorm:"primaryKey"`
    Name string `gorm:"not null"`
    Age  int    `gorm:"not null"`
}
