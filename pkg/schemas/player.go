package schemas

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"column:username"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
}
