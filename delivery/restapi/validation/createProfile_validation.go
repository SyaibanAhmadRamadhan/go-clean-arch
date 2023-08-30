package validation

import (
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/delivery/restapi/response"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/model"
)

func StoreProfileValidation(req *dto.StoreProfileReq) error {
	if req.UserID == "" {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) > 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) < 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}

	return nil
}
