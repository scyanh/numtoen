package main

import (
	router "github.com/scyanh/numtoen/app/infrastructure/routes"
	"go.uber.org/zap"
	"os"
)

var (
	httpRouter = router.NewGinRouter()
)

func main() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)

	//Init Routes
	httpRouter.INIT()
	httpRouter.SETUP_MIDDLEWARES()
	httpRouter.SETUP_ROUTES()

	//Run Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	httpRouter.SERVE(port)

}
