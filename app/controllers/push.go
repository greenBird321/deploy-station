package controllers

import (
	"github.com/gin-gonic/gin"
	"strings"
	"deploy-station/app/service"
	"net/http"
	"deploy-station/app/models"
	"log"
)

func PreparePush(c *gin.Context) {
	name := c.Param("name")
	project, _ := models.GetItem(name)

	workspace := "$HOME/.deploy/" + name
	cmd := `
cd $WORKSPACE
git rev-parse --symbolic --tags
`
	cmd = strings.Replace(cmd, "$WORKSPACE", workspace, -1)
	outStr, _ := service.ExecuteCommand(cmd)
	outStr = strings.Trim(outStr, "\n")
	tags := strings.Split(outStr, "\n")
	c.HTML(http.StatusOK, "deploy/prepare.html", gin.H{
		"project": project,
		"tags":    tags,
	})
}

func ProcessPush(c *gin.Context) {
	syncDir := "www"
	syncPort := "873"

	id := c.PostForm("id")
	name := c.PostForm("name")
	tag := c.PostForm("tag")
	// remark := c.PostForm("remark")

	workspace := "$HOME/.deploy/" + name
	statement := `
cd $WORKSPACE
git checkout $TAG
`
	statement = strings.Replace(statement, "$WORKSPACE", workspace, -1)
	statement = strings.Replace(statement, "$TAG", tag, -1)
	service.ExecuteCommand(statement)

	nodes, err := models.GetNodes(id)
	if err != nil {
		log.Println("no node available")
	}

	// 同步 rsync -az --exclude='.git*' --timeout=30 --port=873 /data/project 192.168.1.2::www
	for _, node := range nodes {
		statement := `
cd $WORKSPACE
$SYNC_CMD
`
		rsyncCmd := "rsync -az --delete --exclude='.git*' --timeout=30 --port=" + syncPort + " . " + node["ip_intranet"] + "::" + syncDir + "/" + name
		statement = strings.Replace(statement, "$WORKSPACE", workspace, -1)
		statement = strings.Replace(statement, "$SYNC_CMD", rsyncCmd, -1)
		go func() {
			service.ExecuteCommand(statement)
		}()
	}
	c.HTML(http.StatusOK, "deploy/result.html", gin.H{
	})
}
