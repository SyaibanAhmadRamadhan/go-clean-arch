package model

import (
	"database/sql"
	"time"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
	uuid "github.com/satori/go.uuid"
)

type Profile struct {
	ProfileID string
	UserID    string
	Quote     sql.NullString
	CreatedAt int64
	CreatedBy string
	UpdatedAt int64
	UpdatedBy sql.NullString
	DeletedAt sql.NullInt64
	DeletedBy sql.NullString
}

func (p *Profile) DefaultValue(userID string) *Profile {
	id := uuid.NewV4().String()
	return &Profile{
		ProfileID: id,
		UserID:    userID,
		Quote:     sql.NullString{},
		CreatedAt: time.Now().Unix(),
		CreatedBy: id,
		UpdatedAt: time.Now().Unix(),
		UpdatedBy: sql.NullString{},
		DeletedAt: sql.NullInt64{},
		DeletedBy: sql.NullString{},
	}
}

func (p *Profile) ToResp() *dto.ProfileResp {
	return &dto.ProfileResp{
		ProfileID: p.ProfileID,
		Quote:     p.Quote,
	}
}
