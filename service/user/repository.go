package user

import (
	"database/sql"
	"fmt"

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

func (repo *Repo) GetUserByID (id uint) (*types.User, error)  {
	return nil,nil	
}

func (repo *Repo) CreateUser (user types.User) error  {
	return nil
}