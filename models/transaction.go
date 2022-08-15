// Create package models, import "time",
package models

import "time"

// Transaction  struct here ...
type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	ProductID int                  `json:"product_id"`
	Product   ProductResponse      `json:"product"`
	BuyerID   int                  `json:"buyer_id"`
	Buyer     UsersProfileResponse `json:"buyer"`
	SellerID  int                  `json:"seller_id"`
	Seller    UsersProfileResponse `json:"seller"`
	Price     int                  `json:"price"`
	Status    string               `json:"status"  gorm:"type:varchar(25)"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

//ctt Terkait transaction ini tidak akan terlalu banyak dibahas karena relasinya sama seperti user dan product yaitu one to many alias belongsTo

//todo Setelah ini, kita siapkan DTO-nya
