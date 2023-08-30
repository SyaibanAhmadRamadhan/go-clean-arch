package usecase

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o ./../mocks . AccountUsecase
type AccountUsecase interface {
	UpdateAccount(context.Context, *dto.UpdateAccountReq) (*dto.UserResp, *dto.ProfileResp, error)
}
