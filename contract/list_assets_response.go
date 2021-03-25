package contract

import "github.com/vkhichar/assets-manager/domain"

type ListAssetsResponse struct {
	Assets []domain.Asset
}
