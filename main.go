package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type OutputCar struct {
	ID         int
	Brand      string
	Model      string
	HorsePower int
}

func carGetterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		ID := r.URL.Path[17:]
		getFromDatabase(w, ID)

	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}

func getFromDatabase(w http.ResponseWriter, ID string) {
	db, err := sql.Open("mysql", "remoteuser:a@tcp(mariadb1:3306)/entrust")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var msg string = "SELECT * FROM car WHERE id = " + ID

	results, err := db.Query(msg)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	hasResults := false
	for results.Next() {
		hasResults = true
		var car OutputCar
		// for each row, scan the result into our tag composite object
		err = results.Scan(&car.ID, &car.Brand, &car.Model, &car.HorsePower)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Fprintln(w, "ID: "+strconv.Itoa(car.ID))
		fmt.Fprintln(w, "Brand: "+car.Brand)
		fmt.Fprintln(w, "Model: "+car.Model)
		fmt.Fprintln(w, "HorsePower: "+strconv.Itoa(car.HorsePower))
		log.Printf("Car selected: %+v\n", car)
	}
	if !hasResults {
		log.Println("There is no stored car with the ID:" + ID)
		fmt.Fprintln(w, "There is no stored car with the ID:"+ID)
	}

}

func main() {
	fmt.Println("GET Server Running...3")
	http.HandleFunc("/service/v1/cars/", carGetterHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
