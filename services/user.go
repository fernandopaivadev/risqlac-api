package services

import (
	"encoding/json"
	"errors"
	"risqlac-api/database"
	"risqlac-api/environment"
	"risqlac-api/types"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

var User UserService

type TokenClaims struct {
	UserId    uint64 `json:"UserId"`
	ExpiresAt int64  `json:"ExpiresAt"`
}

func (service *UserService) GenerateToken(email string, password string) (string, error) {
	var user types.User

	user, err := service.GetByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"UserId":    user.Id,
		"ExpiresAt": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(environment.Variables.JwtSecret))

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func (service *UserService) GeneratePasswordChangeToken(email string) (string, error) {
	var user types.User

	user, err := service.GetByEmail(email)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"UserId":    user.Id,
		"ExpiresAt": time.Now().Add(5 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(environment.Variables.JwtSecret))

	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

func (_ *UserService) ParseToken(tokenString string) (TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(environment.Variables.JwtSecret), nil
	})

	if err != nil {
		return TokenClaims{}, err
	}

	if !token.Valid {
		return TokenClaims{}, errors.New("invalid token")
	}

	jwtClaims, _ := token.Claims.(jwt.MapClaims)
	var claimsObject TokenClaims
	claimsJSON, _ := json.Marshal(jwtClaims)
	err = json.Unmarshal(claimsJSON, &claimsObject)

	if err != nil {
		return TokenClaims{}, err
	}

	return claimsObject, nil
}

func (_ *UserService) ChangePassword(userId uint64, newPassword string) error {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	result := database.Instance.Model(&types.User{
		Id: userId,
	}).Updates(types.User{
		Password: string(passwordHash),
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (_ *UserService) Create(user types.User) error {
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

func (_ *UserService) Update(user types.User) error {
	result := database.Instance.Model(&user).Select(
		"Email", "Name", "Phone", "Is_admin",
	).Updates(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (_ *UserService) GetById(userId uint64) (types.User, error) {
	var user types.User

	result := database.Instance.First(&user, userId)

	if result.Error != nil {
		return types.User{}, result.Error
	}

	return user, nil
}

func (_ *UserService) GetByEmail(email string) (types.User, error) {
	var user types.User

	result := database.Instance.Where(&types.User{
		Email: email,
	}).First(&user)

	if result.Error != nil {
		return types.User{}, result.Error
	}

	return user, nil
}

func (_ *UserService) List() ([]types.User, error) {
	var users []types.User

	result := database.Instance.Find(&users)

	if result.Error != nil {
		return []types.User{}, result.Error
	}

	return users, nil
}

func (_ *UserService) Delete(userId uint64) error {
	result := database.Instance.Delete(&types.User{}, userId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
