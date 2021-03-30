package contract

import (
	"encoding/json"
	"errors"
	"strings"
)

type CreateAssetRequest struct {
	Name          string           `json:"name"`
	Category      string           `json:"category"`
	Specification *json.RawMessage `json:"specification"`
	InitCost      float64          `json:"initCost"`
}

func (req CreateAssetRequest) Validate() error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(req.Category) == "" {
		return errors.New("category is required")
	}

	if req.InitCost < 0 {
		return errors.New("Init cost is Invalid")
	}

	return nil
}
