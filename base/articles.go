package base

import (
	"godc/db"
	"godc/models"
	"time"

	"debug-code.com/dclog"
)

// ArticleOne find article by uid
func ArticleOne(uid string) (models.Article, error) {

	article := models.Article{}
	query := models.Article{Uid: uid}
	dclog.Info(uid)
	mdb := db.Mariadb.Where(&query).First(&article)
	if mdb.Error != nil {
		dclog.Error(mdb.Error.Error())
		return article, mdb.Error
	}

	return article, nil
}

// Articles find some
func Articles(query models.Article, page models.Page) ([]models.Article, error) {

	articles := []models.Article{}
	mdb := db.Mariadb.Where(query).Offset(page.Skip).Limit(page.Limit).Find(&articles)
	if mdb.Error != nil {
		dclog.Error(mdb.Error.Error())
		return articles, mdb.Error
	}

	return articles, nil
}

// ArticlesCount find some
func ArticlesCount(query models.Article) (int, error) {

	total := 0
	mdb := db.Mariadb.Where(query).Find(&[]models.Article{}).Count(&total)
	if mdb.Error != nil {
		dclog.Error(mdb.Error.Error())
		return total, mdb.Error
	}

	return total, nil
}

// ArticlesAdd add
func ArticlesAdd(article models.Article) error {
	article.Uid = "uid"
	article.CreateTime = time.Now().Unix()
	article.UpdateTime = article.CreateTime
	mdb := db.Mariadb.Create(&article)
	if mdb.Error != nil {
		dclog.Error(mdb.Error.Error())
		return mdb.Error
	}

	return nil
}

// ArticlesDel add
func ArticlesDel(ids []string) error {

	mdb := db.Mariadb.Where("uid in (?) ", ids).Delete(&models.Article{})
	if mdb.Error != nil {
		dclog.Error(mdb.Error.Error())
		return mdb.Error
	}

	return nil
}

// ArticlesSetDel add
func ArticlesSetDel(ids []string) error {
	mdb := db.Mariadb.Model(models.Article{}).Where("uid IN (?)", ids).Updates(models.Article{Del: 1})
	if mdb.Error != nil {
		dclog.Error(mdb.Error.Error())
		return mdb.Error
	}
	return nil
}
