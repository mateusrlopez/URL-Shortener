package requests

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type CreateRecord struct {
	URL string `json:"url"`
}

func (req CreateRecord) Validate() error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.URL, validation.Required, is.RequestURL),
	)
}
