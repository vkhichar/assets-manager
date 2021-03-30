package contract

import (
	"encoding/json"
	"errors"
	"strings"
)

type UpadateAssetRequest struct {
	Id            int              `json:"id"`
	Name          string           `json:"name"`
	Category      string           `json:"category"`
	Specification *json.RawMessage `json:"specification"`
	InitCost      float64          `json:"init_cost"`
	Status        int              `json:"status"`
}

func (req UpadateAssetRequest) Validate() error {
	if req.Id < 0 {
		return errors.New("invalid id")
	}
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(req.Category) == "" {
		return errors.New("category is required")
	}

	if req.InitCost < 0 {
		return errors.New("Init cost is Invalid")
	}
	if req.Status < 0 || req.Status > 5 {
		return errors.New("status is invalid")
	}
	return nil
}
