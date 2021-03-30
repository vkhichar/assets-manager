package contract

import (
	"time"

	"github.com/vkhichar/assets-manager/custom_errors"
)

type AllocateAssetRequest struct {
	Asset_id int    `json:"asset_id"`
	Date     string `json:"date"`
}

func (alloc_request AllocateAssetRequest) Validate() error {

	if alloc_request.Asset_id < 0 {
		return custom_errors.InvalidAssetOrUserIdError
	}
	const shortForm = "2006-01-02"
	_, err := time.Parse(shortForm, alloc_request.Date)

	if err != nil {
		return custom_errors.InvalidDateFormatError
	}
	return nil
}
