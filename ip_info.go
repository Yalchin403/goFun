package main

import (
	"fmt"
	ipdata "github.com/ipdata/go"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("API_KEY")
	ipd, _ := ipdata.NewClient(API_KEY)

	var ip string
	fmt.Println("Enter the ip address:")
	fmt.Scanln(&ip)
	data, err := ipd.Lookup(ip)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%s (%s)\n", data.IP, data.ASN)
}
