package controllers

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/praveenmehta010/Decentralized-Cert-Verification/database"
)

func updateCertificateField(certificateID string, fieldName string, newValue interface{}) error {
	ctx := context.Background()
	fbRef := database.FirestoreClient.Collection("certificates")

	query := fbRef.Where("id", "==", certificateID).Limit(1)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		log.Printf("🔴 Error finding certificate: %v", err)
		return err
	}

	docRef := docs[0].Ref
	_, err = docRef.Update(ctx, []firestore.Update{
		{Path: fieldName, Value: newValue},
	})

	if err != nil {
		log.Printf("🔴 Error updating certificate: %v", err)
		return err
	}

	log.Printf("✅ Certificate field :- %v \n updated successfully", fieldName)
	return nil
}
