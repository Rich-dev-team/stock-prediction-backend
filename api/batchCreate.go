package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	db "github.com/Rich-dev-team/stock-prediction-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)


type CreateBatchStockPricesRequest struct {
	CompanysID []int64     `json:"companys_id" binding:"required,companysId"`
	Prices     []int32     `json:"prices" binding:"required,stockPrices"`
	CreatedAts []time.Time `json:"created_ats" binding:"required"`
}

func (server *Server) createBatchStockPrices(ctx *gin.Context) {

	var req CreateBatchStockPricesRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// validate the data has same length
	if len(req.CompanysID) != len(req.Prices) && len(req.CompanysID) != len(req.CreatedAts){
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("data should be same length")))
		return
	}
	// validate account is exists or not
	isErrs := make(chan bool)
	for i := range req.CompanysID {
		id := req.CompanysID[i]
		go func(id int64) {
			isExists := server.validateAccount(ctx, id)
			isErrs <- isExists
		}(id)
	}
	
	// check the channel
	for i :=0; i< len(req.CompanysID); i++ {
		if ! <-isErrs{
			return
		}
	}

	arg := db.CreateBatchStockPriceParams{
		CompanysID: req.CompanysID,
		Prices:     req.Prices,
		CreatedAts: req.CreatedAts,
	}
	result, err := server.store.CreateStocksTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validateAccount(ctx *gin.Context, companyID int64) bool {
	_, err := server.store.GetCompanyById(ctx, companyID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	return true
}
