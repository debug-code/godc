package service

import (
	"godc/base"
	"godc/models"
)

// ArticlesOne one
func ArticlesOne(uid string) (models.Article, error) {

	article, err := base.ArticlesOne(uid)
	if err != nil {
		return article, err
	}
	return article, nil

}

// ArticlesAdd add article
func ArticlesAdd(articleDto models.Article) error {
	err := base.ArticlesAdd(articleDto)
	if err != nil {
		return err
	}
	return nil

}
