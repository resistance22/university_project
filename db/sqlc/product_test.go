package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateConsumable(t *testing.T) {
	createdAt := pgtype.Timestamp{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	arg := CreateConsumableParams{
		Title:     "New_Consumable",
		Uom:       "PCS",
		Remaining: 25,
		CreatedAt: createdAt,
	}

	product, err := testQueries.CreateConsumable(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, product)
	require.Equal(t, product.Title, arg.Title)
	require.Equal(t, product.Uom, arg.Uom)
	require.Equal(t, product.Remaining, arg.Remaining)
	require.Equal(t, product.CreatedAt.Time.Year(), arg.CreatedAt.Time.Year())
	require.Equal(t, product.CreatedAt.Time.Month(), arg.CreatedAt.Time.Month())
	require.Equal(t, product.CreatedAt.Time.Minute(), arg.CreatedAt.Time.Minute())
	require.Equal(t, product.CreatedAt.Time.Second(), arg.CreatedAt.Time.Second())
	require.NotZero(t, product.ID)
}
