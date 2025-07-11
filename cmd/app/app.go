package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/intitalizer"
	// "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/internal/repository"
	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/db"
	sharedRedis "gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/shared/clients/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	cfg    *config.Env
	router *gin.Engine
	server *http.Server
	db     *gorm.DB
	redis  sharedRedis.Client
	log    *zap.Logger
	repo   *intitalizer.BaseRepository
	// service *service.Services
}

func NewApp(cfg *config.Env, log *zap.Logger) *App {
	app := &App{
		cfg: cfg,
		log: log,
	}
	app.initialize()
	return app
}

func (a *App) initialize() {
	// Connect DB
	dbConn, err := db.ConnectDb(a.cfg)
	if err != nil {
		a.log.Fatal("failed to initialize database", zap.Error(err))
	}
	a.db = dbConn

	// Connect Redis
	redisConn, err := sharedRedis.InitRedis(a.cfg)
	if err != nil {
		a.log.Fatal("failed to initialize redis", zap.Error(err))
	}
	a.redis = redisConn

	// Init Repository Access and Factory
	// access := repository.NewRepositoryAccess(a.db, a.redis, a.log, a.cfg)
	a.repo = intitalizer.NewBaseRepository(a.db, a.redis, a.log, a.cfg)

	// // Init Services
	// a.service = service.NewServices(a.repo, a.log, a.cfg)

	// Init Router
	a.router = gin.Default()
	a.registerRoutes()

	// Init HTTP server
	a.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", a.cfg.Port),
		Handler:      a.router,
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
		IdleTimeout:  time.Second * 60,
	}
}

func (a *App) registerRoutes() {
	// Example route
	a.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// TODO: Register actual routes using a.service
}

func (a *App) Run() {
	a.log.Info("starting HTTP server", zap.String("addr", a.server.Addr))
	if err := a.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		a.log.Fatal("server error", zap.Error(err))
	}
}
