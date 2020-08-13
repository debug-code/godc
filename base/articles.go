package base

import (
	"fmt"
	"godc/db"
	"godc/models"
	"time"
)

// Article find article by uid
func ArticlesOne(uid string) (models.Article, error) {

	article := models.Article{}
	query := models.Article{Uid: uid}
	mdb := db.Mariadb.Where(&query).First(&article)
	if mdb.Error != nil {
		fmt.Println(mdb.Error.Error())
		return article, mdb.Error
	}

	return article, nil
}

// ArticlesAdd add
func ArticlesAdd(article models.Article) error {
	article.Uid = "uid"
	article.CreateTime = time.Now().Unix()
	article.UpdateTime = article.CreateTime
	mdb := db.Mariadb.Create(&article)
	if mdb.Error != nil {
		fmt.Println(mdb.Error.Error())
		return mdb.Error
	}

	return nil
}
