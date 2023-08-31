package unit

import (
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/internal/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatEmail(t *testing.T) {
	t.Run("SUCCESS_FormatEmail", func(t *testing.T) {
		expect := "i••••9@gmail.com"
		email := "ibanrama29@gmail.com"
		emailFormat := helpers.EmailFormat(email)
		assert.Equal(t, expect, emailFormat)
	})

	t.Run("ERROR_FormatEmail", func(t *testing.T) {
		expect := "i••••9@gmail.com"
		email := "ibanrama29gmail.com"
		emailFormat := helpers.EmailFormat(email)
		assert.NotEqual(t, expect, emailFormat)
	})
}

func TestConfigValue(t *testing.T) {
	t.Run("SUCCESS_ConfigValueDAILY_NOTIFY", func(t *testing.T) {
		configName := "DAILY_NOTIFY"
		value := "19:00"
		ianaTimezone := "Asia/Jakarta"
		days := []string{
			"monday",
			"yesterday",
		}

		hour := 12
		minute := 0o0
		expect := map[string]any{
			"config_time_user":       value,
			"config_timezone_user":   ianaTimezone,
			"config_time_notify":     fmt.Sprintf("%02d:%02d", hour, minute),
			"config_timezone_notify": "UTC",
			"days":                   days,
		}

		configValue, err := helpers.ConfigValue(configName, value, ianaTimezone, days)
		assert.NoError(t, err)
		assert.Equal(t, expect, configValue)
	})
	t.Run("SUCCESS_ConfigValueMONTHLY_PERIOD", func(t *testing.T) {
		configName := "MONTHLY_PERIOD"
		value := "10"
		ianaTimezone := ""
		var days []string

		expect := map[string]any{
			"config_date": value,
		}

		configValue, err := helpers.ConfigValue(configName, value, ianaTimezone, days)
		assert.NoError(t, err)
		assert.Equal(t, expect, configValue)
	})

	// t.Run("ERROR_ConfigValue", func(t *testing.T) {
	// })
}
