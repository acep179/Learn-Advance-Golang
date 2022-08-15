//. Create package models, import "time",
package models

import "time"

//. Profile struct, ProfileResponse   struct, and ProfileResponse TableName method here ...
type Profile struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Phone     string               `json:"phone" gorm:"type: varchar(255)"`
	Gender    string               `json:"gender" gorm:"type: varchar(255)"`
	Address   string               `json:"address" gorm:"type: text"`
	UserID    int                  `json:"user_id"`
	User      UsersProfileResponse `json:"user"`
	CreatedAt time.Time            `json:"-"`
	UpdatedAt time.Time            `json:"-"`
}

//ctt Dapat kita lihat di atas bahwa struct terlihat masih sama saja
//ctt akan tetapi, pada struct tersebut terdapat User. Padahal, pada tabel profile kita tidak membutuhkan User, adapaun yang dibutuhkan hanyalah UserID-nya saja. Mengapa?
//ctt Karena UserID hanya digunakan untuk melakukan relasi terhadap tabel user. Sedangkan response-nya atau ketika kita Get Data dari User berdasarkan UserID tersebut, maka ia akan melalui UsersProfileResponse
//ctt Nah ketika ia melalui UserProfileResponse tersebut, maka ia akan mengarah ke struct UserProfileResponse yang barusan kita buat pada models.User

//ctt Artinya, ketika nanti kita hit GetProfile yang mana bentuknya dalam struct Profile, maka ia sudah berisikan UserProfileResponse
//ctt Sama halnya pada models.User, kita tidak ingin menampilkan seluruh data Profile di atas
//ctt Maka, kita buatkan struct ProfileResponse di bawah ini yang menunjukkan data apa saja yang akan ditampilkan ketika kita melakukan GetProfile

//. for association relation with another table (user)
type ProfileResponse struct {
	Phone   string `json:"phone"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
	UserID  int    `json:"-"`
}

//todo Sama seperti struct UserProfileResponse, maka kita nyatakan juga kepada gorm bahwa struct di atas tidak akan dimigrasi dengan menggunakan function di bawah ini

func (ProfileResponse) TableName() string {
	return "profiles"
}

//todo Kemudian, kita atur pula relasi pada Product
//ctt Yaitu, antara product dengan category,
