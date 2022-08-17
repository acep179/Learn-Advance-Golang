package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	//todo Using Preload("User") to find data with relation to User and Preload("Category") for relation to Category here ...
	err := r.db.Preload("User").Preload("Category").Find(&products).Error

	return products, err
}

func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	//todo Using Preload("User") to find data with relation to User and Preload("Category") for relation to Category here ...
	err := r.db.Preload("User").Preload("Category").First(&product).Error

	//ctt Seperti biasa, cukup tambahkan Preload-nya saja, yang penting kita paham mengenai cara merelasikannya pada models-nya yaitu pada tahap preparation-nya

	//ctt Tetapi ada satu permasalahan yaitu ketika kita meng-create product dan mengisikan CategoryID-nya. Maka yang akan ter-update hanyalah CategoryID pada Products tersebut. Sedangkan, pada tabel jembatannya yaitu product_category hal tersebut tidaklah ter-handle
	//ctt Maka dari itu, kita perlu meng-handle-nya di handler CreateProduct

	return product, err
}

func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}
