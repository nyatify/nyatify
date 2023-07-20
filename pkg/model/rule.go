package model

import (
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
)

// Rule is a struct that defines when to send notification. It might be complex.
type Rule struct {
	SendAt time.Time `json:"send_at"`
}

// IsValid returns true if the rule is valid.
func (r *Rule) IsValid() bool {
	return !r.SendAt.Before(time.Now())
}

func (r Rule) ToJSON() string {
	b, err := json.Marshal(r)
	if err != nil {
		log.Debug().Err(err).Msg("failed to marshal rule")
		return ""
	}
	return string(b)
}

func (r *Rule) FromJSON(body string) error {
	return json.Unmarshal([]byte(body), r)
}

var (
	// intervals
	Hourly  = time.Hour
	Daily   = time.Hour * 24
	Weekly  = time.Hour * 24 * 7
	Monthly = time.Hour * 24 * 30
)

func SendAtRule(sendAt time.Time) Rule {
	return Rule{
		SendAt: sendAt,
	}
}
