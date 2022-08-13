package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerAuth(userRepository)

	//ctt di handler-nya kita masih mengacu ke userRepository
	//ctt Tapi, jika kita ingin mengacu ke authRepository-nya pun bisa
	/*
		authRepository := repositories.RepositoryAuth(mysql.DB)
		h := handlers.HandlerAuth(authRepository)
	*/
	r.HandleFunc("/register", h.Register).Methods("POST")

	//ctt Untuk login akan sekalian dibahas pada JWT
}
