package DB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Con *gorm.DB = InitDb()

func InitDb() *gorm.DB {
	dsn := "root:@/micro_auth?charset=utf8mb4&parseTime=True&loc=Local"
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
	}
	return con
}


func Migrate(){
	Con.AutoMigrate(&Token{})
	Con.AutoMigrate(&Post{})
	Con.AutoMigrate(&User{})
}