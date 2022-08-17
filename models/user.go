package models

import "time"

type User struct {
	ID       int                   `json:"id"`
	Name     string                `json:"name" gorm:"type: varchar(255)"`
	Email    string                `json:"email" gorm:"type: varchar(255)"`
	Password string                `json:"-" gorm:"type: varchar(255)"`
	Profile  ProfileResponse       `json:"profile"`
	Products []ProductUserResponse `json:"products"`
	//ctt Perlu diperhatikan bahwa sekarang ada propery Products yang akan berelasi ke ProductUserResponse
	//ctt Terlihat juga ada sedikit perbedaan antara Pofile dan Products di atas. Yaitu, pada Products untuk perelasiannya menggunakan [] alias slice karena case-nya adalah one to many yang artinya satu User bisa mempunyai banyak Products. Itulah sebabnya kita menggunakan slice agar data Products dapat ditampung, bukan hanya satu akan tetapi banyak data Products.
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

//todo Sekarang, kita lakukan perubahan pada repository users

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
