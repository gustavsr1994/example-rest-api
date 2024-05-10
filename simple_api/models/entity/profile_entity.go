package entity

type Profile struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:text" json:"username"`
	Email    string `gorm:"type:text" json:"email"`
	Password string `gorm:"type:varchar(8)" json:"password"`
	Photo    string `gorm:"type:text" json:"photo"`
}
