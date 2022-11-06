package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/Rich-dev-team/stock-prediction-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createStockPriceRequest struct {
	CompanyID int64     `json:"company_id" binding:"required,min=1"`
	Price     int32     `json:"price" binding:"required,min=0"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
}

func (server *Server) createStockPrice(ctx *gin.Context) {
	var req createStockPriceRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// check the company is existed
	_, err := server.store.GetCompanyById(ctx, req.CompanyID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// inseret stock price
	arg := db.CreateStockPriceParams{
		CompanyID: req.CompanyID,
		Price:     req.Price,
		CreatedAt: req.CreatedAt,
	}
	price, err := server.store.CreateStockPrice(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, price)
}

type listAllStockPriceRequest struct {
	CompanyID int64 `form:"company_id" binding:"required,min=1"`
	PageID    int32 `form:"page_id" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) listAllStockPrice(ctx *gin.Context) {
	var req listAllStockPriceRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListAllStockPriceParams{
		CompanyID: req.CompanyID,
		Offset:    (req.PageID - 1) * req.PageSize,
		Limit:     req.PageSize,
	}
	prices, err := server.store.ListAllStockPrice(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, prices)
}

type listStockPriceByRangeRequest struct {
	CompanyID int64     `form:"company_id" binding:"required,min=1"`
	Limit     int32     `form:"limit" binding:"required,min=1"`
	Starttime time.Time `form:"starttime" binding:"required"`
	Endtime   time.Time `form:"endtime" binding:"required"`
}

func (server *Server) listStockPriceByRange(ctx *gin.Context) {
	var req listStockPriceByRangeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListStockPriceByRangeForUpdateParams{
		CompanyID: req.CompanyID,
		Limit:     req.Limit,
		Starttime: req.Starttime,
		Endtime:   req.Endtime,
	}

	prices, err := server.store.ListStockPriceByRangeForUpdate(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, prices)
}

type updateStockPriceRequest struct {
	ID    int64 `json:"id" binding:"required,min=1"`
	Price int32 `json:"price" binding:"required,min=0"`
}

func (server *Server) UpdateStockPrice(ctx *gin.Context) {
	var req updateStockPriceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateStockPriceParams{
		ID:    req.ID,
		Price: req.Price,
	}

	price, err := server.store.UpdateStockPrice(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, price)
}

type deleteStockPriceRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteStockPrice(ctx *gin.Context) {
	var req deleteStockPriceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	err := server.store.DeleteStockPrice(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true})
}
