package routes

import (
	"bankapp/internal/events"
	"bankapp/internal/health"
	"bankapp/internal/statements"
	"bankapp/internal/transactions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	routes := gin.Default()

	routes.POST("/events", events.ProcessTransactionEvents)

	// Other API endpoints (placeholder)
	routes.GET("/events", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Events endpoint"})
	})

	routes.GET("/transactions/:id", transactions.GetHistoryTransactions)

	routes.GET("/health", health.Health)

	///statement/{userId}/{AccountType}/{CurrencyType}/{period}
	routes.GET("/statement/:userID/:accountType/:currency/:period", statements.GetStatements)

	return routes
}
