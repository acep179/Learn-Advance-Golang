package middleware

import (
	"context"
	dto "dumbmerch/dto/result"
	jwtToken "dumbmerch/pkg/jwt"
	"encoding/json"
	"net/http"
	"strings"
)

//ctt Konsepnya, proses otentikasi itu bukanlah proses yang kita letakkan seperti layering. Proses otentikasi adalah sebuah proses yang terjadi di antara mereka. Sehingga, posisi proses otentikasi adalah sebagai Middleware atau sebagai jembatan.
//ctt Jadi, ketika user ingin mengakses main-nya kemudian masuk ke routes-nya. Nah, ketika ia ingin mengakses routes ini, ia harus melalui midlleware-nya terlebih dahulu. Yaitu, middleware terkait otentikasi. Sehingga, proses otentikasi terjadi bukan setelah user mengakses routes-nya, akan tetapi hal tesebut terjadi sebelumnya, sebelum user mengakses routes atau endpoint-nya.
//ctt Maka dari itu, di sini kita buatkan sebuah Middleware yang dapat kita panggil di bagian routes-nya agar nantinya kita bisa mengatur agar user melewati middleware-nya terlebih dahulu sebelum mengakses endpoint.

/*
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
*/

//todo untuk Result di atas juga bisa kita gantikan dengan errorResult yang terdapat pada dto sehingga struct Result di atas bisa kita hapus (untuk saat ini saya komentari saja)
//todo Kemudian, untuk kodingan di bawah yang masih menggunakan Result di atas langsung kita ganti saja dengan dto yg telah tersedia

//ctt Jadi, sebagaimana penjelasan sebelumnya bahwasanya konsep dr Middleware itu sebagai perantara. Maka dar itu, pada semua Middleware ada yang namanya sebuah Method http.HandlerFunc yg nantinya akan meneruskan ke proses selanjutnya atau proses intinya.

func Auth(next http.HandlerFunc) http.HandlerFunc {

	//ctt Jika kita ingat-ingat maka kita sadar bahwa HandlerFunc itu kita gunakan pada routes

	//ctt Sekarang, bagaimana caranya agar nanti si Middleware ini meneruskan ke routes-nya.
	//ctt Nah, di bawah ini barulah kita return-kan si HandlerFunc juga yang niali kembalian atau nilai yg dibawanya adalah ResponWriter dan Request selayaknya kita menggunakannya di Handler.

	//ctt Mengapa kita harus menambahakn keduanya? Apakah ini juga akan mengirimkan response? Jawabannya adalah "iya" bahwa si Middleware juga dapat mengirimkan response.
	//ctt Response-nya apa? Response-nya adalah jika berhasil maka akan diteruskan ke proses selanjutnya, jika tidak berhasil atau error maka dia akan berhenti.

	//ctt Sehingga, tidak hanya Handler yg dapat mengirim response dan menerima request, tetapi Middleware pun dapat melakukannya.
	//ctt Bahkan, jika kita tidak meng-handle request-nya, maka bagaimana bisa request yang dikirim dari main sampai ke handler-nya?

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		//ctt di bawah ada variabel token yang kita dapatkan dari request.Header
		//ctt Sedikit intermezo bahwasanya setiap kita mengirimkan request, maka akan ada yang namanya Header dan pada Header tersebut terdapat salah satu kata kunci dengan nama Authorization
		//ctt Nah, token ini kita ambil dari Authorization tersebut

		//! Perlu dipahami kembali bahwa tujuan dari function Auth ini adalah untuk mengecek apakah saat seorang user mengakses suatu endpoint yang telah kita tentukan, apakah si User tersebut memiliki token atau tidak? Maka dari itu, untuk melakukan pengecekannya kita cek atau kita periksa pada bagian Authorization yang terdapat pada Header

		token := r.Header.Get("Authorization")

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		//ctt Jika ternyata pada Authorization tersebut ada isinya, maka kita harus melakukan pengecekan apakah token tersebut benar atau tidak, valit atau tidak.
		//ctt Maka dari itu, di bawah kita panggil proses DecodeToken().

		//ctt Tapi, tidak semua isi dari Authorization tersebut kita Decode-kan. Karena, jika kita perhatikan maka value dari si Authorization tidak hanya terdapat token saja tetapi sebelum token tersebut ada kata "Bearer" dan diikuti oleh satu spasi.
		//ctt Mengapa ada kata "Bearer"? Karena pada Authorization-nya kita menggunakan tipe kemanan berupa Bearer Token.

		//ctt Maka dari itu, sebelum kita Decode-kan token-nya, maka kita Split() terlebih dahulu untuk menghilangkan kata "Bearer" beserta satu spasi setelahnya dengan cara sebagaimana di bawah ini.
		//ctt Kita split value dari Authorization (token) berdasarkan spasi (" "). Lalu Split tersebut akan membuatkan data berupa Array di mana pada index [0] adalah kata "Bearer" dan pada index [1] adalah tokennya. Maka dari itu kita isikan [1] setelah Split()

		token = strings.Split(token, " ")[1]
		claims, err := jwtToken.DecodeToken(token)

		//ctt Terdapat pertanyaan lagi, mengapa langsung kita lakukan DecodeToken? Kapan proses pengecekan apakah token tersebut benar atau tidaknya?
		//ctt Tentu saja, prose tersebut dilakukan sekalian ketika kita melakukan DecodeToken. Karena, pada function jwtToken.DecodeToken() itu akan memanggil function VerifyToken yang berada di atasnya. Untuk penjelasan lengkapnya silakan lihat ke jwtToken.DecodeToken().

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "unauthorized"}
			//todo Sebelum-nya adalah => Result{Code: http.StatusUnauthorized, Message: "unauthorized"}
			json.NewEncoder(w).Encode(response)
			return
		}

		//ctt Berdasarkan context-nya,
		//ctt Jadi, di response itu ada sifatnya context. Context itu berari yang mana membungkus dari si bagian dari Middleware-nya ini.
		//ctt Jadi, dia akan meneruskan untuk ke seluruh proses selanjutnya itu bisa menggunakan yang namanya "userInfo" sebagaimana di bawah
		//ctt Adapun "userInfo" itu akan berisikan payload yang sudah kita tambahkan ke tokennya tadi

		ctx := context.WithValue(r.Context(), "userInfo", claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
