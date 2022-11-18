package main

import (
	dao "web_06/dao/mysql"
	"web_06/router"
)

func main() {
	err := dao.Initmysql()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	router.Initrouter()
}
