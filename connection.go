package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	// TODO fill this in directly or through environment variable
	// Build a DSN e.g. postgres://username:password@url.com:5432/dbName
	DB_DSN = "postgres://postgres:manish@localhost:5432/vaccine?sslmode=disable"
)

var (
	DB *sql.DB
)

type Nurses struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}
type Person struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Sex         string `json:"sex"`
	City        string `json:"city"`
}

type Delete struct {
	Email string `json:"email" binding:"required"`
}
type vaccination struct {
	Recipient       string `json:"recipient"`
	VaccinationTime string `json:"vaccine_time"`
	Vaccine         string `json:"vaccine"`
	Site            string `json:"site"`
	Nurse           string `json:"nurse"`
	Comments        string `json:"comments"`
}

type nursevaccinated struct {
	Nurse           string `json:"nurse"`
	VaccinationTime string `json:"vaccination_time"`
	Recipient       string `json:"recipient"`
	Vaccine         string `json:"vaccine"`
	Site            string `json:"site"`
}

func main() {
	createDBConnection()
	r := gin.Default()
	setupRoutes(r)
	r.Run() //

}

func setupRoutes(r *gin.Engine) {

	r.GET("/nurses", createHandler)
	r.GET("/person", PersonHandler)
	r.GET("/recipient", recipientHandler)
	r.GET("/vaccinated", vaccinatedHandler)
	r.POST("/addnursedetail", NurseDetails)
	r.DELETE("/deleteuser", DeleteUser)

}

func createDBConnection() {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	} else {
		fmt.Println("Connected to database")

	}
	// defer DB.Close()
}
