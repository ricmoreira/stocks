package server

import (
	"stocks/config"
	"stocks/controllers/v1"
	"stocks/handlers"

	"github.com/gin-gonic/gin"
)

// Server is the http layer for role and user resource
type Server struct {
	config             *config.Config
	stockMovController *controllers.StockMovController
	handlers           *handlers.HttpHandlers
}

// NewServer is the Server constructor
func NewServer(cf *config.Config,
	smc *controllers.StockMovController,
	hand *handlers.HttpHandlers) *Server {

	return &Server{
		config:             cf,
		stockMovController: smc,
		handlers:           hand,
	}
}

// Run loads server with its routes and starts the server
func (s *Server) Run() {
	// Instantiate a new router
	r := gin.Default()

	// generic routes
	r.HandleMethodNotAllowed = false
	r.NoRoute(s.handlers.NotFound)

	// Stock Movement resource
	stockMovAPI := r.Group("/api/v1/stock/movement")
	{
		// Create a new invoice
		stockMovAPI.POST("", s.stockMovController.CreateAction)

		// List stocks with filtering and pagination
		stockMovAPI.GET("", s.stockMovController.ListAction)
	}

	// Stock Count resource
	stockCountAPI := r.Group("/api/v1/stock/count")
	{
		// Aggregates stock counts with filtering and pagination
		stockCountAPI.GET("", s.stockMovController.ListStockMovCountAction)
	}

	// Fire up the server
	r.Run(s.config.Host)
}
