package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user models.User) (models.User, error)
	//ctt Untuk parameter yang diterima masih mengacu kepada models.User nilain yang dikembalikannya pun sama.
}

func RepositoryAuth(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

//ctt Method Register di atas sebenarnya sama saja seperti Method Create User, bahkan sama persis. Adapun yg membedakannya hanyalah pada penamaan Method-nya saja.

//todo Sebagaimana dalam pembuatan API lainnya, maka setelah ini melakukan tahap ini barulah kita buatkan handler-nya.
