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
	require.WithinDuration(t, product.CreatedAt.Time, arg.CreatedAt.Time, time.Second)
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
		CreatedAt: createdAt,
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

func TestGetUserByUserName(t *testing.T) {
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
		CreatedAt: createdAt,
		FirstName: pgtype.Text{String: "Amin", Valid: true},
		LastName:  pgtype.Text{String: "Foroutan", Valid: true},
		UserName:  pgtype.Text{String: "amin_f", Valid: true},
		Password:  pgtype.Text{String: "COMPLEX_PASSWORD", Valid: true},
	}

	_, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)

	user, err := testQueries.GetUserByUserName(context.Background(), arg.UserName)

	require.NoError(t, err)
	require.Equal(t, arg.ID.Bytes, user.ID.Bytes)
	require.WithinDuration(t, arg.CreatedAt.Time, arg.CreatedAt.Time, time.Millisecond)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.Password, user.Password)
}
