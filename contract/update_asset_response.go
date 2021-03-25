package contract

import (
	"encoding/json"
)

type UpadateAssetResponse struct {
	Id            int             `json:"id"`
	Name          string          `json:"name"`
	Category      string          `json:"category"`
	Specification json.RawMessage `json:"specification"`
	InitCost      float64         `json:"init_cost"`
	Status        int             `json:"status"`
}
