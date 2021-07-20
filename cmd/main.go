package main

import "asimov-deployer-backend/internal/http"

func main() {
	router := http.InitRouter()
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
