package responses

import "time"

type FindOneRecord struct {
	Key        string     `json:"key"`
	URL        string     `json:"url"`
	Accesses   uint       `json:"accesses"`
	LastAccess *time.Time `json:"lastAccess"`
	CreatedAt  time.Time  `json:"createdAt"`
}
