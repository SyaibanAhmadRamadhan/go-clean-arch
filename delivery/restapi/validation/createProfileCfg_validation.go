package validation

import (
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/delivery/restapi/response"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/dto"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/model"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/internal/utils"
	"strconv"
	"strings"
)

func CreateProfileCfg(req *dto.CreateProfileCfgReq) error {
	badReq := map[string][]string{}

	if req.UserID == "" {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) > 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}
	if len(req.UserID) < 36 {
		return response.Err401(model.ErrUnauthorization.Error(), nil)
	}

	// profile_id validation
	if req.ProfileID == "" {
		return response.Err404("NOT FOUND", nil)
	}
	if len(req.ProfileID) > 36 {
		return response.Err404("NOT FOUND", nil)
	}
	if len(req.ProfileID) < 36 {
		return response.Err404("NOT FOUND", nil)
	}

	// config_name validation
	if req.ConfigName != "DAILY_NOTIFY" && req.ConfigName != "MONTHLY_PERIOD" {
		badReq["config_name"] = append(badReq["config_name"], fmt.Sprintf(InvalidField, "config_name", "your input", "DAILY_NOTIFY or MONTHLY_PERIOD"))
	}

	// config_value validation
	if req.ConfigValue == "" {
		badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(Required, "config_value"))
	}
	if req.ConfigName == "DAILY_NOTIFY" && len(strings.Split(req.ConfigValue, " ")) != 2 {
		badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(InvalidField, "config_value", "your input value", "19:20 Asia/Jakarta"))
	}
	if req.ConfigName == "DAILY_NOTIFY" && len(strings.Split(req.ConfigValue, " ")) == 2 {
		if strings.Split(req.ConfigValue, " ")[0] == "" || strings.Split(req.ConfigValue, " ")[1] == "" {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(InvalidField, "config_value", "your input value", "19:20 Asia/Jakarta"))
		}
	}

	if req.ConfigName == "MONTHLY_PERIOD" {
		configValueInt, err := strconv.Atoi(req.ConfigValue)
		if err != nil {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(Integer, "config_value"))
		}

		if configValueInt > 29 {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(MaxInteger, "config_value", 29))
		}
		if configValueInt < 1 {
			badReq["config_value"] = append(badReq["config_value"], fmt.Sprintf(MinInteger, "config_value", 1))
		}
	}

	if req.ConfigName == "DAILY_NOTIFY" {
		if len(req.Days) < 1 {
			badReq["days"] = append(badReq["days"], fmt.Sprintf(InvalidField, "days", "your days must be >= 1", strings.Join(utils.Days(), ", ")))
		}
		if !dayValidate(req.Days) {
			badReq["days"] = append(badReq["days"], fmt.Sprintf(InvalidField, "days", "your input", strings.Join(utils.Days(), ", ")))
		}
	}

	// status validation
	if req.Status != "on" && req.Status != "off" {
		badReq["status"] = append(badReq["status"], fmt.Sprintf(Enum, "status", strings.Join([]string{"on", "off"}, " or ")))
	}

	// token validation
	if req.Token == "" {
		badReq["token"] = append(badReq["token"], fmt.Sprintf(Required, "token"))
	}
	if len(req.Token) < 10 {
		badReq["token"] = append(badReq["token"], fmt.Sprintf(MinInteger, "token", 10))
	}

	if len(badReq) >= 1 {
		return response.Err400(badReq, nil)
	}
	return nil
}
