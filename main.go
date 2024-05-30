/*
package main

import (

	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"html/template"
	"net/http"

)

// Modelo

	type Student struct {
		ID    uint `gorm:"primaryKey"`
		Name  string
		Email string
	}

var db *gorm.DB
var err error

// Controladores

	func GetStudents(w http.ResponseWriter, r *http.Request) {
		var students []Student
		db.Find(&students)
		json.NewEncoder(w).Encode(students)
	}

	func CreateStudent(w http.ResponseWriter, r *http.Request) {
		var student Student
		json.NewDecoder(r.Body).Decode(&student)
		db.Create(&student)
		json.NewEncoder(w).Encode(student)
	}

	func ShowStudents(w http.ResponseWriter, r *http.Request) {
		var students []Student
		db.Find(&students)
		tmpl := template.Must(template.ParseFiles("students.html"))
		tmpl.Execute(w, map[string]interface{}{
			"Students": students,
		})
	}

	func main() {
		// Conectar a la base de datos SQLite
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrar el esquema
		db.AutoMigrate(&Student{})

		// Crear un nuevo enrutador
		router := mux.NewRouter()

		// Rutas API
		router.HandleFunc("/api/students", GetStudents).Methods("GET")
		router.HandleFunc("/api/students", CreateStudent).Methods("POST")

		// Rutas Web
		router.HandleFunc("/students", ShowStudents).Methods("GET")

		// Iniciar el servidor
		http.ListenAndServe(":8000", router)
	}
*/
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"myapp/models"
)

func main() {
	// Conectar a la base de datos SQLite
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Ejecutar la migración de la base de datos
	models.Migrate(db)

	// Configurar el enrutador
	r := SetupRouter(db)

	// Iniciar el servidor en el puerto 8080
	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
