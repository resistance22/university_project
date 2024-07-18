package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func TestCreateConsumable(t *testing.T) {
	createdAt := pgtype.Timestamp{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	uuid := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}

	arg := CreateConsumableParams{
		ID:        uuid,
		Title:     "New_Consumable",
		Uom:       "PCS",
		Remaining: 25,
		CreatedAt: createdAt,
	}

	product, err := testQueries.CreateConsumable(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, product.ID, uuid)
	require.NotEmpty(t, product)
	require.Equal(t, product.Title, arg.Title)
	require.Equal(t, product.Uom, arg.Uom)
	require.Equal(t, product.Remaining, arg.Remaining)
	require.Equal(t, product.CreatedAt.Time.Year(), arg.CreatedAt.Time.Year())
	require.Equal(t, product.CreatedAt.Time.Month(), arg.CreatedAt.Time.Month())
	require.Equal(t, product.CreatedAt.Time.Minute(), arg.CreatedAt.Time.Minute())
	require.Equal(t, product.CreatedAt.Time.Second(), arg.CreatedAt.Time.Second())
	require.NotZero(t, product.ID)

	consumables, err := testQueries.GetAllConsumable(context.Background())

	require.NoError(t, err)
	require.Len(t, consumables, 1)
}

func TestCreateUser(t *testing.T) {
	createdAt := pgtype.Timestamp{
		Time:  time.Now().UTC(),
		Valid: true,
	}

	uuid := pgtype.UUID{
		Bytes: uuid.New(),
		Valid: true,
	}

	arg := CreateUserParams{
		ID:        uuid,
		CreatedAt: pgtype.Date(createdAt),
		FirstName: pgtype.Text{String: "Amin", Valid: true},
		LastName:  pgtype.Text{String: "Foroutan", Valid: true},
		UserName:  pgtype.Text{String: "amin_foroutan", Valid: true},
		Password:  pgtype.Text{String: "COMPLEX_PASSWORD", Valid: true},
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user.ID, uuid)
	require.NotEmpty(t, user)
	require.Equal(t, user.FirstName, arg.FirstName)
	require.Equal(t, user.LastName, arg.LastName)
	require.Equal(t, user.UserName, arg.UserName)
	require.Equal(t, user.Password, arg.Password)

	users, err := testQueries.GetAllUsers(context.Background())

	require.NoError(t, err)
	require.Len(t, users, 1)

}
