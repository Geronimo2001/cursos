package main

import (
	"fmt"
	"myapp/database"
	"myapp/router"
	"net/http"
)

func main() {
	// Conectar a la base de datos
	db := database.ConnectDB()
	// defer db.Close()

	// Configurar el router
	r := router.SetupRouter(db)

	// Iniciar el servidor
	http.ListenAndServe(":8000", r)
	fmt.Println("Server is running at http://localhost:8000")
	//FIJATE QUE DIFERENCIA HAY ENTRE, COURSE CONTROLLER Y USER CONTROLLER UNO USA SERVICE
}
