//todo Create package repositories here ...
package repositories

//todo Import the required packages here ...
import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

//ctt Isi dari Repository ini sama saja seperti yang ada pada repositories/users.go di mana ada interface dan beberapa function. Sehingga, sebenarnya kita cukup copas saja.

//todo Declare ProfileRepository interface here ...
type ProfileRepository interface {
	GetProfile(ID int) (models.Profile, error)
}

//todo Create RepositoryProfile function here ...
func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

//todo Create GetProfile method here ...
func (r *repository) GetProfile(ID int) (models.Profile, error) {
	//ctt Sebagai mana pada GetUser di mana ia membutuhkan ID User, pada GetProfile ini pun kita membutuhkan ID dari Profile yang ingin kita cari dan tampilkan. Maka dari itu, kita isikan (ID int) pada parameter Method GetProfile ini.

	var profile models.Profile
	err := r.db.Preload("User").First(&profile, ID).Error

	//ctt Karena sekarang konsepnya adalah multi-table di mana ketika kita meng-Get data Profile-nya, kemudian kita hubungkan where UserID-nya
	//ctt Adapun caranya adalah dengan menambahkan Method Preload("User") sebagaimana di atas.
	//ctt Preload ke mana "User" itu? User tsb kita ambil pada models.Profile. Yaitu, pada property dari struct tersebut di mana User tersebut yang akan menghubungkan ke UsersProfileResponse
	//ctt Jika ternyata property struct yang menghubungkan ke (atau yang tipe datanya adalah) UsersProfileResponse adalah UserData, bukan User saja. Maka, pada Preload("User") di atas kita harus mengisinya dengan "UserData" menjadi Preload("UserData"). Intinya, harus disamakan dengan property atau kata kuncinya.

	//ctt Kemudian, First di atas artinya sama saja yaitu mencari satu data pertama yang memiliki ID sesuai dengan yang kita kirimkan ke parameternya

	//ctt Jika pada Sequelize, Preload("") itu bagaikan include
	//ctt Kita juga dapat menerapkan Preload ini ke dalam Method Find

	return profile, err
}

//todo Setelah kita selesai mengatur repository-nya, kini saatnya kita menyiapkan handler-nya.
