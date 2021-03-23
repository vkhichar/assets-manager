package domain

import "encoding/json"

type Asset struct {
	ID            int              `db:"id"`
	Name          string           `db:"name"`
	Category      string           `db:"category"`
	Specification *json.RawMessage `db:"specification"`
	InitCost      float64          `db:"init_cost"`
	Status        int              `db:"status"`
}
