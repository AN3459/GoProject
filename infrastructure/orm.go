package infrastructure

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
    dsn := "user:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local"
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
