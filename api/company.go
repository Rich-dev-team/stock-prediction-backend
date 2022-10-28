package api

import (
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
