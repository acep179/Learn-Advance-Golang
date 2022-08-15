//. Create package models, import "time",
package models

import "time"

//. Product struct, ProductResponse struct, ProductUserResponse  struct, ProductResponse TableName method and ProductUserResponse TableName method here ...

type Product struct {
	ID       int                  `json:"id" gorm:"primary_key:auto_increment"`
	Name     string               `json:"name" form:"name" gorm:"type: varchar(255)"`
	Desc     string               `json:"desc" gorm:"type:text" form:"desc"`
	Price    int                  `json:"price" form:"price" gorm:"type: int"`
	Image    string               `json:"image" form:"image" gorm:"type: varchar(255)"`
	Qty      int                  `json:"qty" form:"qty"`
	UserID   int                  `json:"user_id" form:"user_id"`
	User     UsersProfileResponse `json:"user"`
	Category []Category           `json:"category" gorm:"many2many:product_categories"`

	//ctt Karena relasi antara Product dan Category adalah many to many. Maka, pada gorm: di atas kita tuliskan gorm: "many2many:product_categories" adapun product_category adalah ForeignKey
	//ctt Untuk sekarang tetntu saja []Category di atas masih error karena untuk Category-nya memang belum disiapkan

	//ctt Pada CategoryID di bawah dapat dilihat bahwa pada gorm: itu diisikan "-" menjadi gorm:"-" tujuannya adalah agar ketika proses migration dijalankan maka kolom category_id tidak dibuatkan di tabel product mengingat relasinya Category dengan Product adalah many to many sehingga dia akan membuat tabel jembatannya

	CategoryID []int     `json:"category_id" form:"category_id" gorm:"-"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}

//ctt Di bawah ini adalah ProductResponse-nya sehingga ketika kita ingin GetProduct-nya saja maka response-nya akan seperti di bawah ini

type ProductResponse struct {
	ID         int                  `json:"id"`
	Name       string               `json:"name"`
	Desc       string               `json:"desc"`
	Price      int                  `json:"price"`
	Image      string               `json:"image"`
	Qty        int                  `json:"qty"`
	UserID     int                  `json:"-"`
	User       UsersProfileResponse `json:"user"`
	Category   []Category           `json:"category" gorm:"many2many:product_categories"`
	CategoryID []int                `json:"category_id" form:"category_id" gorm:"-"`
}

//ctt Tapi, kalo ProductUserRespone berarti kita nge-Get dari Products ke Users alias BelongsTo
//ctt Ketika kita menggunakan BelongsTo nanti, maka kita akan menggunakan struct di bawah ini di mana kita memanggil si UserID-nya
//ctt Tujuannya adalha agar tahu setiap product itu dimiliki oleh user siapa

type ProductUserResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	Qty    int    `json:"qty"`
	UserID int    `json:"-"`
}

//ctt Karena kedua struct di atas bukanlah tabel yang akan direlasi, maka kita beritahukan kepada gorm menggunakan function di bawah ini

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductUserResponse) TableName() string {
	return "products"
}

//todo Sekarang, barulah kita siapkan untuk bagian Category-nya
