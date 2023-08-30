package helpers

import (
	"database/sql"
	"time"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/model"
	uuid "github.com/satori/go.uuid"
)

func CreateProfileCfgToModel(req dto.CreateProfileCfgReq, configValue []byte) model.ProfileCfg {
	id := uuid.NewV4().String()
	return model.ProfileCfg{
		ID:          id,
		ProfileID:   req.ProfileID,
		ConfigName:  req.ConfigName,
		ConfigValue: string(configValue),
		Status:      req.Status,
		CreatedAt:   time.Now().Unix(),
		CreatedBy:   req.ProfileID,
		UpdatedAt:   time.Now().Unix(),
		UpdatedBy:   sql.NullString{},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
}

func UpdateProfileCfgToModel(req dto.UpdateProfileCfgReq, configValue []byte, configName, id string) model.ProfileCfg {
	return model.ProfileCfg{
		ID:          id,
		ProfileID:   req.ProfileID,
		ConfigName:  configName,
		ConfigValue: string(configValue),
		Status:      req.Status,
		CreatedAt:   time.Now().Unix(),
		CreatedBy:   req.ProfileID,
		UpdatedAt:   time.Now().Unix(),
		UpdatedBy:   sql.NullString{String: req.ProfileID},
		DeletedAt:   sql.NullInt64{},
		DeletedBy:   sql.NullString{},
	}
}
