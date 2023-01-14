package model

// User model
type User struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Digest   string `gorm:"type:varchar(1024)" json:"-"`
	Nickname string `gorm:"type:varchar(1024)" json:"nickname"`
}
