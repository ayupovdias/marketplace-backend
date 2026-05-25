package repository

import (
	"marketplace-service/internal/models"

	"gorm.io/gorm"
)

type AdvertisementRepository struct {
	DB *gorm.DB
}

func NewAdvertisementRepository(
	db *gorm.DB,
) *AdvertisementRepository {

	return &AdvertisementRepository{
		DB: db,
	}
}

func (r *AdvertisementRepository) Create(
	ad *models.Advertisement,
) error {

	return r.DB.Create(ad).Error
}

func (r *AdvertisementRepository) GetAll() (
	[]models.Advertisement,
	error,
) {

	var ads []models.Advertisement

	err := r.DB.Find(&ads).Error

	return ads, err
}

func (r *AdvertisementRepository) GetByID(
	id uint,
) (*models.Advertisement, error) {

	var ad models.Advertisement

	err := r.DB.First(&ad, id).Error

	return &ad, err
}

func (r *AdvertisementRepository) Delete(
	id uint,
) error {

	return r.DB.Delete(
		&models.Advertisement{},
		id,
	).Error
}
func (r *AdvertisementRepository) Update(
	ad *models.Advertisement,
) error {

	return r.DB.Save(ad).Error
}

func (r *AdvertisementRepository) GetByUserID(
	userID uint,
) ([]models.Advertisement, error) {

	var ads []models.Advertisement

	err := r.DB.
		Where("user_id = ?", userID).
		Find(&ads).Error

	return ads, err
}

func (r *AdvertisementRepository) GetByCategoryID(
	categoryID uint,
) ([]models.Advertisement, error) {

	var ads []models.Advertisement

	err := r.DB.
		Where("category_id = ?", categoryID).
		Find(&ads).Error

	return ads, err
}

func (r *AdvertisementRepository) Search(
	query string,
) ([]models.Advertisement, error) {

	var ads []models.Advertisement

	err := r.DB.
		Where(
			"title ILIKE ? OR description ILIKE ?",
			"%"+query+"%",
			"%"+query+"%",
		).
		Find(&ads).Error

	return ads, err
}
