package test

import (
	"fmt"
	"ginchat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Product struct {
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3307)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{Logger: newLogger})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.UserBasic{})

	user := &models.UserBasic{
		Name: "张晋鹏",
	}
	db.Create(user)

	fmt.Println(db.First(user, 1))

	db.Model(user).Update("PassWord", "123456")
}
