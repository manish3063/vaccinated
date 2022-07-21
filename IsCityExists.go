package main

func isCityExist(city string) bool {
	var count int

	userSQL := "SELECT COUNT(*) FROM cities WHERE city=$1"

	row := DB.QueryRow(userSQL, city)

	err := row.Scan(&count)

	if err != nil {
		//log.Fatal(err)
	}

	if count == 1 {
		return true
	} else {
		return false
	}

}
