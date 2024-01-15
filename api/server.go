package api

import (
	"api/token"
	"api/util"
	"fmt"

	db "api/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new Server instance with the given store.
//
// Parameters:
// - store: a pointer to a db.Store object.
//
// Returns:
// - a pointer to the newly created Server object.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	// tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	// replace the following line with the line above to use JWT instead of Paseto
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/users/:id/logout", server.logoutUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	// middleware
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/users/:id", server.getUser)
	authRoutes.GET("/users", server.listUsers)

	server.router = router
}

// Start starts the server on the specified address.
//
// Parameters:
// - address: the address to listen on.
//
// Returns:
// - error: an error if the server failed to start.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse generates an error response in the form of a gin.H map.
//
// The function takes an error as a parameter and returns a gin.H map
// with a single key "error" and the error message as its value.
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
