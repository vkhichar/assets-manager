package contract

import (
	"errors"
)

type FindAssetRequest struct {
	Id int `json:"id"`
}

func (req *FindAssetRequest) Validate() error {

	if req.Id < 0 || req.Id > 99999 {
		return errors.New("invalid id ")
	}
	return nil
}
