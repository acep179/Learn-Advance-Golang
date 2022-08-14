package handlers

import (
	productdto "dumbmerch/dto/product"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

//todo Create `path_file` Global variable here ...
var path_file = "http://localhost:5000/uploads/"

//ctt variabel di atas akan kita gunakan ketika kita ingin menampilkan data dari hasil FindProducts dan GetProduct agar isi dari "image" tidak hanya nama gambarnya saja sebagaimana di database, tetapi juga didahului oleh path-nya

func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//todo Create Embed Path File on Image property here ...
	for i, p := range products {
		products[i].Image = path_file + p.Image
	}

	//ctt Karena data yang akan ditampilkan itu banyak yaitu berupa Array Object. Maka, kita perlu melakukan looping agar path_file tersebut terpasang pada setiap data yang akan ditampilkan

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var product models.Product
	product, err := h.ProductRepository.GetProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//todo Create Embed Path File on Image property here ...
	product.Image = path_file + product.Image
	//ctt Pada GetProduct ini, kita tidak perlu melakukan looping karena data yang ditampilkan hanya ada satu dan bukan merupakan Array Object, melainkan hanya Object saja.

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(product)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// Get dataFile from midleware and store to filename variable here ...
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code
	//ctt Data dari Context, yaitu "dataFile" di-convert menjadi string agar dapat digunakan nantinya

	//ctt Pada handler ini, yang tadinya kita menggunakan json, sekarang kita menggunakan form data. Maka dari itu, di bawah ini kita membutuhkan yang namanya FormValue()

	price, _ := strconv.Atoi(r.FormValue("price"))
	qty, _ := strconv.Atoi(r.FormValue("qty"))
	category_id, _ := strconv.Atoi(r.FormValue("category_id"))

	//ctt Ketiga value di atas kita dapatkan dari form (bukan json). Sehingga, kita  harus meng-convert value-nya dari string menjadi integer
	//ctt Sedangkan yang memang tipe data yang dibutuhkan itu berupa string, maka tidak perlu kita convert sebagaimana name dan Desc di bawah

	request := productdto.ProductRequest{
		Name:       r.FormValue("name"),
		Desc:       r.FormValue("desc"),
		Price:      price,
		Qty:        qty,
		CategoryID: category_id,
	}

	//! Perlu diingat bahwa ketika kita sudah mulai melakukan integrasi, pastikan bahwa setiap Method pada setiap Handler sudah "melakukan" pengambilan "inputan" dari FormValue(), bukan hanya "json". Karena nantinya di sisi Frontend itu akan mengirimkan data berupa FormValue(), bukan lagi berbentuk "json"
	//! Kendati demikian, kembali lagi tergantung di sisi Frontend-nya apakah data-data dari Form itu akan diubah ke "json" dulu atau tidak. Karena untuk Form yang terdapat uploadFile-nya itu tidak mungkin diubah ke dalam bentuk "json"

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product := models.Product{
		Name:   request.Name,
		Desc:   request.Desc,
		Price:  request.Price,
		Image:  filename,
		Qty:    request.Qty,
		UserID: userId,
	}

	// err := mysql.DB.Create(&product).Error
	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)
}

//todo Tahapan selanjutnya setelah mengatur handler adalah mengatur path file-nya. Untuk melakukan hal tersebut, kita beralih terlebih dahulu ke file main.go

func convertResponseProduct(u models.Product) models.ProductResponse {
	return models.ProductResponse{
		Name:     u.Name,
		Desc:     u.Desc,
		Price:    u.Price,
		Image:    u.Image,
		Qty:      u.Qty,
		User:     u.User,
		Category: u.Category,
	}
}
