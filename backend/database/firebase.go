package database

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	FirestoreClient *firestore.Client
	FirebaseApp     *firebase.App
)

func InitFirebase() {

	ctx := context.Background()
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("❌ Could not load .env file: %v", err)
	}

	creds := map[string]string{
		"type":                        os.Getenv("TYPE"),
		"project_id":                  os.Getenv("PROJECT_ID"),
		"private_key_id":              os.Getenv("PRIVATE_KEY_ID"),
		"private_key":                 strings.ReplaceAll(os.Getenv("PRIVATE_KEY"), `\n`, "\n"),
		"client_email":                os.Getenv("CLIENT_EMAIL"),
		"client_id":                   os.Getenv("CLIENT_ID"),
		"auth_uri":                    os.Getenv("AUTH_URI"),
		"token_uri":                   os.Getenv("TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("AUTH_PROVIDER_X509_CERT_URL"),
		"client_x509_cert_url":        os.Getenv("CLIENT_X509_CERT_URL"),
		"universe_domain":             os.Getenv("UNIVERSE_DOMAIN"),
	}

	credJSON, err := json.Marshal(creds)
	if err != nil {
		log.Fatalf("Failed to marshal Firebase credentials: %v", err)
	}
	opt := option.WithCredentialsJSON(credJSON)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Firestore client: %v", err)
	}

	log.Println("🔥 Connected to Firebase Firestore")
	FirestoreClient = client
	FirebaseApp = app

}

// package database

// import (
// 	"context"
// 	"log"

// 	"cloud.google.com/go/firestore"
// 	firebase "firebase.google.com/go/v4"
// 	"google.golang.org/api/option"
// )

// var (
// 	FirestoreClient *firestore.Client
// 	FirebaseApp     *firebase.App
// )

// func InitFirebase() {

// 	ctx := context.Background()

// 	sa := option.WithCredentialsFile("config/serviceAccountKey.json")
// 	app, err := firebase.NewApp(ctx, nil, sa)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize Firebase app: %v", err)
// 	}

// 	client, err := app.Firestore(ctx)
// 	if err != nil {
// 		log.Fatalf("Failed to initialize Firestore client: %v", err)
// 	}

// 	log.Println("🔥 Connected to Firebase Firestore")
// 	FirestoreClient = client
// 	FirebaseApp = app

// }

// // CloseFirestore closes the Firestore connection on server shutdown
// func CloseFirestore() {
// 	if FirestoreClient != nil {
// 		FirestoreClient.Close()
// 		log.Println("🔌 Firestore connection closed.")
// 	}
// }
