package storage

import (
	"context"
)

// InsertUser inserts a user into the database.
func (db *DB) InsertUser(u string) error {
	_, err := db.Exec(context.Background(), "INSERT INTO users (token, tg_id) VALUES ($1, $2)", u)
	return err
}

// SelectUserByID returns a user by id.
func (db *DB) SelectUserByID(id int64) (string, error) {
	var u string
	err := db.QueryRow(context.Background(), "SELECT token, tg_id FROM users WHERE tg_id = $1", id).Scan(&u)
	return u, err
}

// SelectUserByToken returns a user by token.
func (db *DB) SelectUserByToken(token string) (string, error) {
	var u string
	err := db.QueryRow(context.Background(), "SELECT token, tg_id FROM users WHERE token = $1", token).Scan(&u)
	return u, err
}

// SelectUsers returns all users.
func (db *DB) SelectUsers() ([]string, error) {
	var us []string
	rows, err := db.Query(context.Background(), "SELECT token, tg_id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u string
		err := rows.Scan(&u)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}
