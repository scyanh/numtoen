package main

import (
	"fmt"
	"github.com/go-redis/redis"
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

	addr := os.Getenv("ENV_REDIS_LOCAL_URL")

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if err := c.Ping().Err(); err != nil {
		zap.S().Errorw("Redis Initialize err", "addr", addr, "err", err)
		return
	}

	fmt.Println("running redis ok")

	//Run Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	httpRouter.SERVE(port)

}
