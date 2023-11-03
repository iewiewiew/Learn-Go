package service

import (
	"errors"
	"github.com/wonderivan/logger"
	"k8s-platform/config"
)

/**
 * @author       weimenghua
 * @time         2023-11-04 16:43
 * @description
 */

var Login login

type login struct{}

//验证账号密码
func (l *login) Auth(username, password string) (err error) {
	if username == config.AdminUser && password == config.AdminPwd {
		return nil
	} else {
		logger.Error("登录失败, 用户名或者密码错误")
		return errors.New("登录失败, 用户名或密码错误")
	}
}
