package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"modernc.org/sqlite"
)

type UserModelInterface interface {
	Insert(name string, email string, password string) error
	Authenticate(email string, password string) (int, error)
	Exists(id int) (bool, error)
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name string, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES (?, ?, ?, DATETIME("now"))`
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		var sqliteError *sqlite.Error
		if errors.As(err, &sqliteError) {
			if sqliteError.Code() == 2067 && strings.Contains(sqliteError.Error(), "users.email") {
				return ErrDuplicateEmail
			}
		}
		return err
	}

	// id, err := result.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }

	return nil
}

func (m *UserModel) Authenticate(email string, password string) (int, error) {
	var id int
	var hashedPassword []byte

	stmt := `SELECT id, hashed_password FROM users WHERE email = ?`
	err := m.DB.QueryRow(stmt, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	return id, nil
	// return Snippet{}, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
	var exists bool
	stmt := `SELECT EXISTS(SELECT true FROM users where id = ?)`
	err := m.DB.QueryRow(stmt, id).Scan(&exists)
	return exists, err
}
