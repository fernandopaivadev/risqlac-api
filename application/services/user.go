package services

import (
	"risqlac-api/application/models"
	"risqlac-api/infrastructure"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

var User userService

type tokenClaims struct {
	UserId    uint64 `json:"UserId"`
	ExpiresAt int64  `json:"ExpiresAt"`
}

func (service *userService) GenerateLoginToken(email string, password string) (string, error) {
	var user models.User

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

	tokenString, err := Utils.GenerateJWT(
		user.Id,
		time.Now().Add(24*time.Hour).Unix(),
	)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (service *userService) GeneratePasswordChangeToken(email string) (string, error) {
	var user models.User

	user, err := service.GetByEmail(email)

	if err != nil {
		return "", err
	}

	tokenString, err := Utils.GenerateJWT(
		user.Id,
		time.Now().Add(5*time.Minute).Unix(),
	)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (*userService) ChangePassword(userId uint64, newPassword string) error {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	result := infrastructure.Database.Instance.Model(&models.User{
		Id: userId,
	}).Updates(models.User{
		Password: string(passwordHash),
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*userService) Create(user models.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(passwordHash)

	result := infrastructure.Database.Instance.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*userService) Update(user models.User) error {
	result := infrastructure.Database.Instance.Model(&user).Select(
		"Email", "Name", "Phone", "Is_admin",
	).Updates(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*userService) GetById(userId uint64) (models.User, error) {
	var user models.User

	result := infrastructure.Database.Instance.First(&user, userId)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (*userService) GetByEmail(email string) (models.User, error) {
	var user models.User

	result := infrastructure.Database.Instance.Where(&models.User{
		Email: email,
	}).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (*userService) List() ([]models.User, error) {
	var users []models.User

	result := infrastructure.Database.Instance.Find(&users)

	if result.Error != nil {
		return []models.User{}, result.Error
	}

	return users, nil
}

func (*userService) Delete(userId uint64) error {
	result := infrastructure.Database.Instance.Delete(&models.User{}, userId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}