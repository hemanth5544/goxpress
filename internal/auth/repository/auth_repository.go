package repository

import (
	"github.com/hemanth5544/goxpress/internal/auth/model"
	"github.com/hemanth5544/goxpress/internal/auth/dto"

	"gorm.io/gorm"
)

// IAuthRepository is an interface that defines the methods for user authentication and management.
// in service we will use this interface to call the methods
type IAuthRepository interface {
	CreateUser(userRequest model.User) error
	CheckUserExist(userLogin dto.LoginRequest) (*model.User, error)
	GetUserById(id int) (*model.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

//remember theat in GO we use all varaibiles rather using the module directly in this gorm.DB is passed to db
//this helps that Go had clean arch wiht no golabe deoency or gloabal vabrle just use when need dependcy injection when neede not always

// in GO we dont have classes so taht constictor s also is a fun the
// here NewAuthRepository is acting like constrict and  injecting db as a property to our AuthRepository
// so that we can use db every where use AuthRepository
func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

//---here r *AuthRepository is recvier  imaginere recvier
//like  a mehtod attched to this funtion createuesr so that you can acces AuthRepository properties in ypur
//createuesr  fun
//--------

func (r *AuthRepository) CreateUser(userRequest model.User) error {

	if err := r.db.Create(&userRequest).Error; err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) CheckUserExist(loginRequest dto.LoginRequest) (*model.User, error) {

	var user model.User

	if err := r.db.Where("email = ?", loginRequest.Email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *AuthRepository) GetUserById(id int) (*model.User, error) {

	var user model.User
	// we need to pass teh addres od user &user 
	//gorm internallly fill the fetched db data in hte user struct and gives us the 
	// finally the result will has two retun either .Error or result wiht the user struct 
	result := r.db.Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	userResponse := model.User{
		Username: user.Username,
		Email:    user.Email,
	}
	//ret
	return &userResponse, nil

}
