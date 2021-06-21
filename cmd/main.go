package main

import "asimov-deployer-backend/internal/http"

func main() {
	router := http.InitRouter()
	router.Run("localhost:8080")
}
