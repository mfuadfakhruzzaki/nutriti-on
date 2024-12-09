// backend/controllers/user_service.go
package controllers

import (
	"errors"

	"github.com/mfuadfakhruzzaki/nutriti-on/backend/models"
	"github.com/mfuadfakhruzzaki/nutriti-on/backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
    db      *gorm.DB
    jwtUtil *utils.JWTUtil
}

func NewUserService(db *gorm.DB, jwtUtil *utils.JWTUtil) *UserService {
    return &UserService{
        db:      db,
        jwtUtil: jwtUtil,
    }
}

// Register mendaftarkan user baru
func (s *UserService) Register(name, email, password string) (*models.User, error) {
    // Cek apakah user sudah ada
    var existingUser models.User
    if err := s.db.Where("email = ?", email).First(&existingUser).Error; err == nil {
        return nil, errors.New("user already exists")
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Name:     name,
        Email:    email,
        Password: string(hashedPassword),
    }

    // Simpan user
    if err := s.db.Create(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}

// Login melakukan autentikasi user dan menghasilkan JWT token
func (s *UserService) Login(email, password string) (string, error) {
    var user models.User
    if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
        return "", errors.New("invalid email or password")
    }

    // Cek password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid email or password")
    }

    // Buat JWT token
    token, err := s.jwtUtil.GenerateToken(user.ID, user.Email)
    if err != nil {
        return "", err
    }

    return token, nil
}

// GetUser mengambil data user berdasarkan ID
func (s *UserService) GetUser(id uint) (*models.User, error) {
    var user models.User
    if err := s.db.First(&user, id).Error; err != nil {
        return nil, errors.New("user not found")
    }
    return &user, nil
}
