package api

import (
	db "stock-prediction-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our stock
type Server struct {
	exchange *db.Exchange
	router   *gin.Engine
}

// New server creates a new HTTP Server and setup routing
func NewServer(exchange *db.Exchange) *Server {
	server := &Server{exchange: exchange}

	router := gin.Default()
	router.POST("/companys", server.createCompany)

	server.router = router
	return server

}
