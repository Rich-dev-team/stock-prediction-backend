package api

import (
	db "github.com/Rich-dev-team/stock-prediction-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our stock
type Server struct {
	store  db.Store
	router *gin.Engine
}

// New server creates a new HTTP Server and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	// company
	router.POST("/companys/create", server.createCompany)
	router.GET("/companys/list/:company_name", server.getCompanyByName)
	router.GET("/companys/lists", server.listCompany)
	router.PUT("/companys/update/name", server.updateCompanyName)
	router.PUT("/companys/update/symbol", server.updateCompanyStockSymbol)
	router.DELETE("/companys/:id", server.deleteCompany)
	// stock price
	router.POST("/stockprice/create/one", server.createStockPrice)
	// router.POST("/stockprice/create/batch", server.createBatchStockPrice)
	router.GET("/stockprice/lists/all", server.listAllStockPrice)
	router.GET("/stockprice/lists/range", server.listStockPriceByRange)
	router.PUT("/stockprice/price", server.UpdateStockPrice)
	router.DELETE("/stockprice/:id", server.deleteStockPrice)

	// router
	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
