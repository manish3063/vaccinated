package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createHandler(c *gin.Context) {
	//result := Nurses{}
	//_ := c.Bind(&result)

	data := allnurces()

	ress := gin.H{

		"detail_of_all_nurces": data,
	}

	c.JSON(http.StatusBadRequest, ress)
	return
}

func allnurces() []Nurses {
	Data := []Nurses{}

	SQL := `SELECT "email" FROM nurses`

	rows, err := DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	nurse := Nurses{}

	for rows.Next() {
		rows.Scan(&nurse.Email)

		Data = append(Data, nurse)

	}

	//fmt.Println(Data)
	return Data

}
