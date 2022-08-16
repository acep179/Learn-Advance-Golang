//todo Write package repositories here ...
package repositories

//todo Import the required packages here ...
import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

//todo Declare ProductRepository interface here ...
type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProduct(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
}

//todo Write RepositoryProduct function here ...
func RepositoryProduct(db *gorm.DB) *repository {
	return &repository{db}
}

//ctt Untuk FindProducts dan GetProduct masih sama saja seperti sebelumnya. Hanya saja, kita harus menambahkan Preload("")
//ctt Adapun pada CreateProduct itu tidak ada perubahan sama sekali. Adapun yang akan di-handle pada CreateProduct itu akan kita handle di Handler-nya, bukan di Repository ini.

//todo Write FindProducts method here ...
func (r *repository) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("User").Find(&products).Error

	return products, err
}

//todo Write GetProduct method here ...
func (r *repository) GetProduct(ID int) (models.Product, error) {
	var product models.Product
	//todo not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("User").First(&product, ID).Error

	return product, err
}

//todo Write CreateProduct method here ...
func (r *repository) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	return product, err
}

//todo Sekarang, kita beralih ke Handler-nya
