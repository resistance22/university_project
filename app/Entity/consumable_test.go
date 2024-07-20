package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewConsumable(t *testing.T) {
	now := time.Now()
	title := "New"
	uom := "Barrels"
	remaining := 0.0
	consumable, err := NewConsumable(
		title,
		uom,
		remaining,
	)

	require.NoError(t, err)
	require.NotEmpty(t, consumable.ID)
	require.NotEmpty(t, consumable.CreatedAt)
	require.WithinDuration(t, consumable.CreatedAt, now, time.Second)
	require.Equal(t, title, consumable.Title)
	require.Equal(t, uom, consumable.UOM)
	require.Equal(t, remaining, consumable.Remaining)
}
