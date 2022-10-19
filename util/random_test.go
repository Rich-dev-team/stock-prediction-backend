package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandInt(t *testing.T) {
	min := int64(0)
	max := int64(2000000)

	result := RandomInt(min, max)

	require.NotEmpty(t, result)
	require.NotZero(t, result)
	require.GreaterOrEqual(t, result, min)
	require.LessOrEqual(t, result, max)
}

func TestRandomString(t *testing.T) {
	length := 10
	result := RandomString(length)

	require.NotEmpty(t, result)
	require.Len(t, result, 10)
}

func TestRandomCompanyName(t *testing.T) {
	length := 10
	name1 := RandomCompanyName(length)

	require.NotEmpty(t, name1)
	require.NotZero(t, name1)
	require.Len(t, name1, length)

	name2 := RandomCompanyName(0)
	require.NotEmpty(t, name2)
	require.NotZero(t, name2)
	require.Len(t, name2, 6)

}

func TestRandomStockSymbol(t *testing.T) {
	result := RandomStockSymbol()
	require.NotEmpty(t, result)
	require.NotZero(t, result)
	require.GreaterOrEqual(t, len(result), 1)
	require.LessOrEqual(t, len(result), 999)
}
