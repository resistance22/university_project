package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashString(t *testing.T) {
	stringToHash := "213123123123"
	res, err := HashString(stringToHash)
	require.NoError(t, err)
	require.NotEqual(t, res, stringToHash)
}

type Test struct {
	Hello    string `json:"hello"`
	SomeName string `json:"some_name"`
	NoJSON   string
}

func TestConvertToJsonForm(t *testing.T) {
	res, err := StructToMapWithJSONKeys(&Test{
		Hello:    "hello",
		SomeName: "some_name",
		NoJSON:   "no_json",
	})

	require.NoError(t, err)
	require.Contains(t, res, "hello")
	require.Contains(t, res, "some_name")
	require.Contains(t, res, "NoJSON")
	require.Equal(t, res["some_name"], "some_name")
	require.Equal(t, res["hello"], "hello")
	require.Equal(t, res["NoJSON"], "no_json")
}
