package main

import (
	"fmt"
	"log"
	"myapp/database"
	"myapp/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = 8000

func main() {

	// Conectar a la base de datos
	db := database.ConnectDB()
	// defer db.Close()

	// Configurar el router
	r := gin.Default()
	router.SetupUserRouter(r, db)
	router.SetupCourseRouter(r, db)

	// Iniciar el servidor
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running at http://localhost:8000")
}
