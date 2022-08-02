package services

import (
	"errors"
	"gin-demo/app/common/request"
	"gin-demo/app/models"
	"gin-demo/global"
	"gin-demo/utils"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	result := global.App.DB.Where("mobile=?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{
		Name:     params.Name,
		Mobile:   params.Mobile,
		Password: utils.BcryptMake([]byte(params.Password)),
	}
	err = global.App.DB.Create(&user).Error
	return
}
