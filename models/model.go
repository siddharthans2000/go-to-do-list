package model

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"username"`
}

func init() {

}
