package service

import (
	"os/exec"
	"bytes"
	"log"
)

// @see https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
func ExecuteCommand(statement string) (string, string) {
	cmd := exec.Command("bash", "-c", statement)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	return outStr, errStr
}
