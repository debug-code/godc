package service

import (
	"godc/base"
	"godc/models"

	"debug-code.com/dclog"
)

// ArticleOne one
func ArticleOne(uid string) (models.Article, error) {

	article, err := base.ArticleOne(uid)
	if err != nil {
		dclog.Error(err)
		return article, err
	}
	return article, nil

}

// Articles one
func Articles(query models.Article, page models.Page) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	articles, err := base.Articles(query, page)
	if err != nil {
		dclog.Error(err)
		return result, err
	}
	total, err := base.ArticlesCount(query)
	if err != nil {
		dclog.Error(err)
		return result, err
	}
	result["list"] = articles
	result["total"] = total

	return result, nil

}

// ArticlesAdd add article
func ArticlesAdd(articleDto models.Article) error {
	err := base.ArticlesAdd(articleDto)
	if err != nil {
		return err
	}
	return nil

}

// ArticlesDel add article
func ArticlesDel(ids []string, flag bool) error {
	if flag {
		err := base.ArticlesDel(ids)
		if err != nil {
			return err
		}
	}

	err := base.ArticlesSetDel(ids)
	if err != nil {
		return err
	}
	return nil

}
