package model

import "errors"

//自定义错误
var (
	ERROR_USER_NOTEXIST = errors.New("用户不存在")
	ERROR_USER_EXIST    = errors.New("用户已经存在")
	ERROR_USER_PWD      = errors.New("密码不正确")
)
