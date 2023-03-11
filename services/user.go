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

func GenerateUserToken(email string, password string) (string, error) {
	var user types.User

	user, err := GetUserByEmail(email)

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

func GeneratePasswordChangeToken(email string) (string, error) {
	var user types.User

	user, err := GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id":    user.Id,
		"expires_at": time.Now().Add(5 * time.Minute).Unix(),
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
		return types.TokenClaims{}, errors.New("invalid token")
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

func ChangeUserPassword(userId uint64, newPassword string) error {
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

func CreateUser(user types.User) error {
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

func UpdateUser(user types.User) error {
	result := database.Instance.Model(&user).Select(
		"Email", "Name", "Phone", "Is_admin",
	).Updates(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserById(userId uint64) (types.User, error) {
	var user types.User

	result := database.Instance.First(&user, userId)

	if result.Error != nil {
		return types.User{}, result.Error
	}

	return user, nil
}

func GetUserByEmail(email string) (types.User, error) {
	var user types.User

	result := database.Instance.Where(&types.User{
		Email: email,
	}).First(&user)

	if result.Error != nil {
		return types.User{}, result.Error
	}

	return user, nil
}

func ListUsers() ([]types.User, error) {
	var users []types.User

	result := database.Instance.Find(&users)

	if result.Error != nil {
		return []types.User{}, result.Error
	}

	return users, nil
}

func DeleteUser(userId uint64) error {
	result := database.Instance.Delete(&types.User{}, userId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
