package contract

import "github.com/vkhichar/assets-manager/domain"

type CreateAssetEventReq struct {
	EventType string        `json:"type"`
	Data      *domain.Asset `json:"data"`
}

type UpdateAssetEventReq struct {
	EventType string        `json:"type"`
	Data      *domain.Asset `json:"data"`
}
