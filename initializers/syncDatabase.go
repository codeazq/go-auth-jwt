package initializers

import "github.com/codeazq/go-auth-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
