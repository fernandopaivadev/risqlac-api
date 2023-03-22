package services

import (
	"risqlac-api/infra"
	"risqlac-api/types"
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
	var user types.User

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

	result := infra.Database.Instance.Model(&types.User{
		Id: userId,
	}).Updates(types.User{
		Password: string(passwordHash),
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*userService) Create(user types.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(passwordHash)

	result := infra.Database.Instance.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*userService) Update(user types.User) error {
	result := infra.Database.Instance.Model(&user).Select(
		"Email", "Name", "Phone", "Is_admin",
	).Updates(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*userService) GetById(userId uint64) (types.User, error) {
	var user types.User

	result := infra.Database.Instance.First(&user, userId)

	if result.Error != nil {
		return types.User{}, result.Error
	}

	return user, nil
}

func (*userService) GetByEmail(email string) (types.User, error) {
	var user types.User

	result := infra.Database.Instance.Where(&types.User{
		Email: email,
	}).First(&user)

	if result.Error != nil {
		return types.User{}, result.Error
	}

	return user, nil
}

func (*userService) List() ([]types.User, error) {
	var users []types.User

	result := infra.Database.Instance.Find(&users)

	if result.Error != nil {
		return []types.User{}, result.Error
	}

	return users, nil
}

func (*userService) Delete(userId uint64) error {
	result := infra.Database.Instance.Delete(&types.User{}, userId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
