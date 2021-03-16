package contract

import (
	"encoding/json"
	"errors"
	"fmt"
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
		return errors.New("id is invalid")
	}
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(req.Category) == "" {
		return errors.New("category is required")
	}

	if strings.TrimSpace(fmt.Sprintf("%f", req.InitCost)) == "" {
		return errors.New("Init cost is required")
	}
	if req.Status < 0 || req.Status > 5 {
		return errors.New("status is invalid")
	}
	return nil
}
