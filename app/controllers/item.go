package controllers

import (
	"github.com/gin-gonic/gin"
	"DeployStation/app/service"
	"net/http"
	"log"
	"time"
	"DeployStation/app/models"
	"fmt"
)

func ListItem(c *gin.Context) {
	data, err := models.ListItem()
	if err != nil {
		log.Fatal(err)
	}
	c.HTML(http.StatusOK, "item/list.html", gin.H{
		"items": data,
	})
}

func ShowItem(c *gin.Context) {
	name := c.Param("name")
	data := models.GetItem(name)
	c.HTML(http.StatusOK, "item/show.html", gin.H{
		"item": data,
	})
}

func SettingItem(c *gin.Context) {
	var name string
	var isUpdate string
	name = c.Param("name")

	newName := c.PostForm("name")
	if newName != "" {
		update := make(map[string]string)
		update["name"] = newName
		update["remark"] = c.PostForm("remark")
		//update["color"] = c.PostForm("color")
		update["repo_url"] = c.PostForm("repo_url")
		update["repo_private_key"] = c.PostForm("repo_private_key")
		update["notify"] = c.PostForm("notify")
		update["mtime"] = fmt.Sprintf("%d", time.Now().Unix())

		// 更新
		err := models.UpdateItem(name, update)
		if err == nil {
			isUpdate = "failed"
			name = newName
		}
		isUpdate = "success"
	}

	data := models.GetItem(name)
	c.HTML(http.StatusOK, "item/setting.html", gin.H{
		"item":   data,
		"update": isUpdate,
	})
}
