package service

import (
	"log"

	"github.com/go-resty/resty/v2"
)

type NotificationService struct {
	Client *resty.Client
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		Client: resty.New(),
	}
}

func (s *NotificationService) SendAdCreatedNotification(
	title string,
) {

	resp, err := s.Client.R().
		SetBody(map[string]interface{}{
			"message": "new advertisement created",
			"title":   title,
		}).
		Post("http://notification-service:8082/notify")

	if err != nil {
		log.Println("notification service error:", err)
		return
	}

	log.Println("notification sent:", resp.Status())
}
