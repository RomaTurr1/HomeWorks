package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)



type Employee struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Profession     string  `json:"profession"`
	Salary         float64 `json:"salary"`
	ExperienceYear int     `json:"experience_year"`
}

var db *sql.DB



func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./employees.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS employee (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		profession TEXT,
		salary REAL,
		experience_year INTEGER
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}




func getAllEmployees(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var employees []Employee

	for rows.Next() {
		var e Employee
		rows.Scan(&e.ID, &e.Name, &e.Profession, &e.Salary, &e.ExperienceYear)
		employees = append(employees, e)
	}

	json.NewEncoder(w).Encode(employees)
}


func getOneEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	row := db.QueryRow("SELECT * FROM employee WHERE id = ?", id)

	var e Employee
	err := row.Scan(&e.ID, &e.Name, &e.Profession, &e.Salary, &e.ExperienceYear)
	if err != nil {
		http.Error(w, "Not found", 404)
		return
	}

	json.NewEncoder(w).Encode(e)
}


func getNameExpYear(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT name, experience_year FROM employee")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	type Result struct {
		Name           string `json:"name"`
		ExperienceYear int    `json:"experience_year"`
	}

	var result []Result

	for rows.Next() {
		var r Result
		rows.Scan(&r.Name, &r.ExperienceYear)
		result = append(result, r)
	}

	json.NewEncoder(w).Encode(result)
}


func createEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	json.NewDecoder(r.Body).Decode(&e)

	_, err := db.Exec(
		"INSERT INTO employee(name, profession, salary, experience_year) VALUES(?,?,?,?)",
		e.Name, e.Profession, e.Salary, e.ExperienceYear,
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("Employee created"))
}



func main() {
	initDB()

	http.HandleFunc("/api/getallemployees", getAllEmployees)
	http.HandleFunc("/api/getone", getOneEmployee)
	http.HandleFunc("/api/getnameexpyear", getNameExpYear)
	http.HandleFunc("/api/createemployee", createEmployee)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}