package models

// Article PO ,DTO
type Article struct {
	Id         int    `json:"id" gorm:"primary_key column:id"`
	Uid        string `json:"uid" gorm:"column:uid"`
	Content    string `json:"content" gorm:"column:content"`
	Title      string `json:"title" gorm:"column:title"`
	CreateTime int64  `json:"createTime" gorm:"column:createTime"`
	UpdateTime int64  `json:"updateTime" gorm:"column:updateTime"`
}
