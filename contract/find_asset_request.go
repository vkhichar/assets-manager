package contract

import (
	"errors"
)

type FindAssetRequest struct {
	Id int `json:"id"`
}

func (req *FindAssetRequest) Validate() error {

	if req.Id < 0 {
		return errors.New("invalid id ")
	}
	return nil
}
