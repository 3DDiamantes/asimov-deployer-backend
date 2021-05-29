package main

import "asimov-deployer-backend/internal/http"

func main() {
	router := http.InitRouter()
	router.Run(":8080")
}
