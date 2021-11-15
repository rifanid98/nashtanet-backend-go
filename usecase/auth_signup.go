package usecase

import (
	"nashtanet-backend-go/domain/entity"
	"nashtanet-backend-go/domain/repository"
	"time"
)

type (
	SignupInput struct {
		Email    string `json:"email" validate:"email,required"`
		Password string `json:"password" validate:"min=6,max=32,required"`
	}

	SignupOutput struct {
		Email string `json:"email" validate:"email,required"`
	}

	SignupUseCase interface {
		Execute(SignupInput) (SignupOutput, error)
	}

	SignupPresenter interface {
		Output(user *entity.User) SignupOutput
	}

	signupInteractor struct {
		userRepo   repository.UserRepository
		presenter  SignupPresenter
		ctxTimeout time.Duration
	}
)

func NewSignupInteractor(
	userRepo repository.UserRepository,
	presenter SignupPresenter,
	t time.Duration,
) SignupUseCase {
	return signupInteractor{
		userRepo:   userRepo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (i signupInteractor) Execute(input SignupInput) (SignupOutput, error) {
	user := entity.NewUser()
	user.Email = &input.Email
	user.Password = &input.Password

	registered, err := i.userRepo.CreateUser(user)

	if err != nil {
		return SignupOutput{Email: input.Email}, err
	}

	return i.presenter.Output(registered), nil
}
