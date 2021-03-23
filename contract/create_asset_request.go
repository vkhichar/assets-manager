package contract

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type CreateAssetRequest struct {
	Name          string          `json:"name"`
	Category      string          `json:"category"`
	Specification json.RawMessage `json:"specification"`
	InitCost      float64         `json:"initCost"`
}

func (req CreateAssetRequest) Validate() error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if strings.TrimSpace(req.Category) == "" {
		return errors.New("category is required")
	}

	if strings.TrimSpace(fmt.Sprintf("%f", req.InitCost)) == "" {
		return errors.New("Init cost is required")
	}
	return nil
}
