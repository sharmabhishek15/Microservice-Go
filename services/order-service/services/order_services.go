package services

import (
	"services/order-service/models"
	"services/order-service/repositories"
)

type OrderService struct {
	Repo *repositories.OrderRepository
}

func NewOrderService(repo *repositories.OrderRepository) *OrderService {
	return &OrderService{Repo: repo}
}

func (s *OrderService) CreateOrder(order *models.Order) error {
	order.Status = "Pending"
	return s.Repo.CreateOrder(order)
}

func (s *OrderService) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	return s.Repo.GetOrdersByUserID(userID)
}
