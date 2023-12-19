package entity

import "time"


type {{.ENTITY_UPPER}} struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}