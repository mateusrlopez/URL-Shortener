package models

import (
	"time"
)

type Record struct {
	Key        string `bson:"_id"`
	URL        string `bson:"url"`
	Accesses   uint
	LastAccess *time.Time `bson:"last_access"`
	CreatedAt  time.Time  `bson:"created_at"`
}
