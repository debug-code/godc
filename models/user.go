package models

type User struct {
	Id          int    `json:"id" gorm:"primary_key column:id"`
	Uid         string `json:"uid" gorm:"column:uid"`
	UName       string `json:"uName" gorm:"column:uName"`
	UPasswd     string `json:"uPasswd" gorm:"column:uPasswd"`
	UPasswdHash string `json:"uPasswdHash" gorm:"column:uPasswdHash"`
}
