package main

import (
	"fmt"
	"myapp/database"
	"myapp/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

type aplication struct {
	DSN          string
	DOmain       string
	DB           repository.databaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudiene   string
	CookieDomain string
}

func main() {
	// Conectar a la base de datos
	db := database.ConnectDB()
	// defer db.Close()

	// Configurar el router
	r := gin.Default()
	router.SetupUserRouter(r, db)
	router.SetupCourseRouter(r, db)

	// Iniciar el servidor
	http.ListenAndServe(":8000", r)
	fmt.Println("Server is running at http://localhost:8000")
	//FIJATE QUE DIFERENCIA HAY ENTRE, COURSE CONTROLLER Y USER CONTROLLER UNO USA SERVICE
}
