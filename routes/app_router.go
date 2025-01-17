package routes

import (

	"github.com/easily-mistaken/OpinX-backend/handlers"
	"github.com/gin-gonic/gin"
)

func Run(r *gin.Engine) {
	r = getRoutes()
}

func getRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/balance/inr/:userId", func(c *gin.Context) { handlers.GetInrBalances(c.Request) })
	router.GET("/balance/stock/:userId", func(c *gin.Context) { handlers.GetStockBalances(c.Request) })
	router.POST("/onramp/inr", func(c *gin.Context) { handlers.OnRamp(c.Request) })
	router.GET("/orderbook/:stockSymbol", func(c *gin.Context) { handlers.ViewOrder(c.Request) })
	router.POST("/order/yes", func(c *gin.Context) { handlers.BuyOrder(c.Request) })
	router.POST("/order/no", func(c *gin.Context) { handlers.SellOrder(c.Request) })
	router.POST("/trade/mint/:stockSymbol", func(c *gin.Context) { handlers.MintTokens(c.Request) })

	return router
}
