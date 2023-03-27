package infrastructure

// import (
//     "gorm.io/driver/mysql"
//     "gorm.io/gorm"
// )

// func NewDB() (*gorm.DB, error) {
//     dsn := "user:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local"
//     return gorm.Open(mysql.Open(dsn), &gorm.Config{})
// }


import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "../domain"
)

func initDB() *gorm.DB {
    dsn := "host=localhost user=postgres password=postgres dbname=books port=5432 sslmode=disable TimeZone=Asia/Shanghai"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    // 创建books表
    db.AutoMigrate(&User{})

    return db
}
