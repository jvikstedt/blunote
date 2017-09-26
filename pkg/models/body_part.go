package models

import "time"

type BodyPart struct {
	ID        int
	Meta      string
	Raw       string
	CreatedAt time.Time
	UpdatedAt time.Time
}
