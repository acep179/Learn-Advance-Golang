//todo Create package handlers here ...
package handlers

//todo Import the required packages here ...
import (
	productdto "dumbmerch/dto/product"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

//todo Declare handlerProduct  struct here ...
type handlerProduct struct {
	ProductRepository repositories.ProductRepository
}

//todo HandlerProduct function here ...
func HandlerProduct(ProductRepository repositories.ProductRepository) *handlerProduct {
	return &handlerProduct{ProductRepository}
}

//ctt Pada handler ini, keseluruhannya sama saja dan tidak adanya perubahan dan perbedaan karena kita cukup mengaturnya di Repositroy-nya saja.
//ctt Ada sedikit modifikasi dan catatan pada Method CreateProduct

//todo Create FindProducts method here ...
func (h *handlerProduct) FindProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := h.ProductRepository.FindProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: products}
	json.NewEncoder(w).Encode(response)
}

//todo Create GetProduct method here ...
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

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseProduct(product)}
	json.NewEncoder(w).Encode(response)
}

//todo Create CreateProduct method here ...
func (h *handlerProduct) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(productdto.ProductRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

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
		Image:  request.Image,
		Qty:    request.Qty,
		UserID: request.UserID,
	}

	product, err = h.ProductRepository.CreateProduct(product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	product, _ = h.ProductRepository.GetProduct(product.ID)
	//ctt Untuk menampilkan data yang baru saja kita buat, maka kita laksanakan MethodGet dari Repository seperti di atas

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: product}
	json.NewEncoder(w).Encode(response)
}

//todo Create convertResponseProduct function here ...
func convertResponseProduct(u models.Product) models.ProductResponse {
	return models.ProductResponse{
		ID: u.ID,
		//ctt Agar ketika mengirim response yang kita kirim adalah ID dari Data yang baru kita inputkan, maka kita atur sebagaimana di atas.
		Name:     u.Name,
		Desc:     u.Desc,
		Price:    u.Price,
		Image:    u.Image,
		Qty:      u.Qty,
		User:     u.User,
		Category: u.Category,
	}
}

//todo Sekarang, kta urus terkait routes-nya
