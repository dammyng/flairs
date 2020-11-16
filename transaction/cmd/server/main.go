package main

import (
	cmd "transaction/pkg/cmd/server"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)
func loadEnv() {
	log.Println("env loading...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func main() {
	loadEnv()
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
