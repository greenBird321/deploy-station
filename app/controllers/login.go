package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"deploy-station/app/models"
	"fmt"
	"crypto/md5"
	"deploy-station/app/env"
	"strconv"
)

func Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		username := c.DefaultPostForm("username", "")
		password := c.DefaultPostForm("password", "")
		// 密码 dev123
		data := []byte(password)
		has := md5.Sum(data)
		md5Pass := fmt.Sprintf("%x", has)

		pass, uid, err := models.GetUserByName(username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": fmt.Sprintf("getUser err: %s", err)})
			return
		}

		if md5Pass != pass {
			c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "password error please input again"})
			return
		}

		ENV, _ := env.AppEnv()

		c.SetCookie("uid", strconv.Itoa(int(uid)), 2592000, "/", ENV.Domain, false, true)
		c.Redirect(http.StatusMovedPermanently, "/home")
		return
	}

	c.HTML(http.StatusOK, "login/index.html", gin.H{
	})
}
