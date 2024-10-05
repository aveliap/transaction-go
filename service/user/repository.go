package user

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aveliap/transaction-go/types"
)

type Repo struct{
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo  {
	return &Repo{db:db}
}

func (repo *Repo) GetUserByEmail(email string) (*types.User, error) {
	rows, err := repo.db.Query("SELECT * FROM users WHERE email=?", email)

	if err != nil {
		return nil,err
	}
	u := new(types.User)
	for rows.Next(){
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil,err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}


func (repo *Repo) GetUserByID (id uint) (*types.User, error)  {
	rows, err := repo.db.Query("SELECT * FROM users WHERE id=?", id)

	if err != nil {
		return nil,err
	}
	u := new(types.User)
	for rows.Next(){
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil,err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (repo *Repo) CreateUser (user types.User) error  {
	fmt.Printf("firstName: %s, lastName: %s, email: %s, password: %s\n", user.FirstName, user.LastName, user.Email, user.Password)
	log.Printf("Executing query: INSERT INTO users (\"firstName\", \"lastName\", \"email\", \"password\") VALUES (%s, %s, %s, %s)",
    user.FirstName, user.LastName, user.Email, user.Password)

	_, err := repo.db.Exec("INSERT INTO users (\"firstName\", \"lastName\", \"email\", \"password\") VALUES ($1, $2, $3, $4)",
				user.FirstName,user.LastName,user.Email,user.Password)
	if err != nil {
		return err
	}
	return nil
}

func scanRowIntoUser(row *sql.Rows)(*types.User, error)  {
	user := new(types.User)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil,err
	}

	return user, nil
}
