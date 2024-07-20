package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	entity "github.com/resistance22/university_project/Entity"
	db "github.com/resistance22/university_project/db/sqlc"
)

type IConsumableRepository interface {
	Create(c context.Context, u *entity.Consumable) error
}

type ConsumableRepository struct {
	store *db.Store
}

func createConsumableParams(e *entity.Consumable) *db.CreateConsumableParams {
	return &db.CreateConsumableParams{
		ID: pgtype.UUID{
			Bytes: e.ID,
			Valid: true,
		},
		CreatedAt: pgtype.Timestamp{
			Time:  e.CreatedAt,
			Valid: true,
		},
		Title:     e.Title,
		Uom:       e.UOM,
		Remaining: e.Remaining,
	}
}

func (repo *ConsumableRepository) Create(ctx context.Context, e *entity.Consumable) error {
	params := createConsumableParams(e)
	_, err := repo.store.Queries.CreateConsumable(ctx, *params)
	return err
}

func NewConsumableRepository(store *db.Store) IConsumableRepository {
	return &ConsumableRepository{
		store: store,
	}
}
