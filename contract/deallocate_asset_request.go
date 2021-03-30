package contract

import (
	"time"

	"github.com/vkhichar/assets-manager/custom_errors"
)

type DeallocateAssetRequest struct {
	Date string `json:"date"`
}

func (alloc_request DeallocateAssetRequest) Validate() error {

	const shortForm = "2006-01-02"
	_, err := time.Parse(shortForm, alloc_request.Date)

	if err != nil {
		return custom_errors.InvalidDateFormatError
	}
	return nil
}
