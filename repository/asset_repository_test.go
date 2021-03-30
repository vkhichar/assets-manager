package repository_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/config"
	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/domain"
	"github.com/vkhichar/assets-manager/repository"
)

func configEnvVars() {
	os.Setenv("APP_PORT", "9000")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "12345")
	os.Setenv("DB_NAME", "asset_management")
}
func TestDbConnection(t *testing.T) {

	configEnvVars()
	err := config.Init()
	repository.InitDB()

	assert.NoError(t, err)
}
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
	configEnvVars()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM allocations")
	tx.MustExec("DELETE FROM assets;")
	tx.Commit()

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.CreateAsset(ctx, dummy)

	fmt.Println(err)
	db.Get(&assetExpected, "SELECT * FROM assets WHERE id = $1", asset.Id)
	fmt.Println(assetExpected)

	assert.Equal(t, &assetExpected, asset)
	assert.Nil(t, err)

	fmt.Println()
}
func TestFindAsset_When_ReturnsError(t *testing.T) {

	ctx := context.Background()
	configEnvVars()
	config.Init()
	repository.InitDB()

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.FindAsset(ctx, -1)

	assert.Error(t, err)
	assert.Nil(t, asset)
	assert.Equal(t, custom_errors.InvalidIdError.Error(), err.Error())

}

func TestFindAsset_When_ReturnsAsset(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM assets;")
	tx.Commit()

	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(1,'Test','Test',1000,0)")

	assetRepo := repository.NewAssetRepository()
	asset, err := assetRepo.FindAsset(context.Background(), 1)

	expectedAsset := domain.Asset{
		Id:       1,
		Name:     "Test",
		Category: "Test",
		InitCost: 1000,
		Status:   0,
	}
	assert.NoError(t, err)
	assert.NotNil(t, asset)
	assert.Equal(t, &expectedAsset, asset)
}

func TestUpdateAssets_When_ReturnsError(t *testing.T) {
	configEnvVars()
	config.Init()
	repository.InitDB()
	ctx := context.Background()
	assetRepo := repository.NewAssetRepository()

	asset_request := &contract.UpadateAssetRequest{
		Id:       -1,
		Name:     "hp",
		Category: "laptop",
		InitCost: 1000,
		Status:   0,
	}
	asset, err := assetRepo.UpdateAsset(ctx, asset_request)

	assert.Error(t, err)
	assert.Nil(t, asset)

}

func TestUpdateAssets_When_ReturnsAsset(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	ctx := context.Background()

	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(1,'TEST','TEST',1000,0)")

	asset_request := &contract.UpadateAssetRequest{
		Id:       1,
		Name:     "hp",
		Category: "laptop",
		InitCost: 1000,
		Status:   0,
	}

	assetRepo := repository.NewAssetRepository()
	asset_response, err := assetRepo.UpdateAsset(ctx, asset_request)

	assert.NoError(t, err)
	assert.NotNil(t, asset_response)

}

func TestDeleteAssets_When_ReturnsError(t *testing.T) {
	configEnvVars()
	config.Init()
	repository.InitDB()
	ctx := context.Background()
	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.DeleteAsset(ctx, -1)

	assert.Error(t, err)
	assert.Nil(t, asset)
	assert.Equal(t, custom_errors.InvalidIdError.Error(), err.Error())

}

func TestDeleteAssets_When_ReturnsAsset(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	ctx := context.Background()

	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(1,'TEST','TEST',1000,0)")

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.DeleteAsset(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, asset)
}
func TestGetAllAssets_When_ReturnsListOfAssets(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()
	assetRepo := repository.NewAssetRepository()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM assets;")
	tx.Commit()

	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(0,'TEST0','TEST0',1000,0)")
	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(1,'TEST1','TEST1',1000,0)")
	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(2,'TEST2','TEST2',1000,0)")

	var expectedList = make([]domain.Asset, 3)
	for i := 0; i < 3; i++ {
		expectedList[i] = domain.Asset{Id: i, Name: fmt.Sprintf("TEST%d", i), Category: fmt.Sprintf("TEST%d", i), InitCost: 1000, Status: 0}
	}

	returnedList, err := assetRepo.GetAllAssets()

	assert.NoError(t, err)
	assert.IsType(t, expectedList, returnedList)
	assert.Equal(t, expectedList, returnedList)
}

/*
func TestAssetRepository_GetAsset_When_Success(t *testing.T) {

	config.Init()
	repository.InitDB()
	db := repository.GetDB()
	ctx := context.Background()

	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(1,'TEST','TEST',1000,0)")
	var assetExpected domain.Asset

	assetRepo := repository.NewAssetRepository()

	asset, err := assetRepo.FindAsset(ctx, 1)

	fmt.Println()
	db.Get(&assetExpected, "SELECT * FROM assets WHERE id = $1", 1)
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
	assert.Error(t, err)
	assert.Nil(t, asset)
	fmt.Println()
}
*/
