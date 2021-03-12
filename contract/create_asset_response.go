package contract

import "encoding/json"

type CreateAssetResponse struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Category      string          `json:"category"`
	Specification json.RawMessage `json:"specification"`
	InitCost      float64         `json:"initCost"`
	Status        int             `json:"status"`
}
