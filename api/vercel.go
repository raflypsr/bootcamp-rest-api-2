package api

import (
	"backend-tugas-reactjs/config"
	"backend-tugas-reactjs/docs"
	"backend-tugas-reactjs/routes"
	"backend-tugas-reactjs/utils"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	App *gin.Engine
)

func init() {
	App = gin.New()

	environment := utils.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	docs.SwaggerInfo.Title = "Book REST API"
	docs.SwaggerInfo.Description = "This is REST API Book."
	docs.SwaggerInfo.Version = "1.0"
	if environment == "development" {
		docs.SwaggerInfo.Host = utils.Getenv("HOST", "localhost:9090")
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
	} else {
		docs.SwaggerInfo.Host = os.Getenv("HOST")
		docs.SwaggerInfo.Schemes = []string{"https"}
	}
	db := config.ConnectDataBase()

	routes.SetupRouter(db, App)
}

// Entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	App.ServeHTTP(w, r)
}
