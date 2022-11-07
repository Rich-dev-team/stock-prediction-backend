package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/Rich-dev-team/stock-prediction-backend/db/mock"
	db "github.com/Rich-dev-team/stock-prediction-backend/db/sqlc"
	"github.com/Rich-dev-team/stock-prediction-backend/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetCompanyAPI(t *testing.T) {
	company := randomCompany()

	testCases := []struct {
		name          string
		CompanyName   interface{}
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:        "OK",
			CompanyName: company.CompanyName,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCompanyByName(gomock.Any(), gomock.Eq(company.CompanyName)).
					Times(1).
					Return(company, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchCompany(t, recorder.Body, company)
			},
		},
		{
			name:        "Not Found",
			CompanyName: company.CompanyName,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCompanyByName(gomock.Any(), gomock.Eq(company.CompanyName)).
					Times(1).
					Return(db.Company{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:        "InternalError",
			CompanyName: company.CompanyName,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCompanyByName(gomock.Any(), gomock.Eq(company.CompanyName)).
					Times(1).
					Return(db.Company{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:        "Invalid company Name",
			CompanyName: "1234567890123456789012345678900",
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCompanyByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			// build stubs
			tc.buildStubs(store)

			// start test server and send request
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/companys/list/%s", tc.CompanyName)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			// check response
			tc.checkResponse(t, recorder)
		})
	}
}
func randomCompany() db.Company {
	return db.Company{
		ID:          util.RandomInt(1, 10000),
		CompanyName: util.RandomCompanyName(10),
		StockSymbol: util.RandomStockSymbol(),
		CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
	}
}

func requireBodyMatchCompany(t *testing.T, body *bytes.Buffer, company db.Company) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotCompany db.Company

	err = json.Unmarshal(data, &gotCompany)
	require.NoError(t, err)
	require.Equal(t, company.ID, gotCompany.ID)
	require.Equal(t, company.CompanyName, gotCompany.CompanyName)
	require.Equal(t, company.StockSymbol, gotCompany.StockSymbol)
	// If using the origin time is weird, so using the unix time
	require.Equal(t, company.CreatedAt.Time.Unix(), gotCompany.CreatedAt.Time.Unix())
}
