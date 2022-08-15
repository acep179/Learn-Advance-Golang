package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"type: varchar(255)"`
	Email     string    `json:"email" gorm:"type: varchar(255)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

//ctt Jika kita perhatikan, maka terlihat bahwa pada Password, GeneratedAt, dan UpdatedAt json-nya itu dihilangkan alias hanya ditulis "-"
//ctt Karena untuk inputannya itu kita sudah memiliki DTO. Sehingga, untuk melakukan CreateUser, kita tidak menggunakan Models-nya, tapi kita menggunakan Password yang berasal dari DTO

//ctt Kita tidak perlu khawatir akan hal tersebut, karena ketika kita melakukan CreateUser, maka Password akan tetap terbaca pada DTO dan value-nya tetap akan dimasukkan ke dalam database
//ctt Adapun struct User di atas hanya berfungsi saat kita melakukan proses migrasi

//ctt Untuk struck di atas masih sama saja dan tidak ada perubahan sama sekali kecuali pada Password dll.
//ctt Kendati demikian, kita tetap harus menyiapkan bilamana tabel users dengan tabel profiles telah berelasi satu sama lain
//ctt Maka dari itu, kita buatkan struct UsersProfileResponse di bawah

type UsersProfileResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//ctt Ketika kita melihat kata Response di akhir nama struct di atas, mungkin kita akan berfikir mengapa struct di atas dituliskan pada Models, bukan pada DTO?
//ctt Nah, terkait hal ini di mana para tabel sudah berelasi maka response-response-nya ini kita letakkan langsung di dlm Models-nya agar ketika tabel Profile ingin memanggil referensinya melalui user_id maka ia dapat langsung membaca struct UsersProfileResponse di atas

//ctt Jadi, ketika nanti kita ambil datanya dan kita ingin Profile ini dimiliki oleh Users siapa saja maka nanti data User yang muncul ketika Response-nya itu hanya dua data saja yaitu ID dan Name-nya
//? Nanti akan lebih jelas ketika kita sudah melakukan penggabungannya

//ctt Karena kita meletakkan struct UserProfileResponse di dalam file user.go yang mana package-nya adalah models dan models itu yang akan kita migrasi. Sedangkan, untuk struct UserProfileResponse di atas itu bukan merupakan sebuah tabel yang akan dimigrasi, melainkan hanya Response-nya saja.
//ctt Maka dar itu, kita akan memberitahukan ke gorm bahwa struct tersebut bukan lah sebuah table yang akan dimigrasi
//ctt Tapi, ketika gorm membaca struct tersebut maka nanti akan langsung dikembalikan ke users yaitu struct User

//ctt Sehingga, pada dasarnya code di bawah function di bawah akan memberitahukan kepada gorm bahwa struct tersebut bukanlah hal yang harus dimigrasi

func (UsersProfileResponse) TableName() string {
	return "users"
}

//ctt Dua hal di atas merupakan proses menyiapkan respon relasi

//todo Karena pada models.User sudah, sekarang kita beralih ke models.Profile
