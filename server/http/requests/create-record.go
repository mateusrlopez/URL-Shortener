package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateRecord struct {
	LongURL string `json:"longURL"`
}

func (req CreateRecord) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.LongURL, validation.Required, is.RequestURL),
	)
}
