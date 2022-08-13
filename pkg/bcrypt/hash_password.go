package bcrypt

import "golang.org/x/crypto/bcrypt"

//ctt Kita sudah tidak perlu lagi menginstal bcrypt karena ia sudah menjadi bawaan dari Golang-nya. Shingga, kita bisa langsung mengimport-nya saja.

//todo Dokumentasi bcrypt dapat dilihat pada GOBcrypt

//ctt Function HashingPassword digunakan untuk meng-enkripsi password ketika user melakukan registrasi

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	//ctt GenerateFromPassword([]byte(...) digunakan utk mengubah password yang diinputkan menjadi byte. Setelah menjadi byte barulah password-nya dienkripsi.
	// ctt Angka 10 artinya adalah salt-nya

	if err != nil {
		return "", err
	}
	return string(hashedByte), nil
	//ctt karena tadinya password yang diinputkan diubah terlebih dahulu ke dalam bentuk byte. Maka, string() di atas akan mengembalikan/meng-convert password dari byte menjadi string.
}

//ctt Function CheckPasswordHash digunakan ketika login utk mengecek apakah passsword yang dimasukkan oleh user tapat atau tidak

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

//todo Setelah ini, kita akan membuat sebuah handler dan repository baru yang berbeda dan terpisah dari handler dan repository-nya user walaupun secara teknis antara Register dan Create User itu sama-sama menambah data pada tabel user
