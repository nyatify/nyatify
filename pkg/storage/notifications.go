package storage

import (
	"context"

	"github.com/nyatify/nyatify/pkg/model"
	"github.com/rs/zerolog/log"
)

// InsertNotification inserts a notification into the database.
func (db *DB) InsertNotification(n model.Notification) error {
	_, err := db.Exec(context.Background(), "INSERT INTO notifications (title, body, token, rules, client) VALUES ($1, $2, $3, $4, $5)", n.Title, n.Body, n.By, n.Rule.ToJSON(), n.Client)
	if err != nil {
		log.Debug().Err(err).Msg("failed to insert notification")
		return err
	}
	return nil
}

// SelectNotificationsByToken returns all notifications of a user.
func (db *DB) SelectNotificationsByToken(token string) ([]model.Notification, error) {
	var ns []model.Notification
	rows, err := db.Query(context.Background(), "SELECT id, client, title, body, rules, created_at FROM notifications WHERE token=$1", token)
	if err != nil {
		log.Debug().Err(err).Msg("failed to select notifications")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var n model.Notification
		var json string
		err := rows.Scan(&n.ID, &n.Client, &n.Title, &n.Body, &json, &n.CreatedAt)
		if err != nil {
			return nil, err
		}
		err = n.Rule.FromJSON(json)
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func (db *DB) SelectNotifications() ([]model.Notification, error) {
	var ns []model.Notification
	rows, err := db.Query(context.Background(), "SELECT id, client, title, body, rules, created_at, token FROM notifications")
	if err != nil {
		log.Debug().Err(err).Msg("failed to select notifications")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var n model.Notification
		var json string
		err := rows.Scan(&n.ID, &n.Client, &n.Title, &n.Body, &json, &n.CreatedAt, &n.By)
		if err != nil {
			return nil, err
		}
		err = n.Rule.FromJSON(json)
		if err != nil {
			return nil, err
		}

		ns = append(ns, n)
	}
	return ns, nil
}

func (db *DB) DeleteNotification(id int) error {
	_, err := db.Exec(context.Background(), "DELETE FROM notifications WHERE id=$1", id)
	if err != nil {
		log.Debug().Err(err).Msg("failed to delete notification")
		return err
	}
	return nil
}
