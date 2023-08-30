package repository

import (
	"context"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/model"
)

//counterfeiter:generate -o ./../mocks . ProfileCfgRepo
type ProfileCfgRepo interface {
	StoreProfileCfg(c context.Context, profileCfg model.ProfileCfg) error
	UpdateProfileCfg(c context.Context, profileCfg model.ProfileCfg) error
	GetProfileCfgByNameAndID(c context.Context, profileID string, configName string) (model.ProfileCfg, error)
	GetProfileCfgByScheduler(c context.Context, profileCfgSche dto.ProfileCfgSche) ([]model.ProfileCfg, error)
	UnitOfWork
}
