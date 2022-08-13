package jwtToken

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = "SECRET_KEY"

//! Untuk sementara ini SecretKey-nya kita buatakan sebagai variabel biasa. Tapi nantinya akan kita taruh secretKey-nya ini ke dlm .env

func GenerateToken(claims *jwt.MapClaims) (string, error) {

	//ctt Maksud dari MapClaims adalah apabila kita tidak menyiapkan atau meng-handle terkait bagaimana proses generate token-nya, maka kita gunakan MapClaims tersebut yang nantinya akan mengembalikan dua data
	//ctt Dua data tersebut adalah string yang isinya adalah token yang sudah ter-generate dan data yang ke dua adalah error-nya

	//ctt Adapun cara untuk men-generate tokennya adalah dengan menggunakan NewWithClaims di bawah yang merupakan bawaan dari si jwt-nya langsung tanpa perlu kita buatkan sendiri.

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//ctt Terkait jwt.SigningMethodHS256, jika kita ingin men-spesifik-kan Method atau Algoritma dari si JWT-nya ini ingin menggunakan hash apa, maka akan ada banyak pilihan Method atau Algoritma yang tersedia untuk melakukan hashing.
	//ctt Jika kita ingin menentukan secara spesifik, maka kita gunakan NewWithClaims() di atas.
	//ctt Adapun saat ini kita gunakan Algoritma SigningMethodHS256 sebagaimana di atas.

	//ctt claims adalah data yang ingin kita generate-kan jwt-nya. Claims didapat dari function parameter.

	//todo Setelah kita buatkan token-nya, maka kita tambahkan atau kita kombinasikan dengan SecretKey-nya sebagaimana di bawah.

	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webtoken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}
	//ctt Artinya, tokennya ada tapi tidak valid

	return token, nil
	//ctt Jika tokennya valid maka tokennya akan kita kirimkan
}

//ctt DecodeToken juga digunakan untuk menggunakan data yang sudah di-hashing
//ctt Karena proses DecodeToken adalah sebuah proses penerjemahan dari token menjadi data yang kita simpan di dalamnya.

func DecodeToken(tokenString string) (jwt.MapClaims, error) {

	token, err := VerifyToken(tokenString)

	//ctt Kodingan di bawah adalah proses Decode-nya. Akan tetapi, sebelum proses Decode dilakukan, function ini akan mengecek apakah token yang diberikan itu sudah benar atau belum dengan cara memanggil function VerifyToken() pada kodingan di atas catatan ini.
	//ctt Adapun function VerifyToken() itu terdapat di atas function DecodeToken() ini.

	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)
	//ctt Mengecek apakah datanya oke atau tidak

	if isOk && token.Valid {
		return claims, nil
	}
	//ctt Melakukan pengecekan. Jika datanya oke dan tokennya valid maka akan me-return claims-nya dan nil
	//ctt Seharusnya yg dikembalikan itu claims dan err-nya. Akan tetapi, karena kita tidak meng-handle err-nya. Maka, kita kembalikan nil saja yg mana nil itu pun berasal dari err-nya.

	return nil, fmt.Errorf("invalid token")
}
