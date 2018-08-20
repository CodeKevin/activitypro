package models

import (
	db "activitypro/database"
)

type Person struct {
	ID        int    `json:"id" from:"id"`
	FirstName string `json:"firstname" from:"firstname"`
	LastName  string `json:"lastname" from:"lastname"`
}

func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO person(first_name,last_name) VALUE (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, first_name, last_name FROM person")
	defer rows.Close()

	if err != nil {
		return
	}
	for rows.Next() {
		var person Person
		rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}
