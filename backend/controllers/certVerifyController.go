package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"

	"github.com/praveenmehta010/Decentralized-Cert-Verification/blockchain"
)

func VerifyCertificate(w http.ResponseWriter, r *http.Request) {

	var req struct {
		CertificateID string `json:"certificateHash"`
		NoOfIssue     string `firestore:"noOfIssue"` // nunmber of times the same certificate has been issued
		Course        string `firestore:"course"`    // Course or event name
		StudentEmail  string `firestore:"studentEmail"`
		StudentName   string `firestore:"studentName"` // Name of the student

	}

	certificateID := ""

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || (req.CertificateID == "" && (req.Course == "" || req.NoOfIssue == "" || req.StudentEmail == "" || req.StudentName == "")) {

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.CertificateID != "" {
		// log.Println("I have certificate ID")
		certificateID = req.CertificateID
	} else {
		// log.Println("I have certificate details now creating ID")

		idRawData := req.StudentEmail + "Magical Salt :- jnj864hyf%$7VthI" + req.Course + req.NoOfIssue + req.StudentName
		certIDHash := sha256.Sum256([]byte(idRawData))
		certificateID = hex.EncodeToString(certIDHash[:])
	}

	if !checkForCertificateInDB(certificateID) {
		log.Println("Certificate Does not Exists")
		json.NewEncoder(w).Encode(map[string]bool{"valid": false})
		return
	}

	newCertificateHash, newCertificateID := ValidateViaID(certificateID)
	newCertHashBytes32, err := blockchain.HexToBytes32(newCertificateHash)
	// log.Println("New Certificate Hash :- ", newCertificateHash)
	// log.Println("New Certificate Hash :- ", newCertificateID)
	if err != nil {
		http.Error(w, "Invalid certificate hash format", http.StatusInternalServerError)
		return
	}

	checkForCertificateInDB := checkForCertificateInDB(newCertificateID)
	if !checkForCertificateInDB {
		log.Println("❌ Certificate verification failed, DataBase Error")
	}

	checkForCertificateInBC := checkForCertificateInBC(newCertificateHash, newCertHashBytes32)
	if !checkForCertificateInBC {
		log.Println("❌ Certificate verification failed, Blockchain Error")
	}

	if !checkForCertificateInDB || !checkForCertificateInBC {
		log.Println("❌ Certificate Is Not Valid")
		json.NewEncoder(w).Encode(map[string]bool{"valid": false})
		return
	}

	log.Println("✅ Certificate Is Valid")
	json.NewEncoder(w).Encode(map[string]bool{"valid": true})
}
