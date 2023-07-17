package firebase

import (
  "fmt"
  "context"

  firebase "firebase.google.com/go"
  "firebase.google.com/go/auth"
  "google.golang.org/api/option"
)

type firebaseAppInterface interface {
	VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error)
}

type firebaseApp struct {
	*firebase.App
}

func InitFirebaseApp() (*firebaseApp, error) {
  opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
    fmt.Errorf("error initializing app: %v\n", err)
  }
	return &firebaseApp{app}, nil
}

func (app *firebaseApp) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := app.Auth(ctx)
    if err != nil {
      fmt.Errorf("error getting Auth client: %v\n", err)
    }
	token, err := client.VerifyIDToken(ctx, idToken)
    if err != nil {
      fmt.Errorf("error verifying ID token: %v\n", err)
    }
	return token, nil
}