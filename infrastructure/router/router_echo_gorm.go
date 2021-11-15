package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"nashtanet-backend-go/adapter/handler/handler_echo"
	"nashtanet-backend-go/adapter/presenter"
	"nashtanet-backend-go/domain/entity"
	"nashtanet-backend-go/infrastructure/database"
	"nashtanet-backend-go/infrastructure/persistence/repository_gorm"
	"nashtanet-backend-go/infrastructure/validation"
	"nashtanet-backend-go/usecase"
	"time"
)

type echoEngineGorm struct {
	router     *echo.Echo
	db         database.Gorm
	validator  validation.Validator
	port       Port
	ctxTimeout time.Duration
}

type EchoEngineGormOptions struct {
	Db        database.Gorm
	Validator validation.Validator
	Port      Port
	T         time.Duration
}

func NewEchoServerGorm(options *EchoEngineGormOptions) *echoEngineGorm {
	return &echoEngineGorm{
		router:     echo.New(),
		db:         options.Db,
		validator:  options.Validator,
		port:       options.Port,
		ctxTimeout: options.T,
	}
}

func (e *echoEngineGorm) Listen() {
	e.setAppHandlers()

	// AutoMigrate(domain struct)
	err := e.db.AutoMigrate(&entity.User{}, &entity.Role{})
	if err != nil {
		return
	}

	e.router.Logger.Fatal(e.router.Start(fmt.Sprintf(":%v", e.port)))
}

func (e *echoEngineGorm) setAppHandlers() {
	e.router.POST("/auth/signup", e.buildSignupHandler)
}

func (e *echoEngineGorm) buildSignupHandler(c echo.Context) error {
	var (
		uc = usecase.NewSignupInteractor(
			repository_gorm.NewUserRepositoryGorm(e.db),
			presenter.NewSignupPresenter(),
			e.ctxTimeout,
		)

		act = handler_echo.NewSignupHandler(uc, e.router.Logger, e.validator)
	)

	return act.Execute(c)
}
