package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func recipientHandler(c *gin.Context) {
	//result := Nurses{}
	//_ := c.Bind(&result)

	data := vaccinatedPrsn()

	ress := gin.H{

		"detail_of_all_person": data,
	}

	c.JSON(http.StatusBadRequest, ress)
	return
}

func vaccinatedPrsn() []vaccination {
	Data := []vaccination{}

	SQL := `SELECT "recipient", "vaccination_time", "vaccine", "site", "nurse","comments" FROM "vaccinations"`

	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	vaccinated := vaccination{}

	for rows.Next() {
		rows.Scan(&vaccinated.Recipient, &vaccinated.VaccinationTime, &vaccinated.Vaccine, &vaccinated.Site, &vaccinated.Nurse, vaccinated.Comments)

		Data = append(Data, vaccinated)

	}

	//fmt.Println(Data)
	return Data

}
