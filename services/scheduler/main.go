package main

import (
	"os"
	"time"

	"github.com/nyatify/nyatify/pkg/rabbit"
	"github.com/nyatify/nyatify/pkg/storage"

	"github.com/rs/zerolog/log"
)

func scanDB(db *storage.DB, s *rabbit.Client) error {
	for {
		ns, err := db.SelectNotifications()
		if err != nil {
			return err
		}

		for _, n := range ns {
			t := n.Rule.SendAt
			if t.Before(time.Now().UTC()) || t.Equal(time.Now().UTC()) {
				err = s.Schedule(n)
				if err != nil {
					return err
				}

				err = db.DeleteNotification(n.ID)
				if err != nil {
					log.Debug().Err(err).Msg("failed to delete notification")
				}
			}
		}
	}
}

func main() {
	log.Debug().Msg("Start scheduler service")
	s, err := rabbit.New(os.Getenv("RABBITMQ_HOST"))
	if err != nil {
		panic(err)
	}
	defer s.Close()

	db, err := storage.New()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = scanDB(db, s)
	if err != nil {
		panic(err)
	}
}
