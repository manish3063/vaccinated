package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {

	reqBody := Delete{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := UserDelete(reqBody.Email)
	if result == 0 {
		ress := gin.H{
			"message": "error",
		}
		c.JSON(http.StatusBadRequest, ress)
		c.Abort()
		return
	}
	ress := gin.H{
		"message": "succefully deleted",
		"result":  result,
	}
	c.JSON(http.StatusOK, ress)

}

func UserDelete(email string) int {
	a := 1
	//msg := ""
	sqlStatement := `DELETE FROM nurses where email = $1`

	_, err := DB.Exec(sqlStatement, email)

	if err != nil {
		log.Println("ERror in deleting1: ", err)
		a = 0
		return a
	}

	sqlStatement2 := `DELETE FROM persons where email = $1`

	_, err2 := DB.Exec(sqlStatement2, email)

	if err2 != nil {
		log.Println("ERror in deleting2: ", err2)
		a = 0
		return a
	}

	return a

}
