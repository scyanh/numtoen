package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/scyanh/numtoen/app/infrastructure/controllers/numberController"
	middleware "github.com/scyanh/numtoen/app/infrastructure/middlewares"
	"github.com/scyanh/numtoen/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"time"
)

type ginRouter struct{}

var (
	ginEngine *gin.Engine
)

func NewGinRouter() Router {
	return &ginRouter{}
}

func (*ginRouter) POST(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context)) {
	if groupPath != nil {
		groupPath.POST(relativePath, f)
		return
	}
	ginEngine.POST(relativePath, f)
}
func (*ginRouter) DELETE(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context)) {
	if groupPath != nil {
		groupPath.DELETE(relativePath, f)
		return
	}
	ginEngine.DELETE(relativePath, f)
}
func (*ginRouter) PUT(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context)) {
	if groupPath != nil {
		groupPath.PUT(relativePath, f)
		return
	}
	ginEngine.PUT(relativePath, f)
}

func (*ginRouter) GET(groupPath *gin.RouterGroup, relativePath string, f func(c *gin.Context)) {
	if groupPath != nil {
		groupPath.GET(relativePath, f)
		return
	}
	ginEngine.GET(relativePath, f)
}

func (*ginRouter) SERVE(port string) {
	ginEngine.Run(":" + port)
}

func (*ginRouter) INIT() {

	ENV_NAME := os.Getenv("ENV_NAME")
	if ENV_NAME == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	ginEngine = gin.New()
	ginEngine.Use(gin.Recovery())

	//Documentation start - programatically set swagger info
	docs.SwaggerInfo.Title = "Number to English API"
	docs.SwaggerInfo.Description = "This is a Number to English API Translator."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// use ginSwagger middleware to serve the API docs
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (*ginRouter) SETUP_MIDDLEWARES() {
	ginEngine.Use(middleware.CORSMiddleware())
}

func (*ginRouter) SETUP_ROUTES() {
	logger, _ := zap.NewProduction()
	ginEngine.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	ginEngine.Use(ginzap.RecoveryWithZap(logger, true))

	ginEngine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "server online"})
	})

	ginEngine.GET("/num_to_english", numberController.TranslateToEnglish)
}
