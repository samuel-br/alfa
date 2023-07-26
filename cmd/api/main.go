package main

import (
	"alfa/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	advencePayApiService := api.NewAdvanceApiService()

	port := os.Getenv("PORT")

	router := gin.Default()

	router.POST("/perform_advance/", advencePayApiService.AdvancePay)

	router.Run(port)
}
