package profiledto

import "dumbmerch/models"

// Import "dumbmerch/models"

// Declare ProfileResponse  struct here ...
type ProfileResponse struct {
	ID      int                         `json:"id" gorm:"primary_key:auto_increment"`
	Phone   string                      `json:"phone" gorm:"type: varchar(255)"`
	Gender  string                      `json:"gender" gorm:"type: varchar(255)"`
	Address string                      `json:"address" gorm:"type: text"`
	UserID  int                         `json:"user_id"`
	User    models.UsersProfileResponse `json:"user"`
}

//ctt Dapat kita lihat pada User di atas, karena ia akan beralasi maka akan kita panggil dari models.UsersProfileResponse juga

//todo Persiapan terakhir adalah terkait Modify Repository pada file repositories/repository.go
