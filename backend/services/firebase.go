package services

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
)

var FirebaseAuth *auth.Client
var FirestoreClient *firestore.Client

func InitFirebase() {
	opt := option.WithCredentialsFile("firebase-service-account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase: %v\n", err)
	}

	FirebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing firebase auth: %v\n", err)
	}

	FirestoreClient, err = app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
	}
	log.Println("Firebase initialized successfully")
}
