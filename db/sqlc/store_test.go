package db

import (
	"context"
	"stock-prediction-backend/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomBatchStockParam(t *testing.T, amount int) CreateBatchStockPriceParams {

	company := createRandomCompany(t)
	companys_id := make([]int64, amount)
	prices := make([]int32, amount)
	createdAts := make([]time.Time, amount)

	// create companys_id, prices, createdAts
	for i := range companys_id {
		companys_id[i] = company.ID
		prices[i] = int32(util.RandomInt(10, 500))
		createdAts[i] = time.Now().Add(time.Second * time.Duration(util.RandomInt(1, 10)))
	}
	return CreateBatchStockPriceParams{
		CompanysID: companys_id,
		Prices:     prices,
		CreatedAts: createdAts,
	}
}

func TestCreateBatchStockPrice(t *testing.T) {
	amount := 20
	store := NewStore(testDB)

	n := 10
	// run n concurrent store the stock
	errs := make(chan error)
	results := make(chan []StockPrice)
	companys_id := make(chan int64)

	for i := 0; i < n; i++ {

		go func() {
			arg := createRandomBatchStockParam(t, amount)
			result, err := store.CreateStocksTx(context.Background(), arg)

			errs <- err
			results <- result
			// the company id are same in each epoch
			companys_id <- arg.CompanysID[0]
		}()

	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		company_id := <-companys_id
		require.NotZero(t, company_id)

		// validate the amount is correct
		args := ListAllStockPriceParams{
			CompanyID: company_id,
			Limit:     int32(amount * 2),
			Offset:    0,
		}
		prices, err := testQueries.ListAllStockPrice(context.Background(), args)
		require.NoError(t, err)
		require.Len(t, prices, amount)
	}

}
