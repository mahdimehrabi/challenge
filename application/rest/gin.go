package rest

import (
	"log"

	"challenge/application/rest/trade/controller"
	"challenge/internal/dpi"
	infrastructures "challenge/internal/infrastructure"
	"challenge/internal/repository/trade/pgx"
	"challenge/internal/service"

	"github.com/gin-gonic/gin"
)

func Setup(env *infrastructures.Env) {
	r := gin.Default()

	// for sake of time limitation I didn't create routing files and called controllers directly
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// for best performance
	// I use transient dependency injection for creating controllers and its dependencies
	// even one instance for each endpoint except the ones that use resources like database and socket conns etc.
	// maybe its good idea to use some packages like uber.FX,
	// but I didn't because they decrease the performance and lack of the time
	tc := controller.NewTradeController(service.NewTradeService(pgx.NewTradeRepository(dpi.PGXPool)))
	r.POST("/api/trade", tc.Create)

	if err := r.Run(":" + env.ServerPort); err != nil {
		log.Fatal(err)
	}
}
