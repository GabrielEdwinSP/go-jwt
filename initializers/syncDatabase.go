package initializers

import (
	"github.com/GabrielEdwinSP/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
