package service

import (
	"marketplace-service/internal/models"
	"marketplace-service/internal/repository"
)

type AdvertisementService struct {
	Repo                *repository.AdvertisementRepository
	NotificationService *NotificationService
}

func NewAdvertisementService(
	repo *repository.AdvertisementRepository,
	notificationService *NotificationService,
) *AdvertisementService {

	return &AdvertisementService{
		Repo:                repo,
		NotificationService: notificationService,
	}
}

func (s *AdvertisementService) CreateAd(
	ad *models.Advertisement,
) error {

	err := s.Repo.Create(ad)

	if err != nil {
		return err
	}

	s.NotificationService.SendAdCreatedNotification(
		ad.Title,
	)

	return nil
}
