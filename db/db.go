package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"godc/db/dbMysql"
)

var Mariadb *gorm.DB

func DbInit() {
	var err error
	Mariadb, err = dbMysql.Regist()
	if err != nil {
		fmt.Println("mariadb init error : ", err)
	}

}
