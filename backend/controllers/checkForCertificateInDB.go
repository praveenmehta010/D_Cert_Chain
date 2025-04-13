package controllers

import (
	"context"
	"log"

	"github.com/praveenmehta010/Decentralized-Cert-Verification/database"
)

func checkForCertificateInDB(certificateID string) bool {
	fbRef := database.FirestoreClient.Collection("certificates")
	query := fbRef.Where("id", "==", certificateID).Limit(1)

	docs, errFB := query.Documents(context.Background()).GetAll()
	if errFB != nil {
		log.Println("Cannot find the certificate : ", errFB)
		return false
	}

	if len(docs) > 0 {
		// log.Println("Same Certificate already exists in DataBase")
		return true
	}
	// log.Println("Cannot find Certificate in DataBase")
	return false
}
