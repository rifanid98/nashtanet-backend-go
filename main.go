package main

import (
	"nashtanet-backend-go/infrastructure"
	"nashtanet-backend-go/infrastructure/database"
	"nashtanet-backend-go/infrastructure/logger"
	"nashtanet-backend-go/infrastructure/router"
	"nashtanet-backend-go/infrastructure/validation"
	"os"
)

func main() {
	// initialize application dependencies
	app := infrastructure.NewConfig().
		AppName(os.Getenv("APP_NAME")).
		Logger(logger.InstanceZapLogger).
		Validator(validation.InstanceGoPlayground).
		Database(database.InstanceOrmDatabase, database.InstanceGorm)

	// initialize application server
	app.WebServerPort("8080").
		WebServer(router.InstanceRouterOrm, router.InstanceEchoGorm)

	// start application
	app.Start()
}
