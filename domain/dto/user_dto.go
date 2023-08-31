package dto

import "database/sql"

type UserResp struct {
	ID              string         `json:"id"`
	FullName        string         `json:"full_name"`
	Gender          string         `json:"gender"`
	Image           string         `json:"image"`
	Username        string         `json:"username"`
	Email           string         `json:"email"`
	EmailFormat     string         `json:"email_format"`
	PhoneNumber     sql.NullString `json:"phone_number"`
	EmailVerifiedAt bool           `json:"activated"`
}
