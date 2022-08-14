package main

import (
	"dumbmerch/database"
	"dumbmerch/pkg/mysql"
	"dumbmerch/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// import "godotenv" here ...
)

func main() {

	// Init godotenv here ...

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//todo Initialization "uploads" folder to public here ...
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code
	//ctt Kita buatkan kodingan seperti di atas agar gambarnya bisa ditampilkan
	//ctt "/uploads/" yang pertama adalah yang akan disematkan ke domain
	//ctt "/uploads" yang ke dua menunjukkan folder mana pada direktori kita

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
