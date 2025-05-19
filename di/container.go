package di

import (
	"path/filepath"
	"pongsatorn/go-http-template/internal/adapter/auth/token"
	"pongsatorn/go-http-template/internal/adapter/config"
	"pongsatorn/go-http-template/internal/adapter/http"
	mongoStorage "pongsatorn/go-http-template/internal/adapter/storage/mongo"
	"pongsatorn/go-http-template/internal/adapter/storage/mongo/repository"
	"pongsatorn/go-http-template/internal/core/service"
	"pongsatorn/go-http-template/pkg/logger"
	"runtime"

	"github.com/joho/godotenv"
)

type Application struct {
	HttpServer *http.Server
}

func InitApplication() (*Application, func(), error) {
	_, filename, _, _ := runtime.Caller(0)
	basedir := filepath.Dir(filename)
	configPath := filepath.Join(basedir, "..", ".env")
	err := godotenv.Load(configPath)
	if err != nil {
		logger.Logger.Fatalf("fail to load config, %v", err)
	}

	// init config
	cfg := config.NewConfig()

	// init DB connection
	mongoDB, cleanUpFn1, err := mongoStorage.NewMongoDatabse(cfg.MongoConfig)
	if err != nil {
		return nil, cleanUpFn1, err
	}

	// init auth
	jwtManager := token.NewJwtManager(cfg.JwtConfig)

	// init repository
	templareRepository := repository.NeTemplateRepository(mongoDB)

	// init internal service
	templateService := service.NewTemplateService(templareRepository)

	// init http and related components
	httpMdw := http.NewCustomMiddleware(jwtManager)
	httpHandler := http.NewHttpHandler(templateService)
	httpEchoWrapepr := http.NewEchoWrapper(httpMdw, httpHandler)
	httpServer, cleanUpFn2 := http.NewServer(cfg.HttpServerConfig, httpEchoWrapepr)

	// init background Job
	return &Application{
			HttpServer: httpServer,
		}, func() {
			cleanUpFn2()
			cleanUpFn1()
		}, nil
}
