package repository_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/config"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/repository"
)

func TestAssetRepository_CreateAsset_When_Success(t *testing.T) {
	ctx := context.Background()
	var assetExpected domain.Asset
	spec := json.RawMessage([]byte(`{"ram":"4GB","brand":"acer"}`))

	dummy := &domain.Asset{
		Name:          "Acer-255",
		Category:      "Laptops",
		InitCost:      50000,
		Specification: &spec,
	}

	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("TRUNCATE TABLE assets RESTART IDENTITY;")
	tx.Commit()

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.CreateAsset(ctx, dummy)

	fmt.Println()
	db.Get(&assetExpected, "SELECT * FROM assets WHERE id = $1", 1)
	fmt.Println(assetExpected)

	assert.Equal(t, &assetExpected, asset)
	assert.Nil(t, err)

	fmt.Println()
}

func TestAssetRepository_GetAsset_When_Success(t *testing.T) {
	ctx := context.Background()
	var assetExpected domain.Asset
	id := 1

	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.FindAsset(ctx, id)

	fmt.Println()
	db.Get(&assetExpected, "SELECT * FROM assets WHERE id = $1", id)
	fmt.Println(assetExpected)

	assert.NotNil(t, asset)
	assert.Equal(t, &assetExpected, asset)
	assert.Nil(t, err)
	fmt.Println()
}

func TestAssetRepository_GetAsset_When_NoAssetFound(t *testing.T) {
	ctx := context.Background()
	config.Init()
	repository.InitDB()
	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.FindAsset(ctx, 5)

	fmt.Println()
	assert.Nil(t, asset)
	assert.Nil(t, err)
	fmt.Println()
}
