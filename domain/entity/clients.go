package entity

import "time"

type Clients struct {
	ID           int32     `json:"id"`
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	Channel      int32     `json:"channel"`
	Status       int32     `json:"status"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	CreatedBy    int32     `json:"created_by,omitempty"`
	UpdatedAt    NullTime  `json:"updated_at,omitempty"`
	UpdatedBy    NullInt32 `json:"updated_by,omitempty"`
}
