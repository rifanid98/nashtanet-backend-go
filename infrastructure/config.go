package infrastructure

import (
	"errors"
	"nashtanet-backend-go/infrastructure/database"
	"nashtanet-backend-go/infrastructure/logger"
	"nashtanet-backend-go/infrastructure/router"
	"nashtanet-backend-go/infrastructure/validation"
	"strconv"
	"time"
)

type config struct {
	appName       string
	logger        logger.Logger
	validator     validation.Validator
	dbGorm        database.Gorm
	ctxTimeout    time.Duration
	webServerPort router.Port
	webServer     router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) AppName(name string) *config {
	c.appName = name
	return c
}

func (c *config) Logger(instance int) *config {
	log, err := logger.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.logger = log
	c.logger.Infof("Logger configured successfully")
	return c
}

func (c *config) Validator(instance int) *config {
	validator, err := validation.NewValidatorFactory(instance)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.validator = validator
	c.logger.Infof("Validator configured successfully")
	return c
}


func (c *config) Database(dbType int, dbInstance int) *config {
	db, err := database.NewDatabaseFactory(dbType, dbInstance)
	if err != nil {
		c.logger.Fatalln(err, "could not make a connection to the database")
	}

	switch dbType {
	case database.InstanceOrmDatabase:
		switch dbInstance {
		case database.InstanceGorm:
			c.dbGorm = db.(database.Gorm)
		default:
			err = errors.New(" database orm instance does not exists")
		}
	default:
		err = errors.New(" database type instance does not exists")
	}

	c.logger.Infof("successfully connected to the SQL database")
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.logger.Fatalln(err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) WebServer(routerInstance int, webServerInstance int) *config {
	var (
		webServerOrm = &router.RouterFactoryOptions{
			RouterOrmOptions: &router.RouterOrmFactoryOptions{
				EchoEngineGorm: &router.EchoEngineGormOptions{
					Db:        c.dbGorm,
					Validator: c.validator,
					Port:      c.webServerPort,
					T:         c.ctxTimeout,
				},
			},
		}

		s   router.Server
		err error
	)

	switch routerInstance {
	case router.InstanceRouterOrm:
		switch webServerInstance {
		case router.InstanceEchoGorm:
			s, err = router.NewRouterFactory(
				routerInstance,
				webServerInstance,
				webServerOrm,
			)
		default:
			err = errors.New(" web server orm instance does not exists")
		}
	default:
		err = errors.New("router orm instance does not exists")
	}

	if err != nil {
		c.logger.Fatalln(err)
	}

	c.logger.Infof("router server configured successfully")

	c.webServer = s
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}