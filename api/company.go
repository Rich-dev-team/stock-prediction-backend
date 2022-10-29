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

func (server *Server) getCompany(ctx *gin.Context) {
	var req getCompanyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	company, err := server.store.GetCompany(ctx, req.CompanyName)
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
