package dao

import (
	"fmt"

	"github.com/DuckBroApprentice/GoWeb/models"
	"github.com/DuckBroApprentice/GoWeb/tool"
)

type UserDao struct {
	*tool.Orm
}

func (ud *UserDao) ValidateSmsCode(username string, password string) *models.User {
	var user models.User

	if _, err := ud.Where("phone = ? and code = ?", username, password).Get(&user); err != nil {
		fmt.Println(err.Error())
	}
	return &user
}

func (ud *UserDao) Query(name ,password string)*models.User{
	var user models.User
	password = tool.EncoderSha256(password)

	_,err:= ud.Where("username = ? and passowrd = ?",name,password).Get(&user)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &user
}

// 新用戶的數據庫插入操作
func (ud *UserDao) InsertUsername(username models.User) int64 {
	result, err := ud.InsertOne(&username)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
func (ud *UserDao) InsertPassword(password models.User) int64 {
	result, err := ud.InsertOne(&password)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
func (ud *UserDao) InsertPhone(phone models.User) int64 {
	result, err := ud.InsertOne(&phone)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
func (ud *UserDao) InsertEmail(email models.User) int64 {
	result, err := ud.InsertOne(&email)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
