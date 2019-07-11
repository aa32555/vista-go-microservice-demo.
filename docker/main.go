//////////////////////////////////////////////////////////////////
//								//
// Copyright (c) 2019 YottaDB LLC and/or its subsidiaries.	//
// All rights reserved.						//
//								//
//	This source code contains the intellectual property	//
//	of its copyright holder(s), and is made available	//
//	under a license.  If you do not know the terms of	//
//	the license, please stop and do not read further.	//
//								//
//////////////////////////////////////////////////////////////////

package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Item struct {
	Location string
	Name     string
	Sex      string
	Age      int
	CheckOut string
	CheckIn  string
}

const query string = `
SELECT h1.NAME, p1.NAME, p1.SEX, p1.AGE, v1.CHECK_OUT_DATE_TIME, v1.VISIT_ADMIT_DATE_TIME
FROM HOSPITAL_LOCATION h1
INNER JOIN VISIT v1 ON (v1.HOSPITAL_LOCATION = h1.HOSPITAL_LOCATION_ID)
INNER JOIN PATIENT p1 ON (v1.PATIENT_NAME = p1.PATIENT_ID) 
WHERE h1.NAME = 'ICU/CCU'
`

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		connStr := "host=vista port=1338 user=admin password=admin dbname=hello sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		rows, err := db.Query(query)
		if err != nil {
			panic(err)
		}
		items := []Item{}
		for rows.Next() {
			var loc, name, sex, checkout, checkin string
			var age int
			err = rows.Scan(&loc, &name, &sex, &age, &checkout, &checkin)
			if err != nil {
				panic(err)
			}
			items = append(items, Item{loc, name, sex, age, checkout, checkin})
		}
		json.NewEncoder(w).Encode(items)
	})

	handler := cors.Default().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
