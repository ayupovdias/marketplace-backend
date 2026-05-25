package models

type Favorite struct {
	ID              uint `json:"id" gorm:"primaryKey"`
	UserID          uint `json:"user_id"`
	AdvertisementID uint `json:"advertisement_id"`
}
