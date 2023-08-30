package usecase

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
)

//counterfeiter:generate -o ./../mocks . ProfileCfgUsecase
type ProfileCfgUsecase interface {
	CreateProfileCfg(context.Context, dto.CreateProfileCfgReq) (dto.ProfileCfgResp, error)
	GetProfileCfgByNameAndID(c context.Context, req dto.GetProfileCfgReq) (dto.ProfileCfgResp, error)
	UpdateProfileCfg(c context.Context, req dto.UpdateProfileCfgReq) (dto.ProfileCfgResp, error)
}
