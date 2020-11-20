package model

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
)

var client *db.Client
var ctx context.Context

func init(){
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://project-antrian.firebaseio.com/",
	}

	opt := option.WithCredentialsFile("project-antrian-firebase-adminsdk-56ko9-0d789d2a08.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err =  app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
}