package main

import "fmt"

// user
type User struct {
	ID        int    `json:"id,-"`
	Login     string `json:"login"`
	Password  string `json:"password,-"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Role      string `json:"role,-"`
}

func (u User) Create() error {
	row := connection.QueryRow(`INSERT INTO userr (login, password, firstname, lastname, role) values ($1, $2, $3, $4, 'manager') returning id`, u.Login, u.Password, u.FirstName, u.LastName)
	e := row.Scan(&u.ID)
	if e != nil {
		return e
	}

	fmt.Println("create new user with id: ", u.ID)

	return nil
}

func (u User) Select() error {
	row := connection.QueryRow(`select role, firstname, lastname from userr where login=$1 and password=$2`, u.Login, u.Password)
	e := row.Scan(&u.Role, &u.FirstName, &u.LastName)
	if e != nil {
		return e
	}

	fmt.Println("authorizate new user with role: ", u.Role)

	return nil
}
