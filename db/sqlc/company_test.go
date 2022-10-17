package db

import (
	"context"
	"database/sql"
	"stock-prediction-backend/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCompany(t *testing.T) Company {
	args := CreateCompanyParams{
		CompanyName: util.RandomCompanyName(),
		StockSymbol: util.RandomStockSymbol(),
	}
	company, err := testQueries.CreateCompany(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, company)

	require.Equal(t, args.CompanyName, company.CompanyName)
	require.Equal(t, args.StockSymbol, company.StockSymbol)

	require.NotZero(t, company.ID)
	require.NotZero(t, company.CreatedAt)
	return company
}

func TestCreateCompany(t *testing.T) {
	createRandomCompany(t)
}

func TestGetCompany(t *testing.T) {
	// create company
	company1 := createRandomCompany(t)
	company2, err := testQueries.GetCompany(context.Background(), company1.CompanyName)
	require.NoError(t, err)
	require.NotEmpty(t, company2)

	require.Equal(t, company1.ID, company2.ID)
	require.Equal(t, company1.CompanyName, company2.CompanyName)
	require.Equal(t, company1.StockSymbol, company2.StockSymbol)
	require.Equal(t, company1.CreatedAt, company2.CreatedAt)

}

func TestUpdateCompany(t *testing.T) {
	company1 := createRandomCompany(t)

	args := UpdateCompanyNameParams{
		ID:          company1.ID,
		CompanyName: util.RandomCompanyName(),
	}

	company2, err := testQueries.UpdateCompanyName(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, company2)

	require.Equal(t, company1.ID, company2.ID)
	require.Equal(t, args.CompanyName, company2.CompanyName)
	require.Equal(t, company1.StockSymbol, company2.StockSymbol)
	require.Equal(t, company1.CreatedAt, company2.CreatedAt)
}

func TestDeleteCompany(t *testing.T) {
	company1 := createRandomCompany(t)

	err := testQueries.DeleteCompany(context.Background(), company1.ID)
	require.NoError(t, err)

	company2, err := testQueries.GetCompany(context.Background(), company1.CompanyName)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, company2)
}

func TestListCompany(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCompany(t)
	}

	args := ListCompanyParams{
		Limit:  5,
		Offset: 5,
	}
	companys, err := testQueries.ListCompany(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, companys, 5)

	for _, company := range companys {
		require.NotEmpty(t, company)
	}
}
