package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCompany(t *testing.T) {
	args := CreateCompanyParams{
		CompanyName: "testCompany",
		StockSymbol: "12345",
	}
	company, err := testQueries.CreateCompany(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, company)

	require.Equal(t, args.CompanyName, company.CompanyName)
	require.Equal(t, args.StockSymbol, company.StockSymbol)

	require.NotZero(t, company.CreatedAt)
}
