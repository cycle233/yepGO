package model

type UserRegister struct {
	Username string `form:"username" json:"username" gorm:"type:varchar(20);not null"`
	Email    string `form:"email" json:"email" binding:"email" gorm:"type:varchar(320);not null;unique"`
	Password string `form:"password" json:"password" gorm:"size:255;not null"`
}
