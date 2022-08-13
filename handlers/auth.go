package handlers

import (
	authdto "dumbmerch/dto/auth"
	dto "dumbmerch/dto/result"
	"dumbmerch/models"
	"dumbmerch/pkg/bcrypt"
	"dumbmerch/repositories"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

//ctt struct di atas memberitahukan bahwa handler ini akan meneruskna ke repositroy AuthRepository

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

//ctt untuk Membawa Connection melalui Routes

func (h *handlerAuth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(authdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	//todo di bawah ini adalah cara untuk meng-enkripsi password

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: password,
	}

	//ctt Sebenarnya, untuk Register ini  100% sama dengan handler pada Create User. Yang membedakannya adalah pada password-nya di mana pada handler CreateUser akan langsung mengambil data password dari request-nya sedangkan pada Register ini password dari request akan dienkripsi terlebih dahulu sebagaimana di atas

	data, err := h.AuthRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

//todo Setelah ini jangan lupa untuk menambahkan Routes-nya.
