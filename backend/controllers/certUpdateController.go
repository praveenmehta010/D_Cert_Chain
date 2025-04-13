package controllers

import (
	"encoding/json"
	"net/http"
)

func RevokeCertificate(w http.ResponseWriter, r *http.Request) {

	var req struct {
		CertificateHash string `json:"certificateHash"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.CertificateHash == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	checkForCertificateInDB := checkForCertificateInDB(req.CertificateHash)
	if !checkForCertificateInDB {
		return
	}

	updateCertificateField(req.CertificateHash, "revoked", "true")

	//msg in postman
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Certificate Updated successfully",
	})
}
