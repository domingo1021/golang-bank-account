package api

import (
	db "github.com/domingo1021/golang-bank-account/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for out banking services.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// New Server creates a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address (becaurse Server.router is private)
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
