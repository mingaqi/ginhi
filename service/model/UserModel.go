package model

type UserModel struct {
	Email         string `form:"email" binding:"required" json:"email"`
	Password      string `form:"password" json:"password"`
	PasswordAgain string `form:"passwordAgain" json:"passwordAgain"`
}
