package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"products_api/internal/configuration"
	"products_api/internal/handlers"
	"products_api/internal/repository"
	"products_api/internal/services"
	"strings"

	viper "github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := configuration.Configuration{
		PORT:   viper.GetString("PORT"),
		DB_URL: viper.GetString("DB_URL"),
	}

	dbConn, err := configuration.InitDb(config.DB_URL)
	if err != nil {
		log.Fatal("Failed to initialize : ", err)
	}
	defer dbConn.Close()

	productRepo := repository.NewProductRepository(dbConn)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/products", productHandler.GetAllProducts)

	listener, err := net.Listen("tcp", ":"+config.PORT)

	if err != nil {
		fmt.Println("Error Listening", err.Error())
		return
	}

	fmt.Println("Serve on port ", config.PORT)

	http.Serve(listener, mux)

}
