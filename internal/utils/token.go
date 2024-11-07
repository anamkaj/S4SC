package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AuthToken struct {
	AccessToken   string
	CookieSession string
	ClientTable   string
}

func GetToken() (AuthToken, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
		return AuthToken{}, err
	}

	access_token := os.Getenv("TOKEN")
	cookie_session := os.Getenv("COOKIES_SESSION")
	client_table := os.Getenv("CLIENT_TABLE")

	token := AuthToken{
		AccessToken:   access_token,
		CookieSession: cookie_session,
		ClientTable:   client_table,
	}

	return token, nil
}
