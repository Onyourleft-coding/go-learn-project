package services

import (
	"errors"
	"fmt"
	"todo_list/app/common/request"
	"todo_list/global"
	"todo_list/models"
	"todo_list/utils"
)

//用户注册逻辑

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password))}
	fmt.Println("user", user)
	err = global.App.DB.Create(&user).Error
	return
}
