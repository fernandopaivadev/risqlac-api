package services

import (
	"os"
	"risqlac-api/database"
	"risqlac-api/models"

	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GenerateUserToken(email string, password string) (string, error) {
	var user models.User

	result := database.Instance.First(&user).Where(models.User{
		Email: email,
	})

	if result.Error != nil {
		return "", result.Error
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id":    user.Id,
		"expires_at": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func ValidateUserToken(tokenString string) (bool, jwt.MapClaims, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return false, nil, err
	}

	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		return true, claims, nil
	}

	return false, nil, nil
}

func CreateUser(user models.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(passwordHash)

	result := database.Instance.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdateUser(user models.User) error {
	result := database.Instance.Model(&user).Updates(models.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}).Where("id", user.Id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ListUsers(userId uint64) ([]models.User, error) {
	var users []models.User
	var result *gorm.DB

	if userId == 0 {
		result = database.Instance.Find(&users)
	} else {
		result = database.Instance.Find(&users, []uint64{userId})
	}

	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return users, nil
}

func DeleteUser(userId uint64) error {
	result := database.Instance.Delete(&models.User{}, userId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
