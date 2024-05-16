package entity

import "time"


type {{.ENTITY_UPPER}} struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
