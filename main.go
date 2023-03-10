package main

import (
	"fmt"
	acronis "github.com/gorobot-nz/acronis/client"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env")
	}
}

func main() {
	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	datacenterUrl := os.Getenv("DATACENTER_URL")

	acronisClient, err := acronis.NewAcronisClient(clientId, clientSecret, datacenterUrl)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	acronisClient.SwitchToProduction("19f9321d-4bd6-420b-85b1-5cca6e940b98")
}
