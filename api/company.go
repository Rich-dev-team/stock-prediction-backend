package api

import (
	"database/sql"
	"net/http"
	db "stock-prediction-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createCompanyRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	StockSymbol string `json:"stock_symbol" binding:"required"`
}

func (server *Server) createCompany(ctx *gin.Context) {
	var req createCompanyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateCompanyParams{
		CompanyName: req.CompanyName,
		StockSymbol: req.StockSymbol,
	}

	company, err := server.store.CreateCompany(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, company)
}

type getCompanyRequest struct {
	CompanyName string `uri:"company_name" binding:"required"`
}


func (server *Server) getCompanyByName(ctx *gin.Context) {
	var req getCompanyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	company, err := server.store.GetCompanyByName(ctx, req.CompanyName)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, company)
}

type listCompanyRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=100"`
}

func (server *Server) listCompany(ctx *gin.Context) {
	var req listCompanyRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListCompanyParams{
		Offset: (req.PageID - 1) * req.PageSize,
		Limit:  req.PageSize,
	}
	companys, err := server.store.ListCompany(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, companys)
}

type updateCompanyNameRequest struct {
	ID          int64  `json:"id" binding:"required,min=1"`
	CompanyName string `json:"company_name" binding:"required"`
}

func (server *Server) updateCompanyName(ctx *gin.Context) {
	var req updateCompanyNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCompanyNameParams{
		ID:          req.ID,
		CompanyName: req.CompanyName,
	}

	company, err := server.store.UpdateCompanyName(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, company)
}

type updateCompanyStockSymbolRequest struct {
	ID          int64  `json:"id" binding:"required,min=1"`
	StockSymbol string `json:"stock_symbol" binding:"required"`
}

func (server *Server) updateCompanyStockSymbol(ctx *gin.Context) {
	var req updateCompanyStockSymbolRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCompanyStockSymbolParams{
		ID:          req.ID,
		StockSymbol: req.StockSymbol,
	}

	company, err := server.store.UpdateCompanyStockSymbol(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, company)
}

type deleteCompanyRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteCompany(ctx *gin.Context) {
	var req deleteCompanyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCompany(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, true)
}
