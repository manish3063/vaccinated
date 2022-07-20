package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func vaccinatedHandler(c *gin.Context) {
	//result := Nurses{}
	//_ := c.Bind(&result)

	data := AllVacinated()

	ress := gin.H{

		"detail_of_all_person": data,
	}

	c.JSON(http.StatusBadRequest, ress)
	return
}

func AllVacinated() []nursevaccinated {
	Data := []nursevaccinated{}

	//SQL := `SELECT "recipient", "vaccination_time", "vaccine", "site", "nurse","comments" FROM "vaccinations"`

	SQL := `SELECT vaccinations.nurse,vaccination_time, recipient,vaccine, site from vaccinations inner join nurses 
    on vaccinations.recipient = nurses.email `

	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	vaccinated := nursevaccinated{}

	for rows.Next() {
		rows.Scan(&vaccinated.Nurse, &vaccinated.VaccinationTime, &vaccinated.Recipient, &vaccinated.Vaccine, &vaccinated.Site)

		Data = append(Data, vaccinated)

	}

	//fmt.Println(Data)
	return Data

}
