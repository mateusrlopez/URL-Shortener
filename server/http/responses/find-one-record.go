package responses

import "time"

type FindOneRecord struct {
	ID          int64      `json:"id"`
	ShortURLKey string     `json:"shortURLKey"`
	LongURL     string     `json:"longURL"`
	Accesses    uint       `json:"accesses"`
	LastAccess  *time.Time `json:"lastAccess"`
	CreatedAt   time.Time  `json:"createdAt"`
}
