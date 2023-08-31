package restapi

import (
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/delivery/restapi/response"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/delivery/restapi/validation"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/model"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/usecase"
	"net/http"
)

type ProfileHandler struct {
	profileUsecase usecase.ProfileUsecase
}

func NewProfileHandler(
	profileUsecase usecase.ProfileUsecase,
) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase: profileUsecase,
	}
}

func (h *ProfileHandler) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	req := new(dto.GetProfileReq)

	userId := r.Header.Get("User-Id")

	req.UserID = userId

	err := validation.GetProfileValidation(req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	profile, err := h.profileUsecase.GetProfileByID(r.Context(), req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	resp := model.ResponseSuccess{
		Data: profile,
	}

	response.NewSucc(w, r, resp, 200)
}

func (h *ProfileHandler) StoreProfile(w http.ResponseWriter, r *http.Request) {
	req := new(dto.StoreProfileReq)

	err := response.DecodeReq(r, req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	err = validation.StoreProfileValidation(req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}

	profile, err := h.profileUsecase.StoreProfile(r.Context(), req)
	if err != nil {
		response.NewError(w, r, err)
		return
	}
	resp := model.ResponseSuccess{
		Data: profile,
	}
	response.NewSucc(w, r, resp, 201)
}
