package services

import (
	"encoding/json"
	"risqlac-api/database"
	"risqlac-api/environment"
	"risqlac-api/models"
	"risqlac-api/types"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
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

	tokenString, err := token.SignedString([]byte(environment.Get().JWT_SECRET))

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func ParseUserToken(tokenString string) (types.TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(environment.Get().JWT_SECRET), nil
	})

	if err != nil {
		return types.TokenClaims{}, err
	}

	if !token.Valid {
		return types.TokenClaims{}, types.MakeCustomError("Invalid token")
	}

	jwtClaims, _ := token.Claims.(jwt.MapClaims)
	var claimsObject types.TokenClaims
	claimsJSON, _ := json.Marshal(jwtClaims)
	err = json.Unmarshal(claimsJSON, &claimsObject)

	if err != nil {
		return types.TokenClaims{}, err
	}

	return claimsObject, nil
}

func ResetUserPassword(userId uint64, newPassword string) error {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	result := database.Instance.Model(&models.User{}).Updates(models.User{
		Password: string(passwordHash),
	}).Where("id", userId)

	if result.Error != nil {
		return result.Error
	}

	return nil
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

func GetUser(userId uint64) (models.User, error) {
	var user models.User

	result := database.Instance.First(&user, userId)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func ListUsers() ([]models.User, error) {
	var users []models.User

	result := database.Instance.Find(&users)

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
