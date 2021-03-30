package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/assets-manager/custom_errors"
	"github.com/vkhichar/assets-manager/domain"
)

const (
	getAssetStatusQuery    = "SELECT STATUS FROM ASSETS WHERE id = $1"
	updateAssetStatusQuery = "UPDATE ASSETS SET status = $1 WHERE id = $2"
	allocateAssetQuery     = "INSERT INTO ALLOCATIONS(user_id,asset_id,date_alloc) VALUES($1,$2,$3)"
	findAllocationQuery    = "SELECT user_id,asset_id,date_alloc,date_dealloc FROM ALLOCATIONS WHERE user_id = $1 AND asset_id = $2 AND date_dealloc IS NULL"
	deallocateAssetQuery   = "UPDATE ALLOCATIONS SET date_dealloc = $1 WHERE user_id = $2 AND asset_id = $3 "
)

type AssetsAllocationRepository interface {
	AllocateAsset(context.Context, int, int, string) error
	DeallocateAsset(context.Context, int, int, string) error
}

type assetAllocationRepo struct {
	db *sqlx.DB
}

func NewAssetAllocationRepo() AssetsAllocationRepository {
	return &assetAllocationRepo{
		db: GetDB(),
	}
}

func (asset_repo *assetAllocationRepo) AllocateAsset(context context.Context, user_id int, asset_id int, date string) error {

	var status int

	//check if asset can be allocated
	err := asset_repo.db.Get(&status, getAssetStatusQuery, asset_id)
	if err != nil {
		return err
	}
	if status != 0 {
		return custom_errors.InvalidAssetStatusError
	}

	//allocate asset
	_, err = asset_repo.db.Exec(allocateAssetQuery, user_id, asset_id, date)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// update asset table
	_, err = asset_repo.db.Exec(updateAssetStatusQuery, 1, asset_id)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("asset_id ", asset_id, "is allocated to user_id", user_id, "but status not updated in asset table")
		return errors.New("something went wrong")
	}

	return nil
}

func (asset_repo *assetAllocationRepo) DeallocateAsset(context context.Context, user_id int, asset_id int, dealloc_date string) error {

	//find allocation
	var allocation domain.Allocation
	err := asset_repo.db.Get(&allocation, findAllocationQuery, user_id, asset_id)
	if err != nil {
		fmt.Println("in get", err.Error())

		return err
	}

	// update dealloc date
	_, err = asset_repo.db.Exec(deallocateAssetQuery, dealloc_date, user_id, asset_id)

	if err != nil {

		fmt.Println("ater de allocation req", err.Error())
		return errors.New("Something went wrong")

	}

	//update asset status to available
	_, err = asset_repo.db.Exec(updateAssetStatusQuery, 0, asset_id)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("asset_id ", asset_id, "is deallocated to user_id", user_id, "but status not updated in asset table")

	}
	return nil
}
