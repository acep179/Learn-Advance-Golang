package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User

	//todo Using Preload("profile") to find data with relation to profile and Preload("Products") for relation to Products here ...
	err := r.db.Preload("Profile").Preload("Products").Find(&users).Error // add this code

	//ctt Kita cukup menambahkan satu Preload("") lagi untuk menampilkan data Products-nya
	//ctt Kita tidak perlu lagi menambahkan satu Find lagi di bawahnya, cukup tambahkan atau sisipkan satu Preload lagi untuk Products
	//ctt Karena, jika kita Debug() maka sebenarnya GORM nya itu langsung menjalankan tiga perintah query untuk masing-masing Method

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User

	//todo Using Preload("profile") to find data with relation to profile and Preload("Products") for relation to Products here ...
	err := r.db.Preload("Profile").Preload("Products").First(&user, ID).Error // add this code

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
