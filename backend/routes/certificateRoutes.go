package routes

import (
	"github.com/gorilla/mux"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/controllers"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/middleware"
)

func RegisterCertificateRoutes(router *mux.Router) {
	// Public route
	router.HandleFunc("/verify", controllers.VerifyCertificate).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/").Subrouter()
	protected.Use(middleware.FirebaseAuthMiddleware)

	protected.HandleFunc("/issue", controllers.IssueCertificate).Methods("POST")
	protected.HandleFunc("/revoke", controllers.RevokeCertificate).Methods("POST")
}
