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

func (*sessionService) Delete(sessionId uint64) error {
	result := infrastructure.Database.Instance.Delete(&models.Session{}, sessionId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
