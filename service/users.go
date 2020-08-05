package service

import (
	"fmt"
	"godc/base"
	"godc/models"
)

func Users() ([]models.User, error) {
	users, err := base.Users()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return users, nil
}
