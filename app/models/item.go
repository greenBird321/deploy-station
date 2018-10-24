// http://go-database-sql.org/retrieving.html
// https://github.com/go-sql-driver/mysql/wiki/Examples
package models

import (
	"DeployStation/app/service"
	"log"
	"time"
	"strings"
	"fmt"
)

func ListItem() ([]map[string]string, error) {
	sql := "SELECT name, repo_url, repo_type, remark, ctime, mtime FROM items WHERE 1=1"
	rows, err := service.Db().Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []map[string]string
	for rows.Next() {
		var name, repo_url, repo_type, remark string
		var ctime, mtime int64
		err := rows.Scan(&name, &repo_url, &repo_type, &remark, &ctime, &mtime)
		if err != nil {
			log.Fatal(err)
		}
		r := make(map[string]string)
		r["name"] = name
		r["repo_url"] = repo_url
		r["repo_type"] = repo_type
		r["remark"] = remark
		r["ctime"] = time.Unix(ctime, 0).Format("2006-01-02 15:04")
		r["mtime"] = time.Unix(mtime, 0).Format("2006-01-02 15:04")
		data = append(data, r)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return data, nil
}

func GetItem(name string) (map[string]string, error) {
	var repo_url, repo_type, repo_private_key, remark, notify string
	var id, ctime, mtime int64

	sql := "SELECT id, name, repo_url, repo_type, repo_private_key, remark, notify, ctime, mtime FROM items WHERE name=? LIMIT 1"
	row := service.Db().QueryRow(sql, name)
	err := row.Scan(&id, &name, &repo_url, &repo_type, &repo_private_key, &remark, &notify, &ctime, &mtime)
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	data["id"] = fmt.Sprintf("%d", id)
	data["name"] = name
	data["repo_url"] = repo_url
	data["repo_type"] = repo_type
	data["repo_private_key"] = repo_private_key
	data["remark"] = remark
	data["notify"] = notify
	data["ctime"] = time.Unix(ctime, 0).Format("2006-01-02 15:04")
	data["mtime"] = time.Unix(mtime, 0).Format("2006-01-02 15:04")

	return data, nil
}

func UpdateItem(name string, data map[string]string) error {
	var update string
	for k, v := range data {
		update += "`" + k + "`='" + v + "',"
	}
	update = strings.Trim(update, ",")
	sql := "UPDATE items SET " + update + " WHERE `name`='" + name + "'"
	_, err := service.Db().Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
