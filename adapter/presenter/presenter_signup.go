package presenter

import (
	"nashtanet-backend-go/domain/entity"
	"nashtanet-backend-go/usecase"
)

type signupPresenter struct {
}

func NewSignupPresenter() usecase.SignupPresenter {
	return signupPresenter{}
}

func (s signupPresenter) Output(user *entity.User) usecase.SignupOutput {
	return usecase.SignupOutput{
		Email: *user.Email,
	}
}
