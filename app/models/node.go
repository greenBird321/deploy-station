package models

import (
	"log"
	"deploy-station/app/service"
)

func GetNodes(itemId string) ([]map[string]string, error) {
	sql := "SELECT ip, ip_intranet FROM `nodes` n, items_nodes i WHERE n.id=i.node_id AND i.item_id=?"
	rows, err := service.Db().Query(sql, itemId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []map[string]string
	for rows.Next() {
		var ip, ip_intranet string
		err := rows.Scan(&ip, &ip_intranet)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		r := make(map[string]string)
		r["ip"] = ip
		r["ip_intranet"] = ip_intranet
		data = append(data, r)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return data, nil
}
