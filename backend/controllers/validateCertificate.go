package controllers

import (
	"context"
	"log"

	"github.com/praveenmehta010/Decentralized-Cert-Verification/database"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/models"
)

func ValidateViaID(certificateID string) (newCertificateHash string, newCertificateID string) {
	fbRef := database.FirestoreClient.Collection("certificates")
	query := fbRef.Where("id", "==", certificateID).Limit(1)

	docs, errFB := query.Documents(context.Background()).GetAll()
	if errFB != nil {
		log.Println("Cannot find the certificate : ", errFB)
		return
	}

	var cert models.Certificate
	err := docs[0].DataTo(&cert)

	if err != nil {
		log.Println("Mapping Error:", err)
		return
	}

	newCertificateHash = GenerateCertificateHash(&cert)
	newCertificateID = GenerateCertificateID(&cert)
	// log.Println("New certificate hash :- ", newCertificateHash)
	// log.Println("New certificate ID :- ", newCertificateID)

	return newCertificateHash, newCertificateID

}
