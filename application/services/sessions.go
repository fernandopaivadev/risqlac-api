package services

import (
	"risqlac-api/application/models"
	"risqlac-api/infrastructure"
)

type sessionService struct{}

var Session sessionService

func (*sessionService) Create(session *models.Session) error {
	result := infrastructure.Database.Instance.Create(&session)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*sessionService) GetByToken(token string) (models.Session, error) {
	var session models.Session

	result := infrastructure.Database.Instance.Where(&models.Session{
		Token: token,
	}).First(&session)

	if result.Error != nil {
		return models.Session{}, result.Error
	}

	return session, nil
}

func (*sessionService) GetByUserId(userId uint64) ([]models.Session, error) {
	var sessions []models.Session

	result := infrastructure.Database.Instance.Where(&models.Session{
		UserId: userId,
	}).Find(&sessions)

	if result.Error != nil {
		return []models.Session{}, result.Error
	}

	return sessions, nil
}

func (*sessionService) DeleteByToken(token string) error {
	result := infrastructure.Database.Instance.Where(&models.Session{
		Token: token,
	}).Delete(&models.Session{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (*sessionService) DeleteByUserId(userId uint64) error {
	result := infrastructure.Database.Instance.Where(&models.Session{
		UserId: userId,
	}).Delete(&models.Session{})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
