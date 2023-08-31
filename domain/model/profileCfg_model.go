package model

import "database/sql"

type ProfileCfg struct {
	ID          string
	ProfileID   string
	ConfigName  string
	ConfigValue string
	Status      string
	CreatedAt   int64
	CreatedBy   string
	UpdatedAt   int64
	UpdatedBy   sql.NullString
	DeletedAt   sql.NullInt64
	DeletedBy   sql.NullString
}
