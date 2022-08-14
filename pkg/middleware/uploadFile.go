package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//todo Pertama, kita buatkan sebuah function UploadFile yang akan me-return routes-nya

func UploadFile(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Upload file
		// FormFile returns the first file for the given key `myFile`
		// it also returns the FileHeader so we can get the Filename,
		// the Header and the size of the file
		file, _, err := r.FormFile("image")

		//ctt kata "image" harus disesuaikan dengan name pada form input HTML
		//ctt pada r.FormFile di atas, ia mengembalikan 3 nilai yaitu file-nya, handler, dan error-nya
		//ctt Jika user tidak memasukkan gambar, maka akan terjadi error dan error tersebut akan di-handle di bawah ini

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error Retrieving the File")
			return
		}
		defer file.Close()

		//ctt Adapun terkait handler-nya, terdapat tiga nilai di dalam handler tsb, tiga nilai tsb adalah Filename, Size, dan Header
		//ctt Ketiga hal tsb tidak kita gunakan, karena biasanya tiga hal tersebut hanya digunakan untuk pengecekan saja seperti beberapa contoh di bawah

		//ctt fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		//ctt fmt.Printf("File Size: %+v\n", handler.Size)
		//ctt fmt.Printf("MIME Header: %+v\n", handler.Header)

		//ctt Kita juga dapat mengatur Max Size dari gambar yang di-upload
		const MAX_UPLOAD_SIZE = 10 << 20 // 10MB
		//ctt This is for Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.

		//ctt Kemudian, Max Size-nya kita masukkan ke MultipartForm
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)

		//ctt Setelah itu, kita buat sebuah pengkondisian di mana r.ContentLength adalah size dari gambar yang di-upload dan MAX_UPLOAD_SIZE adalah batas maksimal size gambar yang di-upload
		if r.ContentLength > MAX_UPLOAD_SIZE {
			w.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(w).Encode(response)
			return
		}

		//todo Create a temporary file within our temp-images directory that follows a particular naming pattern
		//ctt Jadi, di bawah ini adalah pengaturan untuk penyimpanan file-nya
		//ctt "uploads" adalah nama folder tempat menyimpan gambar yang di-upload
		//ctt "image-*.png" adalah nama yang akan diterapkan kepada file yang di-upload. Sehingga, nama file yang aslinya akan tergantikan dengan nama "image-*" di mana * adalah satuan milisecond
		//ctt Jika kita ingin memaasukkan nama asli dari file yang kita upload, maka kita bisa menggunakan handler.Filename yang merupakan return ke dua dari r.FormFile di paling atas

		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(w).Encode(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		//ctt Setelah gambar tersimpan di folder uploads, pada tempFile.Name() akan tercatat nama folder sekaligus nama file-nya yaitu: uploads/image-2312432.png
		//ctt Maka dari itu, pada variabel filename di bawah kita split atau kita buang 8 huruf pertama agar kita hanya mendapatkan nama file-nya saja tanpa ada nama folder-nya juga
		//ctt Adapun 8 huruf pertama tersebut adalah uploads/

		data := tempFile.Name()
		filename := data[8:] // split uploads/

		//todo add filename to ctx

		//ctt Agar filname-nya dapat dimasukkan ke dalam database, maka kita simpan filename tersebut ke dalam context
		//ctt "dataFIle" adalah kata kuncinya sedangkan filename adalah value-nya.

		ctx := context.WithValue(r.Context(), "dataFile", filename)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
	//todo Selanjutnya, kita isikan routes CreateProduct
}
