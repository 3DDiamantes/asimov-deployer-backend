package main

import "asimov-deployer-backend/internal/http"

func main() {
	r := http.InitRouter()
	r.Run(":9090")
}
