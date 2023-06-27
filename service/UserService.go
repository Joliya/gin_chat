package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/util"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// GetUserList
// @Tags 用户列表
// @success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUser
// @Tags 创建用户
// @success 200 {string} json{"code", "message"}
// @Router /user/getUserList [POST]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	password := c.Query("password")
	repassword := c.Query("repassword")

	salt := fmt.Sprintf("%06d", rand.Int31())

	if password != repassword {
		c.JSON(-1, gin.H{
			"message": "两次密码不一样",
		})
	}
	//user.PassWord = password
	user.PassWord = util.MakePassword(password, salt)
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增成功",
	})
}

// DeleteUser
// @Tags 删除用户
// @success 200 {string} json{"code", "message"}
// @Router /user/getUserList [POST]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}

// UpdateUser
// @Tags 更新用户
// @success 200 {string} json{"code", "message"}
// @Router /user/updateUser [POST]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "更新成功",
	})
}

// FindUserByNameAndPwd
// @Tags 查找用户
// @success 200 {string} json{"code", "message"}
// @Router /user/getUserList [get]
func FindUserByNameAndPwd(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")

	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(401, gin.H{
			"message": "用户不存在",
		})
		return
	}
	flat := util.ValidPassword(password, user.Salt, user.PassWord)
	if flat == false {
		c.JSON(401, gin.H{
			"message": "用户名或密码错误",
		})
		return
	}

	user = models.FindUserByNameAndPwd(name, util.MakePassword(password, user.Salt))
	c.JSON(200, gin.H{
		"message": user,
	})
	return
}
