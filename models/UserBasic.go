package models

import (
	"ginchat/util"
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogOut      bool
	DeviceInfo    string
}

func TableName(table *UserBasic) string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	return data
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	util.DB.Where("name = ? and pass_word = ?", name, password).First(&user)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	util.DB.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(user UserBasic) *gorm.DB {
	return util.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return util.DB.Delete(user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	return util.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord})
}
