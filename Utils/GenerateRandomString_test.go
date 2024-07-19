package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomString(t *testing.T) {
	res := GenerateRandomString(20)
	require.Equal(t, 20, len(res))
}
