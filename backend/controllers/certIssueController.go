package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/praveenmehta010/Decentralized-Cert-Verification/blockchain"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/database"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/models"
)

func IssueCertificate(w http.ResponseWriter, r *http.Request) {

	var cert models.Certificate //data Model var
	if err := json.NewDecoder(r.Body).Decode(&cert); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//********************************Data to be saved***********************************************************
	cert.DateIssued = time.Now()
	cert.Revoked = "false"

	issuedCertificatHash := GenerateCertificateHash(&cert) //Hash Generation for the certificate
	issuedCertificateID := GenerateCertificateID(&cert)    //ID Generation for the certificate
	// log.Println("Old certificate hash :- ", issuedCertificatHash)
	// log.Println("Old certificate id :- ", issuedCertificateID)

	cert.ID = issuedCertificateID
	certHashBytes32, err := blockchain.HexToBytes32(issuedCertificatHash)
	if err != nil {
		http.Error(w, "Invalid certificate hash format", http.StatusInternalServerError)
		return
	}

	// checking for the same certificate in the DataBase
	certificateExistsDB := checkForCertificateInDB(issuedCertificateID)
	if certificateExistsDB {
		log.Println("Can not generate the same certificate again, Error - DataBase")
		return
	}

	// checking for the same certificate in the BlockChain
	certificateExistsBC := checkForCertificateInBC(issuedCertificatHash, certHashBytes32)
	if certificateExistsBC {
		log.Println("Can not generate the same certificate again, Error - BlockChain")
		return
	}

	//********************************Data storing in DB***********************************************************
	_, _, errIssueFB := database.FirestoreClient.Collection("certificates").Add(context.Background(), cert)
	if errIssueFB != nil {
		http.Error(w, "Failed to store certificate in Firestore", http.StatusInternalServerError)
		return
	}

	//********************************Certificate hash storing in  blockchain***********************************************************
	err = blockchain.StoreCertificateHash(issuedCertificatHash, certHashBytes32)
	if err != nil {
		log.Printf("Blockchain error: %v", err)
		http.Error(w, "Blockchain store failed", http.StatusInternalServerError)
		return
	}

	//********************************Details of Issuer***********************************************************
	uid := r.Context().Value("uid").(string)
	email := r.Context().Value("email").(string)
	fmt.Printf("Certificate issued by :- %s\n UserID :- (%s)\n", email, uid)

	//*******************************************************************************************
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":       "Certificate issued successfully",
		"CertificateID": issuedCertificateID,
		"Name":          cert.StudentName,
	})
}
