package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NurseDetails(c *gin.Context) {

	reqBody := Person{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	//checking either the city is exists or not
	if !isCityExist(reqBody.City) {
		res := gin.H{

			"Error": "City Does not Exist",
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := insertInDb(reqBody)
	//result, _ := create_user(reqBody)

	ress := gin.H{

		"data": result,
	}

	c.JSON(http.StatusOK, ress)
	return
}

func insertInDb(reqBody Person) (bool, error) {
	var err error
	DB, err = sql.Open("postgres", DB_DSN)
	if err != nil {
		log.Println("Failed to open a DB connection: ", err)
		return false, err
	}

	sqlStatement := `INSERT INTO persons(email, first_name,last_name,date_of_birth,sex)VALUES($1, $2,$3,$4,$5)`

	user, err := DB.Exec(sqlStatement, reqBody.Email, reqBody.FirstName, reqBody.LastName, reqBody.DateOfBirth, reqBody.Sex)

	if err != nil {
		log.Println("ERror in insert: ", err)
	}

	sqlStatement2 := `INSERT INTO nurses(email)VALUES($1)`

	user, err2 := DB.Exec(sqlStatement2, reqBody.Email)
	if err != nil {
		log.Println("ERror in insert: ", err2)
	}
	fmt.Println("kkk", sqlStatement2)
	fmt.Println(user)
	return true, err

}
