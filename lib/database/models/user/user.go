package user

import (
	"tobi-api/lib/database"
	"tobi-api/lib/database/models"
)

// User ...
type User struct {
	Account  string `json:"account" gorm:"type:varchar(20)" binding:"required"`  // 帳號
	Password string `json:"password" gorm:"type:varchar(20)" binding:"required"` // 密碼
	models.CommonModel
}

// Migration ...
func Migration() {
	db := database.GetDB()
	// HasTable 無法判斷有 schema 的 table ，
	// 故都會回傳 false
	// ref. https://github.com/jinzhu/gorm/issues/1197
	has := db.HasTable(&User{})
	if !has {
		db.AutoMigrate(&User{})
	}
}

// TableName ...
func (user *User) TableName() string {
	// return "ro.user"
	return "user"

}

// AfterFind : GORM Hooks
// process after user find
// ID 為奇數時， Name 改為 "Huang"
func (user *User) AfterFind() (err error) {
	// switch user.ID % 2 {
	// case 1:
	// 	user.Name = "Huang"
	// }
	return
}
