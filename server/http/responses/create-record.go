package responses

type CreateRecord struct {
	ID int64 `json:"id"`
	ShortURLKey string `json:"shortURLKey"`
}
