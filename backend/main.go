package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/blockchain"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/database"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/routes"
)

func main() {

	// Init Firebase
	database.InitFirebase()

	// Init blockchain connection
	blockchain.Init()

	// For routing
	router := mux.NewRouter()

	// Register routes
	routes.RegisterCertificateRoutes(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running with Firebase 🔥"))
	})

	//Log for the backend
	log.Println("Backend server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)(router)))

}
