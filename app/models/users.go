package models

import (
	"deploy-station/app/service"
	"github.com/pkg/errors"
)

// 通过账号寻找用户
func GetUserByName(username string) (password string, uid int16, err error) {
	err = service.Db().QueryRow("SELECT password, id FROM `users` WHERE username = ?", username).Scan(&password, &uid)
	if err != nil {
		err = errors.New("user not exist")
		return
	}
	return
}
