package handler_echo

import (
	"github.com/labstack/echo/v4"
	"nashtanet-backend-go/common/response"
	"nashtanet-backend-go/infrastructure/validation"
	"nashtanet-backend-go/usecase"
	"net/http"
)

type SignupHandler struct {
	log       echo.Logger
	uc        usecase.SignupUseCase
	validator validation.Validator

	logKey, logMsg string
}

func NewSignupHandler(uc usecase.SignupUseCase, log echo.Logger, v validation.Validator) SignupHandler {
	return SignupHandler{
		log:       log,
		validator: v,
		uc:        uc,
		logKey:    "health_check",
		logMsg:    "health_checking",
	}
}

func (h *SignupHandler) Execute(c echo.Context) error {
	var input usecase.SignupInput
	err := c.Bind(&input)
	if err != nil {
		h.log.Errorf("failed to bind payload", err)
		return c.JSON(
			http.StatusBadRequest,
			response.NewError(response.ErrInvalidInput, http.StatusBadRequest),
		)
	}

	if errs := h.ValidateInput(input); len(errs) > 0 {
		h.log.Info("invalid payload")
		return c.JSON(
			http.StatusBadRequest,
			response.NewErrorMessage(errs, http.StatusBadRequest),
		)
	}

	output, err := h.uc.Execute(input)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			response.NewError(response.ErrInternalServerError, http.StatusInternalServerError),
		)
	}

	return c.JSON(
		http.StatusOK,
		response.NewSuccess(output, http.StatusOK),
	)
}

func (h *SignupHandler) ValidateInput(input usecase.SignupInput) []string {
	var msgs []string

	err := h.validator.Validate(input)
	if err != nil {
		msgs = append(msgs, h.validator.Messages()...)
	}

	return msgs
}
