package dao

import (
	"database/sql"
	_"github.com/pkg/errors"
)

var (
	DB *sql.DB
)

type User struct {
	Id		int
	Name	string
}

type Dao struct {
}

func NewDao() *Dao {
	return &Dao{}
}

func (d *Dao) FindUserById(userID int) (u User, err error) {
	//err = DB.QueryRow("SELECT name FROM users WHERE id = ?",userID).Scan(&u.Name)
	err = sql.ErrNoRows
	//err = sql.ErrConnDone
	///return u, errors.Wrap(err, "Dao miss match!")
	return u, err
}