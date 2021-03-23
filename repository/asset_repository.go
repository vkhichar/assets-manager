package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/vkhichar/assets-manager/domain"
)

const (
	createNewAssetQuery = `INSERT INTO assets(name,category,specification,init_cost,status) values($1,$2,$3,$4,$5) RETURNING id`
	getAssetByIdQuery   = `SELECT * FROM assets WHERE id=$1`
)

//
type AssetRepository interface {
	CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error)
	FindAsset(ctx context.Context, id int) (*domain.Asset, error)
}

type assetRepo struct {
	db *sqlx.DB
}

//
func NewAssetRepository() AssetRepository {
	repo := &assetRepo{
		db: GetDB(),
	}
	return repo
}

func (repo *assetRepo) CreateAsset(ctx context.Context, asset *domain.Asset) (*domain.Asset, error) {
	id := 0
	err := repo.db.QueryRow(createNewAssetQuery, asset.Name, asset.Category, asset.Specification, asset.InitCost, 0).Scan(&id)
	if err != nil {
		return nil, err
	}
	var returnAsset domain.Asset
	err = repo.db.Get(&returnAsset, getAssetByIdQuery, id)

	return &returnAsset, nil
}

func (repo *assetRepo) FindAsset(ctx context.Context, id int) (*domain.Asset, error) {
	var asset domain.Asset
	err := repo.db.Get(&asset, getAssetByIdQuery, id)
	if err == sql.ErrNoRows {
		fmt.Printf("repository: couldn't find asset for id: %v", id)
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &asset, nil
}
