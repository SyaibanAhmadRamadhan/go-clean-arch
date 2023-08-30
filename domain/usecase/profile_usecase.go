package usecase

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
)

//counterfeiter:generate -o ./../mocks . ProfileUsecase
type ProfileUsecase interface {
	GetProfileByID(c context.Context, req *dto.GetProfileReq) (*dto.ProfileResp, error)
	StoreProfile(c context.Context, req *dto.StoreProfileReq) (*dto.ProfileResp, error)
}
