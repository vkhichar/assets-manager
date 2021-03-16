package repository

import (
	"context"
	"errors"

	"github.com/vkhichar/assets-manager/contract"
	"github.com/vkhichar/assets-manager/domain"

	"github.com/jmoiron/sqlx"
)

const (
	FindAssetByIdQuery = "SELECT * FROM assets WHERE id = $1"
	GetAllAssetsQuery  = "SELECT * FROM assets"
	UpdateAssetsQuery  = "UPDATE assets SET name=$2, category=$3, specification=$4,init_cost=$5,status=$6 WHERE id=$1"
	DeleteAssetQuery   = "DELETE FROM assets WHERE ID=$1"
)

type AssetRepository interface {
	FindAsset(context context.Context, id int) (*domain.Asset, error)
	GetAllAssets() (*[]domain.Asset, error)
	UpdateAsset(ctx context.Context, asset *contract.UpadateAssetRequest) (*domain.Asset, error)
	DeleteAsset(ctx context.Context, id int) (*domain.Asset, error)
}

type assetRepo struct {
	db *sqlx.DB
}

func NewAssetRepository() AssetRepository {
	return &assetRepo{
		db: GetDB(),
	}
}
func (repo *assetRepo) FindAsset(context context.Context, id int) (*domain.Asset, error) {

	var asset domain.Asset
	err := repo.db.Get(&asset, FindAssetByIdQuery, id)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (repo *assetRepo) GetAllAssets() (*[]domain.Asset, error) {

	var assets []domain.Asset
	err := repo.db.Select(&assets, GetAllAssetsQuery)
	if err != nil {
		return nil, err
	}
	return &assets, nil

}

func (repo *assetRepo) UpdateAsset(ctx context.Context, asset *contract.UpadateAssetRequest) (*domain.Asset, error) {

	result, err := repo.db.Exec(UpdateAssetsQuery, asset.Id, asset.Name, asset.Category, asset.Specification, asset.InitCost, asset.Status)
	rows_affected, _ := result.RowsAffected()

	if rows_affected < 1 {
		return nil, errors.New("something went wrong")
	}
	if err != nil {
		return nil, err
	}
	updatedAsset, err := repo.FindAsset(ctx, asset.Id)
	if err != nil {
		return nil, err
	}
	return updatedAsset, nil

}

func (repo *assetRepo) DeleteAsset(ctx context.Context, id int) (*domain.Asset, error) {

	asset, err := repo.FindAsset(ctx, id)
	if err != nil {

		return nil, errors.New("something went wrong")
	}

	result, err := repo.db.Exec(DeleteAssetQuery, id)
	rows_affected, _ := result.RowsAffected()

	if rows_affected < 1 {
		return nil, errors.New("something went wrong")
	}
	if err != nil {
		return nil, err
	}

	return asset, nil
}
