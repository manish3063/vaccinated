package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func PersonHandler(c *gin.Context) {
	//result := Nurses{}
	//_ := c.Bind(&result)

	data := allperson()

	ress := gin.H{

		"detail_of_all_person": data,
	}

	c.JSON(http.StatusBadRequest, ress)
	return
}

func allperson() []Person {
	Data := []Person{}

	//SQL := `SELECT "email", "first_name", "last_name", "date_of_birth", "sex" FROM "persons"`

	SQL := `SELECT persons.email,first_name, last_name,date_of_birth, sex from persons inner join vaccinations
    on persons.email = vaccinations.recipient `

	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	detailperson := Person{}

	for rows.Next() {
		rows.Scan(&detailperson.Email, &detailperson.FirstName, &detailperson.LastName, &detailperson.DateOfBirth, &detailperson.Sex)

		Data = append(Data, detailperson)

	}

	//fmt.Println(Data)
	return Data

}
