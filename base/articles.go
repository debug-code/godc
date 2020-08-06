package base

import (
	"fmt"
	"godc/db"
	"godc/models"
)

// Article find article by uid
func Article(uid string) (models.Article, error) {
	users := models.Article{}
	query := models.Article{Uid: uid}
	mdb := db.Mariadb.Where(&query).First(&users)
	if mdb.Error != nil {
		fmt.Println(mdb.Error.Error())
		return users, mdb.Error
	}

	return users, nil
}
