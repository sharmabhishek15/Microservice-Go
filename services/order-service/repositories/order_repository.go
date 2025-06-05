package repositories

import (
	"services/order-service/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) CreateOrder(order *models.Order) error {
	return repo.DB.Create(order).Error
}

func (repo *OrderRepository) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := repo.DB.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
