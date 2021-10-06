package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first Go API")
	//selectCities()
	//loadCsv()
	//insert()
}
func cityId(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service to find City by id")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}

	f, err := os.Open("data/data.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if line[0] == id {
			fmt.Fprintf(w, "\nId     : "+line[0]+"\nCity   : "+line[1]+"\nCountry: "+line[2])
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/cityid/{id}", cityId)
	//	router.HandleFunc("/loadcsv", loadCsv)
	fmt.Println("Server is runnig...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
