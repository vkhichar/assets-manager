package repository_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vkhichar/assets-manager/config"
	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/repository"
)

func TestAllocateAsset_When_Returns_SqlNoRowsError(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()

	assetRepo := repository.NewAssetAllocationRepo()

	err := assetRepo.AllocateAsset(context.Background(), -1, -1, "2020-12-12")

	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)

}

func TestAllocateAsset_When_Returns_InvalidStatusError(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM allocations;")
	tx.MustExec("DELETE FROM assets;")
	tx.Commit()

	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(1,'test','test',1000,3)")
	db.Query("INSERT INTO users(id,name,email,password,is_admin) VALUES(1,'test','test','test',true)")

	assetRepo := repository.NewAssetAllocationRepo()

	err := assetRepo.AllocateAsset(context.Background(), 1, 1, "2020-12-12")

	assert.Error(t, err)
	assert.Equal(t, custom_errors.InvalidAssetStatusError, err)

}

func TestAllocateAsset_When_Success(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM allocations;")
	tx.MustExec("DELETE FROM assets;")
	tx.Commit()

	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(1,'test','test',1000,0)")
	db.Query("INSERT INTO users(id,name,email,password,is_admin) VALUES(1,test,test,test,true)")

	assetRepo := repository.NewAssetAllocationRepo()

	err := assetRepo.AllocateAsset(context.Background(), 1, 1, "2020-12-12")

	assert.NoError(t, err)
	assert.Nil(t, err)

}

func TestDeallocateAsset_When_SqlNoRowsError(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()

	assetRepo := repository.NewAssetAllocationRepo()

	err := assetRepo.DeallocateAsset(context.Background(), -1, -1, "2020-12-12")

	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestDeallocateAsset_When_Success(t *testing.T) {

	configEnvVars()
	config.Init()
	repository.InitDB()
	db := repository.GetDB()

	tx := db.MustBegin()
	tx.MustExec("DELETE FROM allocations;")
	tx.MustExec("DELETE FROM assets;")
	tx.Commit()

	db.Query("INSERT INTO ASSETS(id,name,category,init_cost,status) VALUES(1,'test','test',1000,0)")
	db.Query("INSERT INTO USERS(id,name,email,password,is_admin) VALUES(1,test,test,test,true)")
	db.Query("INSERT INTO ALLOCATIONS(user_id,asset_id,date_alloc) VALUES(1,1,'2020-12-12')")

	assetRepo := repository.NewAssetAllocationRepo()

	err := assetRepo.DeallocateAsset(context.Background(), 1, 1, "2020-12-20")

	assert.NoError(t, err)
	assert.Nil(t, err)
}
