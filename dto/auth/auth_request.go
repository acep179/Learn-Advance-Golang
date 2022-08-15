package authdto

//todo Declare AuthRequest struct here ...
type AuthRequest struct {
	Name     string `gorm:"type: varchar(255)" json:"name"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
}

//ctt Walaupun auth ini belum akan kita handle, tapi untuk sekarang kita siapkan saja untuk proses Registernya

//todo Setelah ini kita siapkan DTO Product Request nya
