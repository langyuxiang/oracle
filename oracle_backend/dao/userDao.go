package dao

import (
	"Oracle/model"
	"errors"
	"github.com/sirupsen/logrus"
)

func Register(name string, email string, pwd string) error {

	user := model.Users{}
	err := DB.Table("users").Where("email = ?", email).First(&user).Error

	if user.Email != "" {
		return errors.New("该邮箱已注册")
	}

	user.Name = name
	user.Email = email
	user.Password = pwd
	err = DB.Table("users").Create(&user).Error
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func Login(email string, pwd string) error {

	user := model.Users{}
	DB.Table("users").Where("email = ?", email).First(&user)
	if user.Password == pwd {
		return nil
	}

	return errors.New("未找到用户，或密码错误！")
}
