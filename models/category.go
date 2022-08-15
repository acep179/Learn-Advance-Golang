//. Create package models, import "time",
package models

import "time"

//. Profile struct here ...
type Category struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

//ctt Sekarang kita sudah mempunya struct mengenai Category pada models di mana struct ini akan dipanggil di bagian product yaitu pada bagian []Category

//todo Sekarang, kita siapkan terkait models.Transaction
