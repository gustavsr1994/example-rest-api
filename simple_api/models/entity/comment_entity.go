package entity

type Comment struct {
	Id      int64  `gorm:"primaryKey" json:"id"`
	Subject string `gorm:"type:varchar(100)" json:"subject"`
	Comment string `gorm:"type:text" json:"comment"`
}
