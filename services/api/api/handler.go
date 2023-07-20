package api

import (
	"io"
	"net/http"

	"github.com/nyatify/nyatify/pkg/model"
	"github.com/rs/zerolog/log"
)

type Server struct {
	*Service
}

func NewServer(s *Service) *Server {
	return &Server{s}
}

func (s *Server) Run() {
	log.Debug().Msg("Start server on :8080")
	http.HandleFunc("/createNotification", s.createNotification)

	http.ListenAndServe(":8080", nil) // TODO: use env
}

func onError(w http.ResponseWriter, err string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(err))
}

func (s *Server) createNotification(w http.ResponseWriter, r *http.Request) {
	j, err := io.ReadAll(r.Body)
	if err != nil {
		log.Err(err).Msg("CreateNotification: io.ReadAll error")
		onError(w, err.Error(), http.StatusBadRequest)
		return
	}

	var n model.Notification

	err = n.FromJSON(j)
	if err != nil {
		log.Err(err).Msg("CreateNotification: json.Unmarshal error")
		onError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !n.Rule.IsValid() {
		log.Error().Msg("CreateNotification: invalid rule")
		onError(w, "Rule is invalid", http.StatusBadRequest)
		return
	}

	n.By = generateToken()

	err = s.db.InsertNotification(n)
	if err != nil {
		log.Err(err).Msg("CreateNotification: s.db.InsertNotification error")
		onError(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Info().Msg("CreateNotification: success " + n.Client)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(n.By))
}
