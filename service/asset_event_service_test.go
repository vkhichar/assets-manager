package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/config"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/service"
	"gopkg.in/h2non/gock.v1"
)

func Test_PostCreateAssetEvent_Success(t *testing.T) {
	defer gock.Off()

	reqAsset := domain.Asset{
		Id:            1,
		Name:          "Ideapad",
		Category:      "Laptops",
		Specification: nil,
		InitCost:      45000,
		Status:        0,
	}

	gock.
		New(config.GetAssetServiceURL()).
		Post("/events").
		Reply(200).
		JSON(map[string]int{"id": 1})

	assetEvServ := service.NewAssetEventService()

	respAsset, err := assetEvServ.PostCreateAssetEvent(context.Background(), &reqAsset)

	assert.Nil(t, err)
	assert.JSONEq(t, "1", respAsset)

}

func Test_PostCreateAssetEvent_When_Fails(t *testing.T) {
	defer gock.Off()

	reqAsset := domain.Asset{
		Id:            1,
		Name:          "Ideapad",
		Category:      "Laptops",
		Specification: nil,
		InitCost:      45000,
		Status:        0,
	}

	gock.
		New(config.GetAssetServiceURL()).
		Post("/events").
		Reply(400)

	assetEvServ := service.NewAssetEventService()

	_, err := assetEvServ.PostCreateAssetEvent(context.Background(), &reqAsset)

	assert.NotNil(t, err)

}

func Test_PostCreateAssetEvent_When_TimeOutError(t *testing.T) {
	defer gock.Off()

	reqAsset := domain.Asset{
		Id:            1,
		Name:          "Ideapad",
		Category:      "Laptops",
		Specification: nil,
		InitCost:      45000,
		Status:        0,
	}
	timeoutErr := errors.New("timeout error")

	gock.
		New(config.GetAssetServiceURL()).
		Post("/events").
		ReplyError(timeoutErr)

	assetEvServ := service.NewAssetEventService()

	_, err := assetEvServ.PostCreateAssetEvent(context.Background(), &reqAsset)
	assert.NotNil(t, err)

}

func Test_PostUpdateAssetEvent_Success(t *testing.T) {
	defer gock.Off()

	reqAsset := domain.Asset{
		Id:            1,
		Name:          "Ideapad",
		Category:      "Laptops",
		Specification: nil,
		InitCost:      45000,
		Status:        0,
	}

	gock.
		New(config.GetAssetServiceURL()).
		Post("/events").
		Reply(200).
		JSON(map[string]int{"id": 1})

	assetEvServ := service.NewAssetEventService()

	respAsset, err := assetEvServ.PostUpdateAssetEvent(context.Background(), &reqAsset)

	assert.Nil(t, err)
	assert.JSONEq(t, "1", respAsset)

}

func Test_PostUpdateAssetEvent_When_Fails(t *testing.T) {
	defer gock.Off()

	reqAsset := domain.Asset{
		Id:            1,
		Name:          "Ideapad",
		Category:      "Laptops",
		Specification: nil,
		InitCost:      45000,
		Status:        0,
	}

	gock.
		New(config.GetAssetServiceURL()).
		Post("/events").
		Reply(400)

	assetEvServ := service.NewAssetEventService()

	_, err := assetEvServ.PostUpdateAssetEvent(context.Background(), &reqAsset)

	assert.NotNil(t, err)

}

func Test_PostUpdateAssetEvent_When_TimeOutError(t *testing.T) {
	defer gock.Off()

	reqAsset := domain.Asset{
		Id:            1,
		Name:          "Ideapad",
		Category:      "Laptops",
		Specification: nil,
		InitCost:      45000,
		Status:        0,
	}
	timeoutErr := errors.New("timeout error")

	gock.
		New(config.GetAssetServiceURL()).
		Post("/events").
		ReplyError(timeoutErr)

	assetEvServ := service.NewAssetEventService()

	_, err := assetEvServ.PostUpdateAssetEvent(context.Background(), &reqAsset)
	assert.NotNil(t, err)

}
