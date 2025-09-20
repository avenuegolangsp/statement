package main

import (
	routes "bankapp/internal/http"
)

func main() {
	r := routes.SetupRoutes()

	r.Run(":8080")
}
