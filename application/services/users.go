package services

import (
	"errors"
	"risqlac-api/application/models"
	"risqlac-api/infrastructure"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

var User userService

func (service *userService) GenerateSessionToken(email string, password string) (string, error) {
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

	token := uuid.NewString()

	err = Session.Create(&models.Session{
		Token:     token,
		UserId:    user.Id,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (*userService) ValidateSessionToken(token string) (models.User, error) {
	var session models.Session

	session, err := Session.GetByToken(token)

	if err != nil {
		return models.User{}, err
	}

	if time.Now().Unix() > session.ExpiresAt.Unix() {
		_ = Session.Delete(session.Id)
		return models.User{}, errors.New("token expired")
	}

	var user models.User

	user, err = User.GetById(session.UserId)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
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

	if user.IsAdmin > 0 {
		user.IsAdmin = 1
	}

	result := infrastructure.Database.Instance.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*userService) Update(user models.User) error {
	if user.IsAdmin > 0 {
		user.IsAdmin = 1
	}

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
