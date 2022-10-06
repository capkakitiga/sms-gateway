package repositories

import (
	"bitbucket.org/capcom6/smsgatewaybackend/internal/smsgateway/models"
	"gorm.io/gorm"
)

var (
	ErrDeviceNotFound = gorm.ErrRecordNotFound
)

type DevicesRepository struct {
	db *gorm.DB
}

func (r *DevicesRepository) GetByToken(token string) (models.Device, error) {
	device := models.Device{}

	return device, r.db.Where("auth_token = ?", token).Take(&device).Error
}

func (r *DevicesRepository) Insert(device *models.Device) error {
	return r.db.Create(device).Error
}

func NewDevicesRepository(db *gorm.DB) *DevicesRepository {
	return &DevicesRepository{
		db: db,
	}
}