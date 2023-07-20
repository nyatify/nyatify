package api

import (
	"math/rand"

	"github.com/nyatify/nyatify/pkg/model"
	"github.com/nyatify/nyatify/pkg/storage"
)

type Service struct {
	db *storage.DB
}

func NewService(db *storage.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) StoreNotification(n model.Notification) error {
	return s.db.InsertNotification(n)
}

func (s *Service) NotificationsByToken(token string) ([]model.Notification, error) {
	return s.db.SelectNotificationsByToken(token)
}

func generateToken() string {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *Service) Users() ([]string, error) {
	return s.db.SelectUsers()
}

func (s *Service) User(token string) (string, error) {
	return s.db.SelectUserByToken(token)
}

func (s *Service) Notifications() ([]model.Notification, error) {
	return s.db.SelectNotifications()
}

func (s *Service) DeleteNotification(id int) error {
	return s.db.DeleteNotification(id)
}
