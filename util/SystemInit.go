package util

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

}

func InitMysql() {
	//DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), *gorm.Config{})
}
