package api

import (
	db "github.com/emmybritt/bank_app/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our banking service
type Server struct{
	store db.Store
	router *gin.Engine
}

// NewServer Create a new http Serever and setup routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)
	server.router = router;
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
