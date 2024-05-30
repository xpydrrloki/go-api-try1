package model

import (

	"gorm.io/datatypes"
)

type Users struct {
	Id        int            `gorm:"type:int;primary_key" json:"id"`
	Username  string         `gorm:"type:varchar(255)" json:"username"`
	Email     string         `gorm:"type:varchar(255)" json:"email"`
	Password  string         `gorm:"type:varchar(255)" json:"password"`
	CreatedAt datatypes.Date `gorm:"type:DATE;DEFAULT:CURRENT_TIMESTAMP" json:"createdAt"`
}
