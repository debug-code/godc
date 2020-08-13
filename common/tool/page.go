package tool

import (
	"godc/models"
	"strconv"
)

func PageSize(pageNum, pageSize string) models.Page {
	page := models.Page{PageNum: 1, PageSize: 10}
	if pageNum != "" {
		page.PageNum, _ = strconv.Atoi(pageNum)
	}
	if pageSize != "" {
		page.PageNum, _ = strconv.Atoi(pageSize)
	}
	page.Skip = page.PageSize * (page.PageNum - 1)
	page.Limit = page.PageSize
	return page

}
