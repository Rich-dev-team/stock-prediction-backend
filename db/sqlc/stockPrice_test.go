package db

import (
	"context"
	"stock-prediction-backend/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomStockPrice(t *testing.T, company Company) StockPrice {

	args := CreateStockPriceParams{
		CompanyID: company.ID,
		Price:     int32(util.RandomInt(10, 100)),
		CreatedAt: time.Now(),
	}

	stockPrice, err := testQueries.CreateStockPrice(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, stockPrice)

	require.Equal(t, args.CompanyID, company.ID)
	require.Equal(t, args.Price, stockPrice.Price)

	require.NotZero(t, stockPrice.ID)
	require.NotZero(t, stockPrice.CreatedAt)

	require.WithinDuration(t, stockPrice.CreatedAt, args.CreatedAt, time.Second)
	return stockPrice
}

func TestCreateStockPrice(t *testing.T) {
	company := createRandomCompany(t)
	createRandomStockPrice(t, company)
}
func TestListPriceByRange(t *testing.T) {
	company := createRandomCompany(t)
	for i := 0; i < 10; i++ {
		createRandomStockPrice(t, company)
	}

	args := ListStockPriceByRangeParams{
		CompanyID:   company.ID,
		CreatedAt:   time.Now().Add(time.Second * -1),
		CreatedAt_2: time.Now().Add(time.Second * 1),
		Limit:       10,
	}
	prices, err := testQueries.ListStockPriceByRange(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, prices)
	require.Len(t, prices, 10)

	for _, price := range prices {
		require.NotEmpty(t, price)
	}
}

func TestUpdateStockPrice(t *testing.T) {
	company := createRandomCompany(t)
	price1 := createRandomStockPrice(t, company)

	args := UpdateStockPriceParams{
		ID:    price1.ID,
		Price: int32(util.RandomInt(100, 200)),
	}

	price2, err := testQueries.UpdateStockPrice(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, price2)

	require.Equal(t, price1.ID, price2.ID)
	require.Equal(t, price1.CompanyID, price2.CompanyID)
	require.Equal(t, price2.Price, args.Price)
	require.Equal(t, price1.CreatedAt, price2.CreatedAt)
}

func TestListAllStockPrice(t *testing.T){
	company := createRandomCompany(t)
	for i := 0; i < 10; i++ {
		createRandomStockPrice(t, company)
	}

	args := ListAllStockPriceParams{
		CompanyID: company.ID,
		Limit: 3,
		Offset: 3,
	}
	prices, err := testQueries.ListAllStockPrice(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, prices, 3)

	for _, price :=range prices{
		require.NotEmpty(t, price)
	}
}