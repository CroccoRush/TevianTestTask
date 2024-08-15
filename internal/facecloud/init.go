package facecloud

import (
	"fmt"
	"log"
	"os"
)

var (
	facecloudAPI   string
	email          string
	password       string
	accessToken    string
	detectEndpoint string
	loginEndpoint  string
)

func init() {

	facecloudAPI = os.Getenv("FACECLOUD_API_URL")
	if facecloudAPI == "" {
		facecloudAPI = "https://backend.facecloud.tevian.ru/api"
	}
	email = os.Getenv("FACECLOUD_EMAIL")
	password = os.Getenv("FACECLOUD_PASSWORD")
	//email = "kiselyovvld@mail.ru"
	//password = "Lesstenpound@159951"

	detectEndpoint = fmt.Sprintf("%s/v1/detect", facecloudAPI)
	loginEndpoint = fmt.Sprintf("%s/v1/login", facecloudAPI)

	response, err := Login()
	if err != nil {
		log.Fatal(err)
	}
	accessToken = response.Data.AccessToken

}
