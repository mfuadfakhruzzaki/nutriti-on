// backend/config/config.go
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config struct untuk menyimpan konfigurasi aplikasi
type Config struct {
    Port        string
    JWTSecret   string
    DatabaseURL string
}

// LoadConfig memuat konfigurasi dari environment variables
func LoadConfig() (*Config, error) {
    // Load file .env
    err := godotenv.Load()
    if err != nil {
        log.Printf("No .env file found")
        // Anda bisa mengembalikan error jika file .env wajib ada
        // return nil, fmt.Errorf("error loading .env file: %w", err)
    }

    cfg := &Config{
        Port:        getEnv("PORT", "8080"),
        JWTSecret:   getEnv("JWT_SECRET", "your_jwt_secret"),
        DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@host:port/dbname?sslmode=disable"),
    }

    return cfg, nil
}

// getEnv mengambil environment variable atau mengembalikan nilai default
func getEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}

// InitDB menginisialisasi koneksi ke PostgreSQL
func InitDB(databaseURL string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    // Migrasi model
    err = db.AutoMigrate(&models.User{})
    if err != nil {
        return nil, fmt.Errorf("failed to migrate database: %w", err)
    }

    return db, nil
}

// CloseDB menutup koneksi database
func CloseDB(db *gorm.DB) {
    sqlDB, err := db.DB()
    if err != nil {
        log.Printf("Error getting DB from GORM: %v", err)
        return
    }
    err = sqlDB.Close()
    if err != nil {
        log.Printf("Error closing database connection: %v", err)
    }
}
