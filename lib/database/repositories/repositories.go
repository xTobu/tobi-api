package repositories

import (
	modelUser "tobi-api/lib/database/models/user"

	"github.com/jinzhu/gorm"
)

// Setup 方便傳遞的 repositories Setup struct
type Setup struct {
	DB *gorm.DB
}

// Init : Initial repositories
func Init() {
	// MigrateTables()
}

// MigrateTables : 檢查並創建所有的表
func MigrateTables() {
	modelUser.Migration()
}
