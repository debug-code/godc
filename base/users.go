package base

import (
	"fmt"
	"godc/db"
	"godc/models"
)

func Users() ([]models.User, error) {
	users := []models.User{}
	mdb := db.Mariadb.Find(&users)
	if mdb.Error != nil {
		fmt.Println(mdb.Error.Error())
		return nil, mdb.Error
	}

	return users, nil
}
