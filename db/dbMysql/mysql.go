package dbMysql

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Regist() (*gorm.DB, error) {
	db, err := gorm.Open("mysql",
		"alex:123qwe=-0@(localhost:33306)/godc?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		fmt.Println("connection succedssed")

	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	//defer db.Close()
	return db, nil

}
