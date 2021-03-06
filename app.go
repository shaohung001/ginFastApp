package ginFastApp

import (
	"github.com/gin-gonic/gin"
)

type IConfig interface {
	GetPort() int
	GetHost() string
	GetDEBUG() bool
	GetDB() map[string]interface{}
	GetRedis() map[string]interface{}
}

type App struct {
	Config IConfig
	beforeMiddlewares []gin.HandlerFunc
	afterMiddlewares []gin.HandlerFunc
	Routes []IRoute
}

func New(config IConfig)*App {
	return &App{Config:config}
}

// Start 服务器正式启动开始
func (app *App) Start() (*gin.Engine, error)  {
	engine := gin.Default()
	app.engineAssembleMiddlewares(engine)
	return engine, nil
}

