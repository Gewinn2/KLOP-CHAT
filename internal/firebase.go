package internal

import (
	"context"
	"log"

	AstraLinux_TCPChat_hackathon "Astra_Linux_chat"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebase() *firebase.App {
	opt := option.WithCredentialsFile(AstraLinux_TCPChat_hackathon.OptPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}
