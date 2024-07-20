package usecase

import (
	"context"

	entity "github.com/resistance22/university_project/Entity"
	repository "github.com/resistance22/university_project/Repository"
	validator "github.com/resistance22/university_project/Validator"
)

type IConsumableUseCase interface {
	Create(ctx context.Context, consumable *validator.CreateConsumableBody) (*entity.Consumable, error)
}

type ConsumableUseCase struct {
	repo repository.IConsumableRepository
}

func NewConsumableUseCase(repo repository.IConsumableRepository) IConsumableUseCase {
	return &ConsumableUseCase{
		repo: repo,
	}
}

func (usecase *ConsumableUseCase) Create(ctx context.Context, consumable *validator.CreateConsumableBody) (*entity.Consumable, error) {
	entity, err := entity.NewConsumable(
		consumable.Title,
		consumable.UOM,
		consumable.Remaining,
	)

	if err != nil {
		return nil, err
	}

	err = usecase.repo.Create(ctx, entity)

	if err != nil {
		return nil, err
	}

	return entity, nil

}
