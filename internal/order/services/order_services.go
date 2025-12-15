package services

import (
	"errors"

	cartRepo "github.com/hemanth5544/goxpress/internal/cart/repository"
	"github.com/hemanth5544/goxpress/internal/order/model"
	"github.com/hemanth5544/goxpress/internal/order/repository"
	productRepo "github.com/hemanth5544/goxpress/internal/product/repository"
)

type OrderServices struct {
	orderRepo   *repository.OrderRepository
	productRepo *productRepo.ProductRepository
	cartRepo    *cartRepo.CartRepository
}

func NewOrderServices(orderRepo *repository.OrderRepository, productRepo *productRepo.ProductRepository, cartRepo *cartRepo.CartRepository) *OrderServices {
	return &OrderServices{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderServices) CheckoutService(userId uint, payment string) error {

	cart, err := s.cartRepo.GetCartByUserID(userId)
	if err != nil {
		return errors.New("cart not found")
	}

	if len(cart.Items) == 0 {
		return errors.New("Cart is empty")
	}

	var totalPrice float64
	var orderItems []model.OrderItem

	for _, item := range cart.Items {

		err := s.productRepo.UpdateProductStock(item.ProductID, item.Quantity)
		if err != nil {
			return errors.New("Failed to update product stock")
		}
		totalPrice += float64(item.Quantity) * item.Price
		orderItems = append(orderItems, model.OrderItem{
			ProductID:   item.ProductID,
			Quantity:    item.Quantity,
			PriceAtTime: item.Price,
		})

	}

	transaction := model.Transaction{
		UserID:     cart.UserID,
		TotalPrice: totalPrice,
	}
	err = s.orderRepo.CreateOrder(transaction, orderItems, payment)
	if err != nil {
		return errors.New("failed to add items to cart")
	}

	err = s.cartRepo.ClearCart(cart.ID)
	if err != nil {
		return errors.New("failed to empty cart")
	}

	return nil

}
