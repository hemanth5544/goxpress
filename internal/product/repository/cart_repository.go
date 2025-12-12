package repository

import (
	"errors"

	"github.com/hemanth5544/goxpress/internal/cart/model"
	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

/*
* This NewCartRepository is construct func while give CartRepository
* as new intialised class(No to say in Go) coz Go doesnt have classes
* we can use NewCartRepository as new insance of cartRepo and use it any where as DI
 */

//? we will use this interface as to tight couple to acces repo function and all ok

type CartRepositoryInterface interface {
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) FindOrCreateCart(userID uint) (*model.Cart, error) {

	/*
		Input: userID is from user model
		Output: cart Model or error

		args: Find user cart or create new cart if it doesn't exist
	*/
	var cart model.Cart

	// Try to find an existing cart
	err := r.db.Where("user_id = ?", userID).First(&cart).Error
	if err == nil {
		// Cart already exists, return it
		return &cart, nil
	}
	//! Checking is error is other tahn not found like any other db error
	//? we need to retun err
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// If no existing cart, create a new one
	newCart := model.Cart{UserID: userID}
	if err := r.db.Create(&newCart).Error; err != nil {
		return nil, err
	}

	return &newCart, nil

}

func (r *CartRepository) AddToCart(cartID uint, item model.CartItem) error {

	item.CartID = cartID

	return r.db.Create(&item).Error

}

func (r *CartRepository) UpdateCartItem(item model.CartItem) error {

	return r.db.Save(&item).Error
}
