package api

import (
	db "stock-prediction-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our stock
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// New server creates a new HTTP Server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()
	router.POST("/companys", server.createCompany)
	router.GET("/companys/:company_name", server.getCompany)
	router.GET("/companys", server.listCompany)

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
