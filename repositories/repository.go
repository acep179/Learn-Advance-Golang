package repositories

//todo Iimport "gorm.io/gorm"
import "gorm.io/gorm"

//todo Declare repository struct here ...
type repository struct {
	db *gorm.DB
}

//ctt Tadinya, struct repository di atas kita masukkan ke dalam repositories/users.go
//ctt Sekarang, kita pindahkan struct tersebut ke sebuah file terpisah agar lebih rapi. Karena, struct tersebut akan digunakan oleh seluruh repositories, tidak hanya user saja.
