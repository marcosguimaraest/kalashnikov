package api

import (
	"os"

	"github.com/joho/godotenv"
)

var UriApi string
var ApiKey string
var ApiToken string
var ApiIDDefaultList string
var APiIDDefaultBoard string

func InitializeApi() {
	godotenv.Load("secrets.env")
	UriApi = os.Getenv("TRELLOAPIURL")
	ApiKey = os.Getenv("TRELLOAPIKEY")
	ApiToken = os.Getenv("TRELLOAPITOKEN")
	APiIDDefaultBoard = os.Getenv("APIIDDEFAULTBOARD")
	ApiIDDefaultList = os.Getenv("APIIDDEFAULTLIST")
}
